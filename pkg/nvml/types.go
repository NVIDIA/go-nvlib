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
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// Return defines an NVML return type
type Return nvml.Return

//go:generate moq -out nvml_mock.go . Interface
// Interface defines the functions implemented by an NVML library
type Interface interface {
	Init() Return
	Shutdown() Return
	DeviceGetCount() (int, Return)
	DeviceGetHandleByIndex(Index int) (Device, Return)
	DeviceGetHandleByUUID(UUID string) (Device, Return)
	SystemGetDriverVersion() (string, Return)
	ErrorString(r Return) string
}

//go:generate moq -out device_mock.go . Device
// Device defines the functions implemented by an NVML device
type Device interface {
	GetIndex() (int, Return)
	GetPciInfo() (PciInfo, Return)
	GetMemoryInfo() (Memory, Return)
	GetUUID() (string, Return)
	GetMinorNumber() (int, Return)
	IsMigDeviceHandle() (bool, Return)
	GetDeviceHandleFromMigDeviceHandle() (Device, Return)
	SetMigMode(Mode int) (Return, Return)
	GetMigMode() (int, int, Return)
	GetGpuInstanceProfileInfo(Profile int) (GpuInstanceProfileInfo, Return)
	GetGpuInstances(Info *GpuInstanceProfileInfo) ([]GpuInstance, Return)
	GetMaxMigDeviceCount() (int, Return)
	GetMigDeviceHandleByIndex(Index int) (Device, Return)
	GetGpuInstanceId() (int, Return)
	GetComputeInstanceId() (int, Return)
}

//go:generate moq -out gi_mock.go . GpuInstance
// GpuInstance defines the functions implemented by a GpuInstance
type GpuInstance interface {
	GetInfo() (GpuInstanceInfo, Return)
	GetComputeInstanceProfileInfo(Profile int, EngProfile int) (ComputeInstanceProfileInfo, Return)
	CreateComputeInstance(Info *ComputeInstanceProfileInfo) (ComputeInstance, Return)
	GetComputeInstances(Info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return)
	Destroy() Return
}

//go:generate moq -out ci_mock.go . ComputeInstance
// ComputeInstance defines the functions implemented by a ComputeInstance
type ComputeInstance interface {
	GetInfo() (ComputeInstanceInfo, Return)
	Destroy() Return
}

// GpuInstanceInfo holds info about a GPU Instance
type GpuInstanceInfo struct {
	Device    Device
	Id        uint32
	ProfileId uint32
	Placement GpuInstancePlacement
}

// ComputeInstanceInfo holds info about a Compute Instance
type ComputeInstanceInfo struct {
	Device      Device
	GpuInstance GpuInstance
	Id          uint32
	ProfileId   uint32
	Placement   ComputeInstancePlacement
}

// Memory holds info about GPU device memory
type Memory nvml.Memory

//PciInfo holds info about the PCI connections of a GPU dvice
type PciInfo nvml.PciInfo

// GpuInstanceProfileInfo holds info about a GPU Instance Profile
type GpuInstanceProfileInfo nvml.GpuInstanceProfileInfo

// GpuInstancePlacement holds placement info about a GPU Instance
type GpuInstancePlacement nvml.GpuInstancePlacement

// ComputeInstanceProfileInfo holds info about a Compute Instance Profile
type ComputeInstanceProfileInfo nvml.ComputeInstanceProfileInfo

// ComputeInstancePlacement holds placement info about a Compute Instance
type ComputeInstancePlacement nvml.ComputeInstancePlacement
