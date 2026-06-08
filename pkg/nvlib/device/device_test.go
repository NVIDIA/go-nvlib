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

package device

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/NVIDIA/go-nvml/pkg/nvml/mock"
)

func pciInfoWithBusID(busID string) nvml.PciInfo {
	var info nvml.PciInfo
	for i := 0; i < len(busID) && i < len(info.BusId); i++ {
		info.BusId[i] = int8(busID[i])
	}
	return info
}

func deviceWithPCIBusID(busID string) *device {
	return &device{
		Device: &mock.Device{
			GetPciInfoFunc: func() (nvml.PciInfo, nvml.Return) {
				return pciInfoWithBusID(busID), nvml.SUCCESS
			},
		},
	}
}

func TestGetPCIBusID(t *testing.T) {
	testCases := []struct {
		name          string
		busIDFromNVML string
		expected      string
	}{
		{
			// Typical legacy NVML 4-digit domain: must not strip "0000".
			name:          "four_digit_legacy_domain",
			busIDFromNVML: "0000:0A:00.0",
			expected:      "0000:0a:00.0",
		},
		{
			// Non-zero 4-digit domain: must not trim.
			name:          "nonzero_four_digit_domain",
			busIDFromNVML: "0001:03:00.0",
			expected:      "0001:03:00.0",
		},
		{
			// 8-digit domain 00000000: trim prefix "0000".
			name:          "eight_digit_domain_padded_with_zeros",
			busIDFromNVML: "00000000:0a:00.0",
			expected:      "0000:0a:00.0",
		},
		{
			// 8-digit domain does not match padded "0000xxxx": leave unchanged.
			name:          "eight_digit_domain_id",
			busIDFromNVML: "0001ABCD:03:00.0",
			expected:      "0001abcd:03:00.0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := deviceWithPCIBusID(tc.busIDFromNVML).GetPCIBusID()
			require.NoError(t, err)
			require.Equal(t, tc.expected, got)
		})
	}
}
