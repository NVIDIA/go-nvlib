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
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsGpuIndex(t *testing.T) {
	testCases := []struct {
		id       string
		expected bool
	}{
		{"", false},
		{"-1", false},
		{"0", true},
		{"1", true},
		{"not an integer", false},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			actual := Identifier(tc.id).IsGpuIndex()
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsMigIndex(t *testing.T) {
	testCases := []struct {
		id       string
		expected bool
	}{
		{"", false},
		{"0", false},
		{"not an integer", false},
		{"0:0", true},
		{"0:0:0", false},
		{"0:0.0", false},
		{"-1:0", false},
		{"0:-1", false},
		{"0:foo", false},
		{"foo:0", false},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			actual := Identifier(tc.id).IsMigIndex()
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsGpuUUID(t *testing.T) {
	testCases := []struct {
		id       string
		expected bool
	}{
		{"", false},
		{"0", false},
		{"not an integer", false},
		{"GPU-foo", false},
		{"GPU-ebd34bdf-1083-eaac-2aff-4b71a022f9bd", true},
		{"MIG-ebd34bdf-1083-eaac-2aff-4b71a022f9bd", false},
		{"ebd34bdf-1083-eaac-2aff-4b71a022f9bd", false},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			actual := Identifier(tc.id).IsGpuUUID()
			require.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsMigUUID(t *testing.T) {
	testCases := []struct {
		id       string
		expected bool
	}{
		{"", false},
		{"0", false},
		{"not an integer", false},
		{"MIG-foo", false},
		{"MIG-ebd34bdf-1083-eaac-2aff-4b71a022f9bd", true},
		{"GPU-ebd34bdf-1083-eaac-2aff-4b71a022f9bd", false},
		{"ebd34bdf-1083-eaac-2aff-4b71a022f9bd", false},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			actual := Identifier(tc.id).IsMigUUID()
			require.Equal(t, tc.expected, actual)
		})
	}
}
