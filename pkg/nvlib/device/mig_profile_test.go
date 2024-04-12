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

package device

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	nvmlmock "github.com/NVIDIA/go-nvml/pkg/nvml/mock"
)

type MigProfileInfoWrapper struct {
	MigProfileInfo
}

func newMockDeviceLib() Interface {
	mockDevice := &nvmlmock.Device{
		GetNameFunc: func() (string, nvml.Return) {
			return "MockDevice", nvml.SUCCESS
		},
		GetMigModeFunc: func() (int, int, nvml.Return) {
			return nvml.DEVICE_MIG_ENABLE, nvml.DEVICE_MIG_ENABLE, nvml.SUCCESS
		},
		GetMemoryInfoFunc: func() (nvml.Memory, nvml.Return) {
			memory := nvml.Memory{
				Total: 40 * 1024 * 1024 * 1024,
			}
			return memory, nvml.SUCCESS
		},
		GetGpuInstanceProfileInfoFunc: func(Profile int) (nvml.GpuInstanceProfileInfo, nvml.Return) {
			info := nvml.GpuInstanceProfileInfo{}
			switch Profile {
			case nvml.GPU_INSTANCE_PROFILE_1_SLICE,
				nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV1:
				info.MemorySizeMB = 5 * 1024
			case nvml.GPU_INSTANCE_PROFILE_1_SLICE_REV2:
				info.MemorySizeMB = 10 * 1024
			case nvml.GPU_INSTANCE_PROFILE_2_SLICE,
				nvml.GPU_INSTANCE_PROFILE_2_SLICE_REV1:
				info.MemorySizeMB = 10 * 1024
			case nvml.GPU_INSTANCE_PROFILE_3_SLICE:
				info.MemorySizeMB = 20 * 1024
			case nvml.GPU_INSTANCE_PROFILE_4_SLICE:
				info.MemorySizeMB = 20 * 1024
			case nvml.GPU_INSTANCE_PROFILE_7_SLICE:
				info.MemorySizeMB = 40 * 1024
			case nvml.GPU_INSTANCE_PROFILE_6_SLICE,
				nvml.GPU_INSTANCE_PROFILE_8_SLICE:
				fallthrough
			default:
				return info, nvml.ERROR_NOT_SUPPORTED
			}
			return info, nvml.SUCCESS
		},
	}
	mockNvml := &nvmlmock.Interface{
		DeviceGetCountFunc: func() (int, nvml.Return) {
			return 1, nvml.SUCCESS
		},
		DeviceGetHandleByIndexFunc: func(Index int) (nvml.Device, nvml.Return) {
			return mockDevice, nvml.SUCCESS
		},
	}

	return New(WithNvml(mockNvml), WithVerifySymbols(false))
}

func TestParseMigProfile(t *testing.T) {
	testCases := []struct {
		description string
		device      string
		validFormat bool
		validDevice bool
	}{
		{
			"Empty device type",
			"",
			false,
			false,
		},
		{
			"Valid 1g.5gb",
			"1g.5gb",
			true,
			true,
		},
		{
			"Valid 1c.1g.5gb",
			"1c.1g.5gb",
			true,
			true,
		},
		{
			"Valid 1g.5gb+me",
			"1g.5gb+me",
			true,
			true,
		},
		{
			"Valid 1c.1g.5gb+me",
			"1c.1g.5gb+me",
			true,
			true,
		},
		{
			"Invalid 0g.0gb",
			"0g.0gb",
			true,
			false,
		},
		{
			"Invalid 0c.0g.0gb",
			"0c.0g.0gb",
			true,
			false,
		},
		{
			"Invalid 10000g.500000gb",
			"10000g.500000gb",
			true,
			false,
		},
		{
			"Invalid 10000c.10000g.500000gb",
			"10000c.10000g.500000gb",
			true,
			false,
		},
		{
			"Invalid ' 1c.1g.5gb'",
			" 1c.1g.5gb",
			false,
			false,
		},
		{
			"Invalid '1 c.1g.5gb'",
			"1 c.1g.5gb",
			false,
			false,
		},
		{
			"Invalid '1c .1g.5gb'",
			"1c .1g.5gb",
			false,
			false,
		},
		{
			"Invalid '1c. 1g.5gb'",
			"1c. 1g.5gb",
			false,
			false,
		},
		{
			"Invalid '1c.1 g.5gb'",
			"1c.1 g.5gb",
			false,
			false,
		},
		{
			"Invalid '1c.1g .5gb'",
			"1c.1g .5gb",
			false,
			false,
		},
		{
			"Invalid '1c.1g. 5gb'",
			"1c.1g. 5gb",
			false,
			false,
		},
		{
			"Invalid '1c.1g.5 gb'",
			"1c.1g.5 gb",
			false,
			false,
		},
		{
			"Invalid '1c.1g.5g b'",
			"1c.1g.5g b",
			false,
			false,
		},
		{
			"Invalid '1c.1g.5gb '",
			"1c.1g.5gb ",
			false,
			false,
		},
		{
			"Invalid '1c . 1g . 5gb'",
			"1c . 1g . 5gb",
			false,
			false,
		},
		{
			"Invalid 1c.f1g.5gb",
			"1c.f1g.5gb",
			false,
			false,
		},
		{
			"Invalid 1r.1g.5gb",
			"1r.1g.5gb",
			false,
			false,
		},
		{
			"Invalid 1g.5gbk",
			"1g.5gbk",
			false,
			false,
		},
		{
			"Invalid 1g.5",
			"1g.5",
			false,
			false,
		},
		{
			"Invalid g.5gb",
			"g.5gb",
			false,
			false,
		},
		{
			"Invalid 1g.gb",
			"1g.gb",
			false,
			false,
		},
		{
			"Invalid 1g.5gb+me,me",
			"1g.5gb+me,me",
			false,
			false,
		},
		{
			"Invalid 1g.5gb+me,you,them",
			"1g.5gb+me,you,them",
			true,
			false,
		},
		{
			"Invalid 1c.1g.5gb+me,you,them",
			"1c.1g.5gb+me,you,them",
			true,
			false,
		},
		{
			"Invalid 1g.5gb+",
			"1g.5gb+",
			false,
			false,
		},
		{
			"Invalid 1g.5gb +",
			"1g.5gb +",
			false,
			false,
		},
		{
			"Invalid 1g.5gb+ ",
			"1g.5gb+ ",
			false,
			false,
		},
		{
			"Invalid 1g.5gb+ ,",
			"1g.5gb+ ,",
			false,
			false,
		},
		{
			"Invalid 1g.5gb+,,",
			"1g.5gb+,,",
			false,
			false,
		},
		{
			"Invalid 1g.5gb+me,",
			"1g.5gb+me,",
			false,
			false,
		},
		{
			"Invalid 1g.5gb+me,,",
			"1g.5gb+me,,",
			false,
			false,
		},
		{
			"Invalid 1g.5gb+me, ",
			"1g.5gb+me, ",
			false,
			false,
		},
		{
			"Invalid 1g.5gb+2me",
			"1g.5gb+2me",
			false,
			false,
		},
		{
			"Inavlid 1g.5gb*me",
			"1g.5gb*me",
			false,
			false,
		},
		{
			"Invalid 1c.1g.5gb*me",
			"1c.1g.5gb*me",
			false,
			false,
		},
		{
			"Invalid 1g.5gb*me,you,them",
			"1g.5gb*me,you,them",
			false,
			false,
		},
		{
			"Invalid 1c.1g.5gb*me,you,them",
			"1c.1g.5gb*me,you,them",
			false,
			false,
		},
		{
			"Invalid bogus",
			"bogus",
			false,
			false,
		},
	}

	d := newMockDeviceLib()
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			err := d.AssertValidMigProfileFormat(tc.device)
			if tc.validFormat {
				require.Nil(t, err)
			} else {
				require.Error(t, err)
			}
			_, err = d.ParseMigProfile(tc.device)
			if tc.validDevice {
				require.Nil(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestParseMigProfileEquals(t *testing.T) {
	testCases := []struct {
		description string
		profile1    string
		profile2    string
		valid       bool
	}{
		{
			"Exactly equal",
			"1g.5gb",
			"1g.5gb",
			true,
		},
		{
			"Equal when expanded",
			"1c.1g.5gb",
			"1g.5gb",
			true,
		},
		{
			"Equal with attributes",
			"1g.5gb+me",
			"1g.5gb+me",
			true,
		},
		{
			"Not equal C slices",
			"1c.2g.10gb",
			"2c.2g.10gb",
			false,
		},
		{
			"Not equal G slices",
			"1c.1g.10gb",
			"1c.2g.10gb",
			false,
		},
		{
			"Not equal attributes",
			"1g.5gb",
			"1g.5gb+me",
			false,
		},
	}

	d := newMockDeviceLib()
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			p1, err := d.ParseMigProfile(tc.profile1)
			require.Nil(t, err)
			p2, err := d.ParseMigProfile(tc.profile2)
			require.Nil(t, err)
			wrapper := MigProfileInfoWrapper{p2.GetInfo()}
			if tc.valid {
				require.True(t, p1.Equals(p2))
				require.True(t, p1.Equals(wrapper))
			} else {
				require.False(t, p1.Equals(p2))
				require.False(t, p1.Equals(wrapper))
			}
		})
	}
}

func TestGetMigMemorySizeGB(t *testing.T) {
	type testCase struct {
		totalDeviceMemory    uint64
		migMemorySizeMB      uint64
		expectedMemorySizeGB uint64
	}

	const maxMemorySlices = 8
	const oneMB = uint64(1024 * 1024)
	const oneGB = uint64(1024 * 1024 * 1024)

	totalDeviceMemory := []uint64{
		24 * oneGB,
		40 * oneGB,
		80 * oneGB,
	}

	testCases := []testCase{}
	for _, tdm := range totalDeviceMemory {
		sliceSize := tdm / maxMemorySlices

		const stepSize = oneGB / 32
		for i := stepSize; i <= tdm; i += stepSize {
			tc := testCase{
				totalDeviceMemory: tdm,
				migMemorySizeMB:   i / oneMB,
			}
			for j := uint64(sliceSize); j <= tdm; j += sliceSize {
				if i <= j {
					tc.expectedMemorySizeGB = j / oneGB
					break
				}
			}
			testCases = append(testCases, tc)
		}
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.migMemorySizeMB), func(t *testing.T) {
			memorySizeGB := getMigMemorySizeGB(tc.totalDeviceMemory, tc.migMemorySizeMB)
			require.Equal(t, int(tc.expectedMemorySizeGB), int(memorySizeGB))
		})
	}
}
