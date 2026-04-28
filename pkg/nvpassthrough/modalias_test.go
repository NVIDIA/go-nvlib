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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseModAliasString(t *testing.T) {
	testCases := []struct {
		description    string
		input          string
		expectedOutput *modAlias
		expectedError  bool
	}{
		{
			description:   "empty string",
			input:         "",
			expectedError: true,
		},
		{
			description:   "more than one semicolon delimiter",
			input:         "pci:foo:bar",
			expectedError: true,
		},
		{
			description: "all wildcards",
			input:       "pci:v*d*sv*sd*bc*sc*i*",
			expectedOutput: &modAlias{
				vendor:               "*",
				device:               "*",
				subvendor:            "*",
				subdevice:            "*",
				baseClass:            "*",
				subClass:             "*",
				programmingInterface: "*",
			},
		},
		{
			description: "some wildcards",
			input:       "pci:v000010DEd00002941sv*sd*bc*sc*i*",
			expectedOutput: &modAlias{
				vendor:               "000010DE",
				device:               "00002941",
				subvendor:            "*",
				subdevice:            "*",
				baseClass:            "*",
				subClass:             "*",
				programmingInterface: "*",
			},
		},
		{
			description: "no wildcards",
			input:       "pci:v000010DEd00002941sv000010DEsd00002046bc03sc02i00",
			expectedOutput: &modAlias{
				vendor:               "000010DE",
				device:               "00002941",
				subvendor:            "000010DE",
				subdevice:            "00002046",
				baseClass:            "03",
				subClass:             "02",
				programmingInterface: "00",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			modAlias, err := parseModAliasString(tc.input)
			if tc.expectedError {
				require.Error(t, err)
				return
			}
			require.NotNil(t, modAlias)
			require.EqualValues(t, *tc.expectedOutput, *modAlias)
		})
	}
}

func TestGetVFIOAliases(t *testing.T) {
	testCases := []struct {
		description    string
		input          string
		expectedOutput []vfioAlias
	}{
		{
			description:    "empty string",
			input:          "",
			expectedOutput: nil,
		},
		{
			description: "no vfio aliases",
			input: `
alias foo:v*d*sv*sd*bc*sc*i* bar
alias pci:v000010DEd00002941sv*sd*bc*sc*i* foo
`,
			expectedOutput: nil,
		},
		{
			description: "vfio aliases present",
			input: `
alias foo:v*d*sv*sd*bc*sc*i* bar
alias pci:v000010DEd00002941sv*sd*bc*sc*i* foo
alias vfio_pci:v*d*sv*sd*bc*sc*i* vfio_pci
alias vfio_pci:v000010DEd00002941sv*sd*bc*sc*i* nvgrace_gpu_vfio_pci
`,
			expectedOutput: []vfioAlias{
				{
					driver: "vfio_pci",
					modAlias: &modAlias{
						vendor:               "*",
						device:               "*",
						subvendor:            "*",
						subdevice:            "*",
						baseClass:            "*",
						subClass:             "*",
						programmingInterface: "*",
					},
				},
				{
					driver: "nvgrace_gpu_vfio_pci",
					modAlias: &modAlias{
						vendor:               "000010DE",
						device:               "00002941",
						subvendor:            "*",
						subdevice:            "*",
						baseClass:            "*",
						subClass:             "*",
						programmingInterface: "*",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			vfioAliases := getVFIOAliases(tc.input)
			require.EqualValues(t, tc.expectedOutput, vfioAliases)
		})
	}
}

func TestMatchModalias(t *testing.T) {
	testCases := []struct {
		description           string
		modalias              modAlias
		compareTo             modAlias
		expectedMatch         bool
		expectedWildcardCount int
	}{
		{
			description: "all wildcards",
			modalias: modAlias{
				vendor:               "000010DE",
				device:               "00002941",
				subvendor:            "000010DE",
				subdevice:            "00002046",
				baseClass:            "03",
				subClass:             "02",
				programmingInterface: "00",
			},
			compareTo: modAlias{
				vendor:               "*",
				device:               "*",
				subvendor:            "*",
				subdevice:            "*",
				baseClass:            "*",
				subClass:             "*",
				programmingInterface: "*",
			},
			expectedMatch:         true,
			expectedWildcardCount: 7,
		},
		{
			description: "some wildcards, match",
			modalias: modAlias{
				vendor:               "000010DE",
				device:               "00002941",
				subvendor:            "000010DE",
				subdevice:            "00002046",
				baseClass:            "03",
				subClass:             "02",
				programmingInterface: "00",
			},
			compareTo: modAlias{
				vendor:               "000010DE",
				device:               "00002941",
				subvendor:            "*",
				subdevice:            "*",
				baseClass:            "*",
				subClass:             "*",
				programmingInterface: "*",
			},
			expectedMatch:         true,
			expectedWildcardCount: 5,
		},
		{
			description: "some wildcards, not a match",
			modalias: modAlias{
				vendor:               "000010DE",
				device:               "00002941",
				subvendor:            "000010DE",
				subdevice:            "00002046",
				baseClass:            "03",
				subClass:             "02",
				programmingInterface: "00",
			},
			compareTo: modAlias{
				vendor:               "000010DE",
				device:               "00002900",
				subvendor:            "*",
				subdevice:            "*",
				baseClass:            "*",
				subClass:             "*",
				programmingInterface: "*",
			},
			expectedMatch: false,
		},
		{
			description: "no wildcards, match",
			modalias: modAlias{
				vendor:               "000010DE",
				device:               "00002941",
				subvendor:            "000010DE",
				subdevice:            "00002046",
				baseClass:            "03",
				subClass:             "02",
				programmingInterface: "00",
			},
			compareTo: modAlias{
				vendor:               "000010DE",
				device:               "00002941",
				subvendor:            "000010DE",
				subdevice:            "00002046",
				baseClass:            "03",
				subClass:             "02",
				programmingInterface: "00",
			},
			expectedMatch:         true,
			expectedWildcardCount: 0,
		},
		{
			description: "no wildcards, not a match",
			modalias: modAlias{
				vendor:               "000010DE",
				device:               "00002941",
				subvendor:            "000010DE",
				subdevice:            "00002046",
				baseClass:            "03",
				subClass:             "02",
				programmingInterface: "00",
			},
			compareTo: modAlias{
				vendor:               "00001111",
				device:               "00002222",
				subvendor:            "0000333",
				subdevice:            "00004444",
				baseClass:            "05",
				subClass:             "06",
				programmingInterface: "07",
			},
			expectedMatch: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			match, wildcardCount := matchModalias(&tc.modalias, &tc.compareTo)
			require.EqualValues(t, tc.expectedMatch, match)
			if tc.expectedMatch {
				require.EqualValues(t, tc.expectedWildcardCount, wildcardCount)
			}
		})
	}
}
