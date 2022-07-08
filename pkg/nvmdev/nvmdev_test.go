/*
 * Copyright (c) 2022, NVIDIA CORPORATION.  All rights reserved.
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

package nvmdev

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNvmdev(t *testing.T) {
	nvmdev, err := NewMock()
	require.Nil(t, err, "Error creating MockNvmdev")
	defer nvmdev.Cleanup()

	err = nvmdev.AddMockA100Parent("0000:3b:04.1", 0)
	require.Nil(t, err, "Error adding Mock A100 parent device to MockNvmdev")
	parentDevs, err := nvmdev.GetAllParentDevices()
	require.Nil(t, err, "Error getting parent GPU devices")
	require.Equal(t, 1, len(parentDevs), "Wrong number of parent GPU devices")

	parentA100 := parentDevs[0]
	supported := parentA100.IsMDEVTypeSupported("A100-4C")
	require.True(t, supported, "A100-4C should be a supported vGPU type")

	available, err := parentA100.IsMDEVTypeAvailable("A100-4C")
	require.Nil(t, err, "Error checking if A100-4Q vGPU type is available for creation")
	require.True(t, available, "A100-4C should be available to create")

	err = nvmdev.AddMockA100Mdev("b1914f0a-15cf-416e-8967-55fc7cb68e20", "A100-4C", "nvidia-500", parentDevs[0].Path)
	require.Nil(t, err, "Error adding Mock A100 mediated device")

	mdevs, err := nvmdev.GetAllDevices()
	require.Nil(t, err, "Error getting NVIDIA MDEV (vGPU) devices")
	require.Equal(t, 1, len(mdevs), "Wrong number of NVIDIA MDEV (vGPU) devices")
	require.Equal(t, "vfio_mdev", mdevs[0].Driver, "Wrong driver detected for mdev device")
}
