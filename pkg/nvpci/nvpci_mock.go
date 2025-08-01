// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package nvpci

import (
	"sync"
)

// Ensure, that InterfaceMock does implement Interface.
// If this is not the case, regenerate this file with moq.
var _ Interface = &InterfaceMock{}

// InterfaceMock is a mock implementation of Interface.
//
//	func TestSomethingThatUsesInterface(t *testing.T) {
//
//		// make and configure a mocked Interface
//		mockedInterface := &InterfaceMock{
//			Get3DControllersFunc: func() ([]*NvidiaPCIDevice, error) {
//				panic("mock out the Get3DControllers method")
//			},
//			GetAllDevicesFunc: func() ([]*NvidiaPCIDevice, error) {
//				panic("mock out the GetAllDevices method")
//			},
//			GetDPUsFunc: func() ([]*NvidiaPCIDevice, error) {
//				panic("mock out the GetDPUs method")
//			},
//			GetGPUByIndexFunc: func(n int) (*NvidiaPCIDevice, error) {
//				panic("mock out the GetGPUByIndex method")
//			},
//			GetGPUByPciBusIDFunc: func(s string) (*NvidiaPCIDevice, error) {
//				panic("mock out the GetGPUByPciBusID method")
//			},
//			GetGPUsFunc: func() ([]*NvidiaPCIDevice, error) {
//				panic("mock out the GetGPUs method")
//			},
//			GetNVSwitchesFunc: func() ([]*NvidiaPCIDevice, error) {
//				panic("mock out the GetNVSwitches method")
//			},
//			GetNetworkControllersFunc: func() ([]*NvidiaPCIDevice, error) {
//				panic("mock out the GetNetworkControllers method")
//			},
//			GetPciBridgesFunc: func() ([]*NvidiaPCIDevice, error) {
//				panic("mock out the GetPciBridges method")
//			},
//			GetVGAControllersFunc: func() ([]*NvidiaPCIDevice, error) {
//				panic("mock out the GetVGAControllers method")
//			},
//		}
//
//		// use mockedInterface in code that requires Interface
//		// and then make assertions.
//
//	}
type InterfaceMock struct {
	// Get3DControllersFunc mocks the Get3DControllers method.
	Get3DControllersFunc func() ([]*NvidiaPCIDevice, error)

	// GetAllDevicesFunc mocks the GetAllDevices method.
	GetAllDevicesFunc func() ([]*NvidiaPCIDevice, error)

	// GetDPUsFunc mocks the GetDPUs method.
	GetDPUsFunc func() ([]*NvidiaPCIDevice, error)

	// GetGPUByIndexFunc mocks the GetGPUByIndex method.
	GetGPUByIndexFunc func(n int) (*NvidiaPCIDevice, error)

	// GetGPUByPciBusIDFunc mocks the GetGPUByPciBusID method.
	GetGPUByPciBusIDFunc func(s string) (*NvidiaPCIDevice, error)

	// GetGPUsFunc mocks the GetGPUs method.
	GetGPUsFunc func() ([]*NvidiaPCIDevice, error)

	// GetNVSwitchesFunc mocks the GetNVSwitches method.
	GetNVSwitchesFunc func() ([]*NvidiaPCIDevice, error)

	// GetNetworkControllersFunc mocks the GetNetworkControllers method.
	GetNetworkControllersFunc func() ([]*NvidiaPCIDevice, error)

	// GetPciBridgesFunc mocks the GetPciBridges method.
	GetPciBridgesFunc func() ([]*NvidiaPCIDevice, error)

	// GetVGAControllersFunc mocks the GetVGAControllers method.
	GetVGAControllersFunc func() ([]*NvidiaPCIDevice, error)

	// calls tracks calls to the methods.
	calls struct {
		// Get3DControllers holds details about calls to the Get3DControllers method.
		Get3DControllers []struct {
		}
		// GetAllDevices holds details about calls to the GetAllDevices method.
		GetAllDevices []struct {
		}
		// GetDPUs holds details about calls to the GetDPUs method.
		GetDPUs []struct {
		}
		// GetGPUByIndex holds details about calls to the GetGPUByIndex method.
		GetGPUByIndex []struct {
			// N is the n argument value.
			N int
		}
		// GetGPUByPciBusID holds details about calls to the GetGPUByPciBusID method.
		GetGPUByPciBusID []struct {
			// S is the s argument value.
			S string
		}
		// GetGPUs holds details about calls to the GetGPUs method.
		GetGPUs []struct {
		}
		// GetNVSwitches holds details about calls to the GetNVSwitches method.
		GetNVSwitches []struct {
		}
		// GetNetworkControllers holds details about calls to the GetNetworkControllers method.
		GetNetworkControllers []struct {
		}
		// GetPciBridges holds details about calls to the GetPciBridges method.
		GetPciBridges []struct {
		}
		// GetVGAControllers holds details about calls to the GetVGAControllers method.
		GetVGAControllers []struct {
		}
	}
	lockGet3DControllers      sync.RWMutex
	lockGetAllDevices         sync.RWMutex
	lockGetDPUs               sync.RWMutex
	lockGetGPUByIndex         sync.RWMutex
	lockGetGPUByPciBusID      sync.RWMutex
	lockGetGPUs               sync.RWMutex
	lockGetNVSwitches         sync.RWMutex
	lockGetNetworkControllers sync.RWMutex
	lockGetPciBridges         sync.RWMutex
	lockGetVGAControllers     sync.RWMutex
}

// Get3DControllers calls Get3DControllersFunc.
func (mock *InterfaceMock) Get3DControllers() ([]*NvidiaPCIDevice, error) {
	if mock.Get3DControllersFunc == nil {
		panic("InterfaceMock.Get3DControllersFunc: method is nil but Interface.Get3DControllers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGet3DControllers.Lock()
	mock.calls.Get3DControllers = append(mock.calls.Get3DControllers, callInfo)
	mock.lockGet3DControllers.Unlock()
	return mock.Get3DControllersFunc()
}

// Get3DControllersCalls gets all the calls that were made to Get3DControllers.
// Check the length with:
//
//	len(mockedInterface.Get3DControllersCalls())
func (mock *InterfaceMock) Get3DControllersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGet3DControllers.RLock()
	calls = mock.calls.Get3DControllers
	mock.lockGet3DControllers.RUnlock()
	return calls
}

// GetAllDevices calls GetAllDevicesFunc.
func (mock *InterfaceMock) GetAllDevices() ([]*NvidiaPCIDevice, error) {
	if mock.GetAllDevicesFunc == nil {
		panic("InterfaceMock.GetAllDevicesFunc: method is nil but Interface.GetAllDevices was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetAllDevices.Lock()
	mock.calls.GetAllDevices = append(mock.calls.GetAllDevices, callInfo)
	mock.lockGetAllDevices.Unlock()
	return mock.GetAllDevicesFunc()
}

// GetAllDevicesCalls gets all the calls that were made to GetAllDevices.
// Check the length with:
//
//	len(mockedInterface.GetAllDevicesCalls())
func (mock *InterfaceMock) GetAllDevicesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetAllDevices.RLock()
	calls = mock.calls.GetAllDevices
	mock.lockGetAllDevices.RUnlock()
	return calls
}

// GetDPUs calls GetDPUsFunc.
func (mock *InterfaceMock) GetDPUs() ([]*NvidiaPCIDevice, error) {
	if mock.GetDPUsFunc == nil {
		panic("InterfaceMock.GetDPUsFunc: method is nil but Interface.GetDPUs was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetDPUs.Lock()
	mock.calls.GetDPUs = append(mock.calls.GetDPUs, callInfo)
	mock.lockGetDPUs.Unlock()
	return mock.GetDPUsFunc()
}

// GetDPUsCalls gets all the calls that were made to GetDPUs.
// Check the length with:
//
//	len(mockedInterface.GetDPUsCalls())
func (mock *InterfaceMock) GetDPUsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetDPUs.RLock()
	calls = mock.calls.GetDPUs
	mock.lockGetDPUs.RUnlock()
	return calls
}

// GetGPUByIndex calls GetGPUByIndexFunc.
func (mock *InterfaceMock) GetGPUByIndex(n int) (*NvidiaPCIDevice, error) {
	if mock.GetGPUByIndexFunc == nil {
		panic("InterfaceMock.GetGPUByIndexFunc: method is nil but Interface.GetGPUByIndex was just called")
	}
	callInfo := struct {
		N int
	}{
		N: n,
	}
	mock.lockGetGPUByIndex.Lock()
	mock.calls.GetGPUByIndex = append(mock.calls.GetGPUByIndex, callInfo)
	mock.lockGetGPUByIndex.Unlock()
	return mock.GetGPUByIndexFunc(n)
}

// GetGPUByIndexCalls gets all the calls that were made to GetGPUByIndex.
// Check the length with:
//
//	len(mockedInterface.GetGPUByIndexCalls())
func (mock *InterfaceMock) GetGPUByIndexCalls() []struct {
	N int
} {
	var calls []struct {
		N int
	}
	mock.lockGetGPUByIndex.RLock()
	calls = mock.calls.GetGPUByIndex
	mock.lockGetGPUByIndex.RUnlock()
	return calls
}

// GetGPUByPciBusID calls GetGPUByPciBusIDFunc.
func (mock *InterfaceMock) GetGPUByPciBusID(s string) (*NvidiaPCIDevice, error) {
	if mock.GetGPUByPciBusIDFunc == nil {
		panic("InterfaceMock.GetGPUByPciBusIDFunc: method is nil but Interface.GetGPUByPciBusID was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockGetGPUByPciBusID.Lock()
	mock.calls.GetGPUByPciBusID = append(mock.calls.GetGPUByPciBusID, callInfo)
	mock.lockGetGPUByPciBusID.Unlock()
	return mock.GetGPUByPciBusIDFunc(s)
}

// GetGPUByPciBusIDCalls gets all the calls that were made to GetGPUByPciBusID.
// Check the length with:
//
//	len(mockedInterface.GetGPUByPciBusIDCalls())
func (mock *InterfaceMock) GetGPUByPciBusIDCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockGetGPUByPciBusID.RLock()
	calls = mock.calls.GetGPUByPciBusID
	mock.lockGetGPUByPciBusID.RUnlock()
	return calls
}

// GetGPUs calls GetGPUsFunc.
func (mock *InterfaceMock) GetGPUs() ([]*NvidiaPCIDevice, error) {
	if mock.GetGPUsFunc == nil {
		panic("InterfaceMock.GetGPUsFunc: method is nil but Interface.GetGPUs was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetGPUs.Lock()
	mock.calls.GetGPUs = append(mock.calls.GetGPUs, callInfo)
	mock.lockGetGPUs.Unlock()
	return mock.GetGPUsFunc()
}

// GetGPUsCalls gets all the calls that were made to GetGPUs.
// Check the length with:
//
//	len(mockedInterface.GetGPUsCalls())
func (mock *InterfaceMock) GetGPUsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetGPUs.RLock()
	calls = mock.calls.GetGPUs
	mock.lockGetGPUs.RUnlock()
	return calls
}

// GetNVSwitches calls GetNVSwitchesFunc.
func (mock *InterfaceMock) GetNVSwitches() ([]*NvidiaPCIDevice, error) {
	if mock.GetNVSwitchesFunc == nil {
		panic("InterfaceMock.GetNVSwitchesFunc: method is nil but Interface.GetNVSwitches was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetNVSwitches.Lock()
	mock.calls.GetNVSwitches = append(mock.calls.GetNVSwitches, callInfo)
	mock.lockGetNVSwitches.Unlock()
	return mock.GetNVSwitchesFunc()
}

// GetNVSwitchesCalls gets all the calls that were made to GetNVSwitches.
// Check the length with:
//
//	len(mockedInterface.GetNVSwitchesCalls())
func (mock *InterfaceMock) GetNVSwitchesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetNVSwitches.RLock()
	calls = mock.calls.GetNVSwitches
	mock.lockGetNVSwitches.RUnlock()
	return calls
}

// GetNetworkControllers calls GetNetworkControllersFunc.
func (mock *InterfaceMock) GetNetworkControllers() ([]*NvidiaPCIDevice, error) {
	if mock.GetNetworkControllersFunc == nil {
		panic("InterfaceMock.GetNetworkControllersFunc: method is nil but Interface.GetNetworkControllers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetNetworkControllers.Lock()
	mock.calls.GetNetworkControllers = append(mock.calls.GetNetworkControllers, callInfo)
	mock.lockGetNetworkControllers.Unlock()
	return mock.GetNetworkControllersFunc()
}

// GetNetworkControllersCalls gets all the calls that were made to GetNetworkControllers.
// Check the length with:
//
//	len(mockedInterface.GetNetworkControllersCalls())
func (mock *InterfaceMock) GetNetworkControllersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetNetworkControllers.RLock()
	calls = mock.calls.GetNetworkControllers
	mock.lockGetNetworkControllers.RUnlock()
	return calls
}

// GetPciBridges calls GetPciBridgesFunc.
func (mock *InterfaceMock) GetPciBridges() ([]*NvidiaPCIDevice, error) {
	if mock.GetPciBridgesFunc == nil {
		panic("InterfaceMock.GetPciBridgesFunc: method is nil but Interface.GetPciBridges was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPciBridges.Lock()
	mock.calls.GetPciBridges = append(mock.calls.GetPciBridges, callInfo)
	mock.lockGetPciBridges.Unlock()
	return mock.GetPciBridgesFunc()
}

// GetPciBridgesCalls gets all the calls that were made to GetPciBridges.
// Check the length with:
//
//	len(mockedInterface.GetPciBridgesCalls())
func (mock *InterfaceMock) GetPciBridgesCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPciBridges.RLock()
	calls = mock.calls.GetPciBridges
	mock.lockGetPciBridges.RUnlock()
	return calls
}

// GetVGAControllers calls GetVGAControllersFunc.
func (mock *InterfaceMock) GetVGAControllers() ([]*NvidiaPCIDevice, error) {
	if mock.GetVGAControllersFunc == nil {
		panic("InterfaceMock.GetVGAControllersFunc: method is nil but Interface.GetVGAControllers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetVGAControllers.Lock()
	mock.calls.GetVGAControllers = append(mock.calls.GetVGAControllers, callInfo)
	mock.lockGetVGAControllers.Unlock()
	return mock.GetVGAControllersFunc()
}

// GetVGAControllersCalls gets all the calls that were made to GetVGAControllers.
// Check the length with:
//
//	len(mockedInterface.GetVGAControllersCalls())
func (mock *InterfaceMock) GetVGAControllersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetVGAControllers.RLock()
	calls = mock.calls.GetVGAControllers
	mock.lockGetVGAControllers.RUnlock()
	return calls
}
