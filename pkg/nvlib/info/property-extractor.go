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

package info

import (
	"fmt"
	"os"
	"strings"
)

type propertyExtractor struct {
	root root
}

var _ Interface = &propertyExtractor{}

// HasDXCore returns true if DXCore is detected on the system.
func (i *propertyExtractor) HasDXCore() (bool, string) {
	const (
		libraryName = "libdxcore.so"
	)
	if err := i.root.assertHasLibrary(libraryName); err != nil {
		return false, fmt.Sprintf("could not load DXCore library: %v", err)
	}

	return true, "found DXCore library"
}

// HasNvml returns true if NVML is detected on the system.
func (i *propertyExtractor) HasNvml() (bool, string) {
	const (
		libraryName = "libnvidia-ml.so.1"
	)
	if err := i.root.assertHasLibrary(libraryName); err != nil {
		return false, fmt.Sprintf("could not load NVML library: %v", err)
	}

	return true, "found NVML library"
}

// IsTegraSystem returns true if the system is detected as a Tegra-based system.
// Deprecated: Use HasTegraFiles instead.
func (i *propertyExtractor) IsTegraSystem() (bool, string) {
	return i.HasTegraFiles()
}

// HasTegraFiles returns true if tegra-based files are detected on the system.
func (i *propertyExtractor) HasTegraFiles() (bool, string) {
	tegraReleaseFile := i.root.join("/etc/nv_tegra_release")
	tegraFamilyFile := i.root.join("/sys/devices/soc0/family")

	if info, err := os.Stat(tegraReleaseFile); err == nil && !info.IsDir() {
		return true, fmt.Sprintf("%v found", tegraReleaseFile)
	}

	if info, err := os.Stat(tegraFamilyFile); err != nil || info.IsDir() {
		return false, fmt.Sprintf("%v file not found", tegraFamilyFile)
	}

	contents, err := os.ReadFile(tegraFamilyFile)
	if err != nil {
		return false, fmt.Sprintf("could not read %v", tegraFamilyFile)
	}

	if strings.HasPrefix(strings.ToLower(string(contents)), "tegra") {
		return true, fmt.Sprintf("%v has 'tegra' prefix", tegraFamilyFile)
	}

	return false, fmt.Sprintf("%v has no 'tegra' prefix", tegraFamilyFile)
}
