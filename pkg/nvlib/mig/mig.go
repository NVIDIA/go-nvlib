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

package mig

import (
	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvml"
)

// Interface provides the API to the mig package
type Interface interface {
	NewProfile(giProfileID, ciProfileID, ciEngProfileID int, migMemorySizeMB, deviceMemorySizeBytes uint64) (Profile, error)
	ParseProfile(profile string) (Profile, error)
	NewDevice(d nvml.Device) (Device, error)
}

type miglib struct {
	nvml nvml.Interface
}

var _ Interface = &miglib{}

// New creates a new instance of the 'mig' interface
func New(opts ...Option) Interface {
	m := &miglib{}
	for _, opt := range opts {
		opt(m)
	}
	if m.nvml == nil {
		m.nvml = nvml.New()
	}
	return m
}

// WithNvml provides an Option to set the NVML library used by the 'mig' interface
func WithNvml(nvml nvml.Interface) Option {
	return func(m *miglib) {
		m.nvml = nvml
	}
}

// Option defines a function for passing options to the New() call
type Option func(*miglib)
