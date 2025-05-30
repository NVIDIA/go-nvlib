/**
# Copyright 2024 NVIDIA CORPORATION
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

// Interface provides the API to the info package.
type Interface interface {
	PlatformResolver
	PropertyExtractor
}

// PlatformResolver defines a function to resolve the current platform.
type PlatformResolver interface {
	ResolvePlatform() Platform
}

// PropertyExtractor provides a set of functions to query capabilities of the
// system.
//
//go:generate moq  -rm -fmt=goimports -out property-extractor_mock.go . PropertyExtractor
type PropertyExtractor interface {
	HasDXCore() (bool, string)
	HasNvml() (bool, string)
	HasTegraFiles() (bool, string)
	// Deprecated: Use HasTegraFiles instead.
	IsTegraSystem() (bool, string)
	// Deprecated: Use HasOnlyIntegratedGPUs
	UsesOnlyNVGPUModule() (bool, string)
	HasOnlyIntegratedGPUs() (bool, string)
}
