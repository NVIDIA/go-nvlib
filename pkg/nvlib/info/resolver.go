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

import "github.com/go-logr/logr"

// Platform represents a supported plaform.
type Platform string

const (
	PlatformAuto    = Platform("auto")
	PlatformNVML    = Platform("nvml")
	PlatformTegra   = Platform("tegra")
	PlatformWSL     = Platform("wsl")
	PlatformUnknown = Platform("unknown")
)

type platformResolver struct {
	logger            logr.Logger
	platform          Platform
	propertyExtractor PropertyExtractor
}

func (p platformResolver) ResolvePlatform() Platform {
	if p.platform != PlatformAuto {
		p.logger.Info("Using requested platform", "platform", p.platform)
		return p.platform
	}

	hasDXCore, reason := p.propertyExtractor.HasDXCore()
	p.logger.V(4).Info("Is WSL-based system?", "hasDXhasDXCore", hasDXCore, "reason", reason)

	hasTegraFiles, reason := p.propertyExtractor.HasTegraFiles()
	p.logger.V(4).Info("Is Tegra-based system?", "hasTegraFiles", hasTegraFiles, "reason", reason)

	hasNVML, reason := p.propertyExtractor.HasNvml()
	p.logger.V(4).Info("Is NVML-based system?", "hasNVML", hasNVML, "reason", reason)

	hasAnIntegratedGPU, reason := p.propertyExtractor.HasAnIntegratedGPU()
	p.logger.V(4).Info("Has an integrated GPU?", "hasAnIntegratedGPU", hasAnIntegratedGPU, "reason", reason)

	switch {
	case hasDXCore:
		return PlatformWSL
	case (hasTegraFiles && !hasNVML), hasAnIntegratedGPU:
		return PlatformTegra
	case hasNVML:
		return PlatformNVML
	default:
		return PlatformUnknown
	}
}
