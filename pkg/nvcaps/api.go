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

package nvcaps

// Interface provides the API to the 'nvcaps' package
type Interface interface {
	NewMigCapabilities() (Capabilities, error)
}

type nvcapslib struct {
	procRoot    string
	devRoot     string
	deviceMajor int
}

// Capability is a data structure that stores the properties of a capability device
type Capability struct {
	ProcPath    string
	DevicePath  string
	DeviceMajor int
	DeviceMinor int
}

// Capabilities defines the operations that be executed on a set of Capabilities
// TODO: Does an interface make sense here?
type Capabilities interface {
}

// New creates a new nvcaps interface with the supplied options
func New(opts ...Option) (Interface, error) {
	a := &nvcapslib{}
	for _, opt := range opts {
		opt(a)
	}

	return a, nil
}

// WithProcRoot sets the root for generated proc paths
func WithProcRoot(root string) Option {
	return func(a *nvcapslib) {
		a.procRoot = root
	}
}

// WithDevRoot sets the root for generated device paths
func WithDevRoot(root string) Option {
	return func(a *nvcapslib) {
		a.devRoot = root
	}
}

// WithDeviceMajor sets the device major number for cap devices
func WithDeviceMajor(major int) Option {
	return func(a *nvcapslib) {
		a.deviceMajor = major
	}
}

// Option defines a function for passing options to the New() call
type Option func(*nvcapslib)
