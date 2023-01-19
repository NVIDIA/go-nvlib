/**
# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
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

package nvcaps

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	nvidiaProcDriverPath   = "/proc/driver/nvidia"
	nvidiaCapabilitiesPath = nvidiaProcDriverPath + "/capabilities"

	nvcapsProcDriverPath = "/proc/driver/nvidia-caps"
	nvcapsMigMinorsPath  = nvcapsProcDriverPath + "/mig-minors"
	nvcapsDevicePath     = "/dev/nvidia-caps"
)

// migCapabilities stores a map of MIG cap file paths to MIG minors
type migCapabilities struct {
	lib         *nvcapslib
	pathToMinor migCaps
}

type migMinor int
type migCap string
type migCaps map[migCap]migMinor

// NewMigCapabilities creates a
func (c *nvcapslib) NewMigCapabilities() (Capabilities, error) {
	migCaps, err := newMigCaps()
	if err != nil {
		return nil, fmt.Errorf("failed to construct MIG caps: %v", err)
	}

	caps := migCapabilities{
		lib:         c,
		pathToMinor: migCaps,
	}

	return &caps, nil
}

func (m *migCapabilities) GlobalMigMonitor() (Capability, error) {
	return m.getCapability(migCap("mig/monitor"))
}

func (m *migCapabilities) GlobalMigConfig() (Capability, error) {
	return m.getCapability(migCap("mig/config"))
}

func (m *migCapabilities) GPUInstanceAccess(gpu, gi int) (Capability, error) {
	cap := migCap(fmt.Sprintf("gpu%d/gi%d/access", gpu, gi))
	return m.getCapability(cap)
}

func (m *migCapabilities) ComputeInstanceAccess(gpu, gi, ci int) (Capability, error) {
	cap := migCap(fmt.Sprintf("gpu%d/gi%d/ci%d/access", gpu, gi, ci))
	return m.getCapability(cap)
}

func (m *migCapabilities) getCapability(cap migCap) (Capability, error) {
	minor, exists := m.pathToMinor[cap]
	if !exists {
		return Capability{}, fmt.Errorf("invalid MIG capability %v", cap)
	}

	c := Capability{
		ProcPath:    filepath.Join(m.lib.procRoot, cap.ProcPath()),
		DevicePath:  filepath.Join(m.lib.devRoot, minor.DevicePath()),
		DeviceMajor: m.lib.deviceMajor,
		DeviceMinor: int(minor),
	}

	return c, nil
}

// newMigCaps creates a MigCaps structure based on the contents of the MIG minors file.
func newMigCaps() (migCaps, error) {
	// Open nvcapsMigMinorsPath for walking.
	// If the nvcapsMigMinorsPath does not exist, then we are not on a MIG
	// capable machine, so there is nothing to do.
	// The format of this file is discussed in:
	//     https://docs.nvidia.com/datacenter/tesla/mig-user-guide/index.html#unique_1576522674
	minorsFile, err := os.Open(nvcapsMigMinorsPath)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error opening MIG minors file: %v", err)
	}
	defer minorsFile.Close()

	return processMinorsFile(minorsFile), nil
}

func processMinorsFile(minorsFile io.Reader) migCaps {
	// Walk each line of nvcapsMigMinorsPath and construct a mapping of nvidia
	// capabilities path to device minor for that capability
	migCaps := make(migCaps)
	scanner := bufio.NewScanner(minorsFile)
	for scanner.Scan() {
		cap, minor, err := processMigMinorsLine(scanner.Text())
		if err != nil {
			log.Printf("Skipping line in MIG minors file: %v", err)
			continue
		}
		migCaps[cap] = minor
	}
	return migCaps
}

func processMigMinorsLine(line string) (migCap, migMinor, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("error processing line: %v", line)
	}

	cap := migCap(parts[0])
	if !cap.isValid() {
		return "", 0, fmt.Errorf("invalid MIG minors line: '%v'", line)
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, fmt.Errorf("error reading MIG minor from '%v': %v", line, err)
	}

	return cap, migMinor(minor), nil
}

func (m migCap) isValid() bool {
	cap := string(m)
	switch cap {
	case "config", "monitor":
		return true
	default:
		var gpu int
		var gi int
		var ci int
		// Look for a CI access file
		n, _ := fmt.Sscanf(cap, "gpu%d/gi%d/ci%d/access", &gpu, &gi, &ci)
		if n == 3 {
			return true
		}
		// Look for a GI access file
		n, _ = fmt.Sscanf(cap, "gpu%d/gi%d/access %d", &gpu, &gi)
		if n == 2 {
			return true
		}
	}
	return false
}

// ProcPath returns the proc path associated with the MIG capability
func (m migCap) ProcPath() string {
	id := string(m)

	var path string
	switch id {
	case "config", "monitor":
		path = "mig/" + id
	default:
		parts := strings.SplitN(id, "/", 2)
		path = strings.Join([]string{parts[0], "mig", parts[1]}, "/")
	}
	return filepath.Join(nvidiaCapabilitiesPath, path)
}

// DevicePath returns the path for the nvidia-caps device with the specified
// minor number
func (m migMinor) DevicePath() string {
	return fmt.Sprintf(nvcapsDevicePath+"/nvidia-cap%d", m)
}
