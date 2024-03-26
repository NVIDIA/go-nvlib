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

import (
	"github.com/NVIDIA/go-nvlib/pkg/nvlib/device"
	"github.com/NVIDIA/go-nvlib/pkg/nvml"
)

type builder struct {
	root      string
	nvmllib   nvml.Interface
	devicelib device.Interface
}

// New creates a new instance of the 'info' Interface
func New(opts ...Option) Interface {
	b := &builder{}
	for _, opt := range opts {
		opt(b)
	}
	if b.root == "" {
		b.root = "/"
	}
	if b.nvmllib == nil {
		b.nvmllib = nvml.New()
	}
	if b.devicelib == nil {
		b.devicelib = device.New(device.WithNvml(b.nvmllib))
	}
	return b.build()
}

func (b *builder) build() Interface {
	return &infolib{
		root:      b.root,
		nvmllib:   b.nvmllib,
		devicelib: b.devicelib,
	}
}
