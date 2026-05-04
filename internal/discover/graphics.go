/**
# Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package discover

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/NVIDIA/go-nvlib/internal/logger"
	"github.com/NVIDIA/go-nvlib/internal/lookup"
	"github.com/NVIDIA/go-nvlib/internal/lookup/root"
)

// NewGraphicsMountsDiscoverer creates a discoverer for the mounts required by graphics tools such as vulkan.
func NewGraphicsMountsDiscoverer(logger logger.Interface, driver *root.Driver, hookCreator HookCreator) (Discover, error) {
	libraries, err := newGraphicsLibrariesDiscoverer(logger, driver, hookCreator)
	if err != nil {
		return nil, fmt.Errorf("failed to construct discoverer for graphics libraries: %w", err)
	}

	configs := NewMounts(
		logger,
		driver.Configs(),
		driver.Root,
		[]string{
			"glvnd/egl_vendor.d/10_nvidia.json",
			"egl/egl_external_platform.d/15_nvidia_gbm.json",
			"egl/egl_external_platform.d/10_nvidia_wayland.json",
			"egl/egl_external_platform.d/09_nvidia_wayland2.json",
			"nvidia/nvoptix.bin",
			"X11/xorg.conf.d/10-nvidia.conf",
			"X11/xorg.conf.d/nvidia-drm-outputclass.conf",
		},
	)

	discover := Merge(
		libraries,
		configs,
		newVulkanConfigsDiscover(logger, driver),
	)

	return discover, nil
}

// newVulkanConfigsDiscover creates a discoverer for vulkan ICD files.
// For these files we search the standard driver config paths as well as the
// driver root itself. This allows us to support GKE installations where the
// vulkan ICD files are at {{ .driverRoot }}/vulkan instead of in /etc/vulkan.
func newVulkanConfigsDiscover(logger logger.Interface, driver *root.Driver) Discover {
	locator := lookup.First(driver.Configs(), driver.Files())

	required := []string{
		"vulkan/icd.d/nvidia_icd.json",
		"vulkan/icd.d/nvidia_layers.json",
		"vulkan/implicit_layer.d/nvidia_layers.json",
	}
	// For some RPM-based driver packages, the vulkan ICD files are installed to
	// /usr/share/vulkan/icd.d/nvidia_icd.%{_target_cpu}.json
	// We also include this in the list of candidates for the ICD file.
	switch runtime.GOARCH {
	case "amd64":
		required = append(required, "vulkan/icd.d/nvidia_icd.x86_64.json")
	case "arm64":
		required = append(required, "vulkan/icd.d/nvidia_icd.aarch64.json")
	}
	return &mountsToContainerPath{
		logger:        logger,
		locator:       locator,
		required:      required,
		containerRoot: "/etc",
	}
}

type graphicsDriverLibraries struct {
	Discover
	logger      logger.Interface
	hookCreator HookCreator
}

var _ Discover = (*graphicsDriverLibraries)(nil)

func newGraphicsLibrariesDiscoverer(logger logger.Interface, driver *root.Driver, hookCreator HookCreator) (Discover, error) {
	cudaVersionPattern, err := driver.Version()
	if err != nil {
		return nil, fmt.Errorf("failed to get driver version: %w", err)
	}
	cudaLibRoot, err := driver.GetDriverLibDirectory()
	if err != nil {
		return nil, fmt.Errorf("failed to get libcuda.so parent directory: %w", err)
	}

	libraries := NewMounts(
		logger,
		driver.Libraries(),
		driver.Root,
		[]string{
			// The libnvidia-egl-gbm and libnvidia-egl-wayland libraries do not
			// have the RM version. Use the *.* pattern to match X.Y.Z versions.
			"libnvidia-egl-gbm.so.*.*",
			"libnvidia-egl-wayland.so.*.*",
			"libnvidia-egl-wayland2.so.*.*",
			// We include the following libraries to have them available for
			// symlink creation below:
			// If CDI injection is used, these should already be detected as:
			// * libnvidia-allocator.so.RM_VERSION
			// * libnvidia-vulkan-producer.so.RM_VERSION
			// but need to be handled for the legacy case too.
			"libnvidia-allocator.so." + cudaVersionPattern,
			"libnvidia-vulkan-producer.so." + cudaVersionPattern,
		},
	)

	xorgLibraries := NewMounts(
		logger,
		lookup.NewFileLocator(
			lookup.WithLogger(logger),
			lookup.WithRoot(driver.Root),
			lookup.WithSearchPaths(buildXOrgSearchPaths(cudaLibRoot)...),
			lookup.WithCount(1),
		),
		driver.Root,
		[]string{
			"nvidia_drv.so",
			"libglxserver_nvidia.so." + cudaVersionPattern,
		},
	)

	return &graphicsDriverLibraries{
		Discover:    Merge(libraries, xorgLibraries),
		logger:      logger,
		hookCreator: hookCreator,
	}, nil
}

// Mounts discovers the required libraries and filters out libnvidia-allocator.so.
// The library libnvidia-allocator.so is already handled by either the *.RM_VERSION
// injection or by libnvidia-container. We therefore filter it out here as a
// workaround for the case where libnvidia-container will re-mount this in the
// container, which causes issues with shared mount propagation.
func (d graphicsDriverLibraries) Mounts() ([]Mount, error) {
	mounts, err := d.Discover.Mounts()
	if err != nil {
		return nil, fmt.Errorf("failed to get library mounts: %v", err)
	}

	var filtered []Mount
	for _, mount := range mounts {
		if d.isDriverLibrary(filepath.Base(mount.Path), "libnvidia-allocator.so") {
			continue
		}
		filtered = append(filtered, mount)
	}
	return filtered, nil
}

// Create necessary library symlinks for graphics drivers.
func (d graphicsDriverLibraries) Hooks() ([]Hook, error) {
	mounts, err := d.Discover.Mounts()
	if err != nil {
		return nil, fmt.Errorf("failed to get library mounts: %v", err)
	}

	var links []string
	for _, mount := range mounts {
		dir, filename := filepath.Split(mount.Path)
		switch {
		case d.isDriverLibrary(filename, "libnvidia-allocator.so"):
			// gbm/nvidia-drm_gbm.so is a symlink to ../libnvidia-allocator.so.1 which
			// in turn symlinks to libnvidia-allocator.so.RM_VERSION.
			// The libnvidia-allocator.so.1 -> libnvidia-allocator.so.RM_VERSION symlink
			// is created when ldconfig is run against the container and there
			// is no explicit need to create it.
			// create gbm/nvidia-drm_gbm.so -> ../libnvidia-allocate.so.1 symlink
			linkPath := filepath.Join(dir, "gbm", "nvidia-drm_gbm.so")
			links = append(links, fmt.Sprintf("%s::%s", "../libnvidia-allocator.so.1", linkPath))
		case d.isDriverLibrary(filename, "libnvidia-vulkan-producer.so"):
			// libnvidia-vulkan-producer.so is a drirect symlink to libnvidia-vulkan-producer.so.RM_VERSION
			// create libnvidia-vulkan-producer.so -> libnvidia-vulkan-producer.so.RM_VERSION symlink
			linkPath := filepath.Join(dir, "libnvidia-vulkan-producer.so")
			links = append(links, fmt.Sprintf("%s::%s", filename, linkPath))
		case d.isDriverLibrary(filename, "libglxserver_nvidia.so"):
			// libglxserver_nvidia.so is a directl symlink to libglxserver_nvidia.so.RM_VERSION
			// create libglxserver_nvidia.so -> libglxserver_nvidia.so.RM_VERSION symlink
			linkPath := filepath.Join(dir, "libglxserver_nvidia.so")
			links = append(links, fmt.Sprintf("%s::%s", filename, linkPath))
		}
	}
	if len(links) == 0 {
		return nil, nil
	}

	hook := d.hookCreator.Create("create-symlinks", links...)

	return hook.Hooks()
}

// isDriverLibrary checks whether the specified filename is a specific driver library.
func (d graphicsDriverLibraries) isDriverLibrary(filename string, libraryName string) bool {
	// TODO: Instead of `.*.*` we could use the driver version.
	pattern := strings.TrimSuffix(libraryName, ".") + ".*.*"
	match, _ := filepath.Match(pattern, filename)
	return match
}

// buildXOrgSearchPaths returns the ordered list of search paths for XOrg files.
func buildXOrgSearchPaths(libRoot string) []string {
	var paths []string
	if libRoot != "" {
		paths = append(paths,
			filepath.Join(libRoot, "nvidia/xorg"),
			filepath.Join(libRoot, "xorg", "modules", "drivers"),
			filepath.Join(libRoot, "xorg", "modules", "extensions"),
			filepath.Join(libRoot, "xorg", "modules/updates", "drivers"),
			filepath.Join(libRoot, "xorg", "modules/updates", "extensions"),
		)
	}

	return append(paths,
		filepath.Join("/usr/lib/xorg", "modules", "drivers"),
		filepath.Join("/usr/lib/xorg", "modules", "extensions"),
		filepath.Join("/usr/lib/xorg", "modules/updates", "drivers"),
		filepath.Join("/usr/lib/xorg", "modules/updates", "extensions"),
		filepath.Join("/usr/lib64/xorg", "modules", "drivers"),
		filepath.Join("/usr/lib64/xorg", "modules", "extensions"),
		filepath.Join("/usr/lib64/xorg", "modules/updates", "drivers"),
		filepath.Join("/usr/lib64/xorg", "modules/updates", "extensions"),
		filepath.Join("/usr/X11R6/lib", "modules", "drivers"),
		filepath.Join("/usr/X11R6/lib", "modules", "extensions"),
		filepath.Join("/usr/X11R6/lib", "modules/updates", "drivers"),
		filepath.Join("/usr/X11R6/lib", "modules/updates", "extensions"),
		filepath.Join("/usr/X11R6/lib64", "modules", "drivers"),
		filepath.Join("/usr/X11R6/lib64", "modules", "extensions"),
		filepath.Join("/usr/X11R6/lib64", "modules/updates", "drivers"),
		filepath.Join("/usr/X11R6/lib64", "modules/updates", "extensions"),
	)
}
