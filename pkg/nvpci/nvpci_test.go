/*
 * Copyright (c) 2021, NVIDIA CORPORATION.  All rights reserved.
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

package nvpci

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	ga100PmcID = uint32(0x170000a1)
)

func TestNvpci(t *testing.T) {
	nvpci, err := NewMockNvpci()
	require.Nil(t, err, "Error creating NewMockNvpci")
	defer nvpci.Cleanup()

	err = nvpci.AddMockA100("0000:80:05.1", 0)
	require.Nil(t, err, "Error adding Mock A100 device to MockNvpci")

	devices, err := nvpci.GetGPUs()
	require.Nil(t, err, "Error getting GPUs")
	require.Equal(t, 1, len(devices), "Wrong number of GPU devices")
	require.Equal(t, 1, len(devices[0].Resources), "Wrong number GPU resources found")
	require.Equal(t, "0000:80:05.1", devices[0].Address, "Wrong Address found for device")
	require.Equal(t, 0, devices[0].NumaNode, "Wrong NUMA node found for device")

	config, err := devices[0].Config.Read()
	require.Nil(t, err, "Error reading config")
	require.Equal(t, devices[0].Vendor, config.GetVendorID(), "Vendor IDs do not match")
	require.Equal(t, devices[0].Device, config.GetDeviceID(), "Device IDs do not match")
	require.Equal(t, "nvidia", devices[0].Driver, "Wrong driver detected for device")
	require.Equal(t, 20, devices[0].IommuGroup, "Wrong iommu_group detected for device")

	capabilities, err := config.GetPCICapabilities()
	require.Nil(t, err, "Error getting PCI capabilities")
	require.Equal(t, 0, len(capabilities.Standard), "Wrong number of standard PCI capabilities")
	require.Equal(t, 0, len(capabilities.Extended), "Wrong number of extended PCI capabilities")

	resource0 := devices[0].Resources[0]
	bar0, err := resource0.OpenRW()
	require.Nil(t, err, "Error opening bar0")
	defer func() {
		err := bar0.Close()
		if err != nil {
			t.Errorf("Error closing bar0: %v", err)
		}
	}()
	require.Equal(t, int(resource0.End-resource0.Start+1), bar0.Len())
	require.Equal(t, ga100PmcID, bar0.Read32(0))
}

func TestNvpciNUMANode(t *testing.T) {
	testCases := []struct {
		Description string
		NumaNode    int
	}{
		{
			Description: "Numa Node -1",
			NumaNode:    -1,
		},
		{
			Description: "Numa Node 0",
			NumaNode:    0,
		},
		{
			Description: "Numa Node 1",
			NumaNode:    1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			nvpci, err := NewMockNvpci()
			require.Nil(t, err, "Error creating NewMockNvpci")
			defer nvpci.Cleanup()

			err = nvpci.AddMockA100("0000:80:05.1", tc.NumaNode)
			require.Nil(t, err, "Error adding Mock A100 device to MockNvpci")

			devices, err := nvpci.GetGPUs()
			require.Nil(t, err, "Error getting GPUs")
			require.Equal(t, 1, len(devices), "Wrong number of GPU devices")
			require.Equal(t, tc.NumaNode, devices[0].NumaNode, "Wrong NUMA node found for device")
		})
	}
}
