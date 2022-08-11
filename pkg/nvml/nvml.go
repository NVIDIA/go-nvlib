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

package nvml

import (
	"sync"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

type nvmlLib struct {
	sync.Mutex
	refcount int
}

var _ Interface = (*nvmlLib)(nil)

// New creates a new instance of the NVML Interface
func New() Interface {
	return &nvmlLib{}
}

// Init initializes an NVML Interface
func (n *nvmlLib) Init() Return {
	ret := nvml.Init()
	if ret != nvml.SUCCESS {
		return Return(ret)
	}

	n.Lock()
	defer n.Unlock()
	if n.refcount == 0 {
		errorStringFunc = nvml.ErrorString
	}
	n.refcount += 1

	return SUCCESS
}

// Shutdown shuts down an NVML Interface
func (n *nvmlLib) Shutdown() Return {
	ret := nvml.Shutdown()
	if ret != nvml.SUCCESS {
		return Return(ret)
	}

	n.Lock()
	defer n.Unlock()
	n.refcount -= 1
	if n.refcount == 0 {
		errorStringFunc = defaultErrorStringFunc
	}

	return SUCCESS
}

// DeviceGetCount returns the total number of GPU Devices
func (n *nvmlLib) DeviceGetCount() (int, Return) {
	c, r := nvml.DeviceGetCount()
	return c, Return(r)
}

// DeviceGetHandleByIndex returns a Device handle given its index
func (n *nvmlLib) DeviceGetHandleByIndex(index int) (Device, Return) {
	d, r := nvml.DeviceGetHandleByIndex(index)
	return nvmlDevice(d), Return(r)
}

// DeviceGetHandleByUUID returns a Device handle given its UUID
func (n *nvmlLib) DeviceGetHandleByUUID(uuid string) (Device, Return) {
	d, r := nvml.DeviceGetHandleByUUID(uuid)
	return nvmlDevice(d), Return(r)
}

// SystemGetDriverVersion returns the version of the installed NVIDIA driver
func (n *nvmlLib) SystemGetDriverVersion() (string, Return) {
	v, r := nvml.SystemGetDriverVersion()
	return v, Return(r)
}

// ErrorString returns the error string associated with a given return value
func (n *nvmlLib) ErrorString(ret Return) string {
	return nvml.ErrorString(nvml.Return(ret))
}
