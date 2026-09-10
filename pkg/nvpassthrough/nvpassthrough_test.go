/*
 * Copyright (c) NVIDIA CORPORATION.  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nvpassthrough

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/NVIDIA/go-nvlib/pkg/nvpci"
)

// pciFunc describes one PCI function to materialise in the fake sysfs tree.
type pciFunc struct {
	address string
	vendor  uint16
	driver  string // empty means unbound
}

// newFakePCITree builds a minimal /sys/bus/pci/devices layout and returns its path.
func newFakePCITree(t *testing.T, funcs []pciFunc, consumerLinksOn string, consumers []string) string {
	t.Helper()
	root := t.TempDir()
	devicesRoot := filepath.Join(root, "devices")
	driversRoot := filepath.Join(root, "drivers")
	require.NoError(t, os.MkdirAll(devicesRoot, 0755))
	require.NoError(t, os.MkdirAll(driversRoot, 0755))

	for _, f := range funcs {
		devPath := filepath.Join(devicesRoot, f.address)
		require.NoError(t, os.MkdirAll(devPath, 0755))
		vendor := fmt.Sprintf("0x%04x\n", f.vendor)
		require.NoError(t, os.WriteFile(filepath.Join(devPath, "vendor"), []byte(vendor), 0644))
		if f.driver != "" {
			drvPath := filepath.Join(driversRoot, f.driver)
			require.NoError(t, os.MkdirAll(drvPath, 0755))
			require.NoError(t, os.Symlink(drvPath, filepath.Join(devPath, "driver")))
		}
	}

	// "consumer:pci:<addr>" entries live in the GPU's own sysfs directory.
	for _, c := range consumers {
		require.NoError(t, os.MkdirAll(
			filepath.Join(devicesRoot, consumerLinksOn, consumerPrefix+c), 0755))
	}
	return devicesRoot
}

func addressesOf(devs []*nvidiaPCIAuxDevice) []string {
	out := make([]string, 0, len(devs))
	for _, d := range devs {
		out = append(out, d.Address)
	}
	sort.Strings(out)
	return out
}

func TestGetAuxDevices(t *testing.T) {
	const nvidia = nvpci.PCINvidiaVendorID

	testCases := []struct {
		description string
		funcs       []pciFunc
		consumers   []string
		gpu         string
		class       uint32 // defaults to PCIVgaControllerClass
		expected    []string
	}{
		{
			description: "turing board: audio, VirtualLink xHCI and USB-C UCSI are all returned",
			// .2/.3 have no consumer link -- they are only reachable via sibling enumeration
			funcs: []pciFunc{
				{"0000:16:00.0", nvidia, "vfio-pci"},
				{"0000:16:00.1", nvidia, "vfio-pci"},
				{"0000:16:00.2", nvidia, "xhci_hcd"},
				{"0000:16:00.3", nvidia, ""},
			},
			consumers: []string{"0000:16:00.1"},
			gpu:       "0000:16:00.0",
			expected:  []string{"0000:16:00.1", "0000:16:00.2", "0000:16:00.3"},
		},
		{
			description: "single-function GPU has no auxiliary devices",
			funcs:       []pciFunc{{"0000:9b:00.0", nvidia, "vfio-pci"}},
			gpu:         "0000:9b:00.0",
			expected:    []string{},
		},
		{
			description: "a sibling function from another vendor is not touched",
			funcs: []pciFunc{
				{"0000:16:00.0", nvidia, "vfio-pci"},
				{"0000:16:00.1", 0x8086, "snd_hda_intel"},
			},
			gpu:      "0000:16:00.0",
			expected: []string{},
		},
		{
			description: "functions on a different slot are not siblings",
			funcs: []pciFunc{
				{"0000:16:00.0", nvidia, "vfio-pci"},
				{"0000:17:00.0", nvidia, "vfio-pci"},
			},
			gpu:      "0000:16:00.0",
			expected: []string{},
		},
		{
			description: "GPU with a single HDA audio function is still returned",
			funcs: []pciFunc{
				{"0000:16:00.0", nvidia, "vfio-pci"},
				{"0000:16:00.1", nvidia, "snd_hda_intel"},
			},
			consumers: []string{"0000:16:00.1"},
			gpu:       "0000:16:00.0",
			expected:  []string{"0000:16:00.1"},
		},
		{
			description: "auxiliary device reachable only via a consumer link is still found",
			funcs: []pciFunc{
				{"0000:16:00.0", nvidia, "vfio-pci"},
				{"0000:20:00.0", nvidia, "snd_hda_intel"},
			},
			consumers: []string{"0000:20:00.0"},
			gpu:       "0000:16:00.0",
			expected:  []string{"0000:20:00.0"},
		},
		{
			description: "3D-controller-class GPU has its auxiliary functions returned",
			funcs: []pciFunc{
				{"0000:16:00.0", nvidia, "vfio-pci"},
				{"0000:16:00.1", nvidia, "snd_hda_intel"},
			},
			gpu:      "0000:16:00.0",
			class:    nvpci.PCI3dControllerClass,
			expected: []string{"0000:16:00.1"},
		},
		{
			description: "NVSwitch has no auxiliary functions handled",
			funcs: []pciFunc{
				{"0000:16:00.0", nvidia, "vfio-pci"},
				{"0000:16:00.1", nvidia, "snd_hda_intel"},
			},
			gpu:      "0000:16:00.0",
			class:    nvpci.PCINvSwitchClass,
			expected: []string{},
		},
		{
			description: "a function found both as a sibling and via a consumer link appears once",
			funcs: []pciFunc{
				{"0000:16:00.0", nvidia, "vfio-pci"},
				{"0000:16:00.1", nvidia, "snd_hda_intel"},
			},
			consumers: []string{"0000:16:00.1"},
			gpu:       "0000:16:00.0",
			expected:  []string{"0000:16:00.1"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			devicesRoot := newFakePCITree(t, tc.funcs, tc.gpu, tc.consumers)
			class := tc.class
			if class == 0 {
				class = nvpci.PCIVgaControllerClass
			}
			dev := &nvpci.NvidiaPCIDevice{
				Address: tc.gpu,
				Path:    filepath.Join(devicesRoot, tc.gpu),
				Class:   class,
			}

			auxDevs, err := getAuxDevices(devicesRoot, dev)
			require.NoError(t, err)
			require.ElementsMatch(t, tc.expected, addressesOf(auxDevs))
		})
	}
}

// The driver currently bound to each auxiliary function must be reported, since
// BindToVFIODriver skips functions already on the vfio driver and unbinds the rest.
func TestGetAuxDevicesReportsCurrentDriver(t *testing.T) {
	const nvidia = nvpci.PCINvidiaVendorID
	funcs := []pciFunc{
		{"0000:16:00.0", nvidia, "vfio-pci"},
		{"0000:16:00.1", nvidia, "vfio-pci"},
		{"0000:16:00.2", nvidia, "xhci_hcd"},
		{"0000:16:00.3", nvidia, ""},
	}
	devicesRoot := newFakePCITree(t, funcs, "0000:16:00.0", nil)
	dev := &nvpci.NvidiaPCIDevice{
		Address: "0000:16:00.0",
		Path:    filepath.Join(devicesRoot, "0000:16:00.0"),
		Class:   nvpci.PCIVgaControllerClass,
	}

	auxDevs, err := getAuxDevices(devicesRoot, dev)
	require.NoError(t, err)

	got := map[string]string{}
	for _, d := range auxDevs {
		got[d.Address] = d.Driver
	}
	require.Equal(t, map[string]string{
		"0000:16:00.1": "vfio-pci",
		"0000:16:00.2": "xhci_hcd",
		"0000:16:00.3": "",
	}, got)
}
