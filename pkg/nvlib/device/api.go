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

	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvml"
	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvpci"
)

// Interface provides the API to the 'device' package
type Interface interface {
	GetDevices() ([]Device, error)
	GetMigDevices() ([]MigDevice, error)
	GetMigProfiles() ([]MigProfile, error)
	NewDevice(d nvml.Device) (Device, error)
	NewMigDevice(d nvml.Device) (MigDevice, error)
	NewMigProfile(giProfileID, ciProfileID, ciEngProfileID int, migMemorySizeMB, deviceMemorySizeBytes uint64) (MigProfile, error)
	ParseMigProfile(profile string) (MigProfile, error)
	VisitDevices(func(i int, d Device) error) error
	VisitMigDevices(func(i int, d Device, j int, m MigDevice) error) error
	VisitMigProfiles(func(p MigProfile) error) error
}

type devicelib struct {
	nvml                  nvml.Interface
	selectedDeviceClasses map[Class]struct{}
}

var _ Interface = &devicelib{}

// New creates a new instance of the 'device' interface
func New(opts ...Option) (Interface, error) {
	d := &devicelib{}
	for _, opt := range opts {
		opt(d)
	}
	if d.selectedDeviceClasses == nil {
		option, err := defaultSelectedDeviceClasses()
		if err != nil {
			return nil, fmt.Errorf("error setting default selected device classes: %v", err)
		}
		option(d)
	}
	if d.nvml == nil {
		d.nvml = nvml.New()
	}

	return d, nil
}

// WithNvml provides an Option to set the NVML library used by the 'device' interface
func WithNvml(nvml nvml.Interface) Option {
	return func(d *devicelib) {
		d.nvml = nvml
	}
}

// WithSelectedDeviceClasses selects the specified device classes when filtering devices
func WithSelectedDeviceClasses(classes ...Class) Option {
	return func(d *devicelib) {
		if d.selectedDeviceClasses == nil {
			d.selectedDeviceClasses = make(map[Class]struct{})
		}
		for _, c := range classes {
			d.selectedDeviceClasses[c] = struct{}{}
		}
	}
}

// Option defines a function for passing options to the New() call
type Option func(*devicelib)

// defaultSelectedDeviceClasses sets the default device class selection based on the devices included.
func defaultSelectedDeviceClasses() (Option, error) {
	gpus, err := nvpci.New().GetGPUs()
	if err != nil {
		return nil, fmt.Errorf("error getting PCI devices: %v", err)
	}

	classes := make(map[Class]struct{})
	for _, gpu := range gpus {
		class := Class(gpu.Class)
		classes[class] = struct{}{}
	}

	var option Option
	if len(classes) == 1 {
		option = func(d *devicelib) {
		}
	} else {
		option = WithSelectedDeviceClasses(ClassCompute)
	}

	return option, nil
}
