/**
# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResolvePlatform(t *testing.T) {
	testCases := []struct {
		platform           string
		hasTegraFiles      bool
		hasDXCore          bool
		hasNVML            bool
		hasAnIntegratedGPU bool
		expected           string
	}{
		{
			platform:  "auto",
			hasDXCore: true,
			expected:  "wsl",
		},
		{
			platform:      "auto",
			hasDXCore:     false,
			hasTegraFiles: true,
			hasNVML:       false,
			expected:      "tegra",
		},
		{
			platform:      "auto",
			hasDXCore:     false,
			hasTegraFiles: false,
			hasNVML:       false,
			expected:      "unknown",
		},
		{
			platform:      "auto",
			hasDXCore:     false,
			hasTegraFiles: true,
			hasNVML:       true,
			expected:      "nvml",
		},
		{
			platform:           "auto",
			hasDXCore:          false,
			hasTegraFiles:      true,
			hasNVML:            true,
			hasAnIntegratedGPU: true,
			expected:           "tegra",
		},
		{
			platform:      "nvml",
			hasDXCore:     true,
			hasTegraFiles: true,
			expected:      "nvml",
		},
		{
			platform:  "wsl",
			hasDXCore: false,
			expected:  "wsl",
		},
		{
			platform:  "not-auto",
			hasDXCore: true,
			expected:  "not-auto",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			l := New(
				WithPropertyExtractor(&PropertyExtractorMock{
					HasDXCoreFunc: func() (bool, string) {
						return tc.hasDXCore, ""
					},
					HasNvmlFunc: func() (bool, string) {
						return tc.hasNVML, ""
					},
					HasTegraFilesFunc: func() (bool, string) {
						return tc.hasTegraFiles, ""
					},
					HasAnIntegratedGPUFunc: func() (bool, string) {
						return tc.hasAnIntegratedGPU, ""
					},
				}),
				WithPlatform(Platform(tc.platform)),
			)

			require.Equal(t, Platform(tc.expected), l.ResolvePlatform())
		})
	}
}
