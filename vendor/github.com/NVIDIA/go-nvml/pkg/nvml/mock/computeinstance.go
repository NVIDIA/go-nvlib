// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"sync"
)

// Ensure, that ComputeInstance does implement nvml.ComputeInstance.
// If this is not the case, regenerate this file with moq.
var _ nvml.ComputeInstance = &ComputeInstance{}

// ComputeInstance is a mock implementation of nvml.ComputeInstance.
//
//	func TestSomethingThatUsesComputeInstance(t *testing.T) {
//
//		// make and configure a mocked nvml.ComputeInstance
//		mockedComputeInstance := &ComputeInstance{
//			DestroyFunc: func() nvml.Return {
//				panic("mock out the Destroy method")
//			},
//			GetInfoFunc: func() (nvml.ComputeInstanceInfo, nvml.Return) {
//				panic("mock out the GetInfo method")
//			},
//		}
//
//		// use mockedComputeInstance in code that requires nvml.ComputeInstance
//		// and then make assertions.
//
//	}
type ComputeInstance struct {
	// DestroyFunc mocks the Destroy method.
	DestroyFunc func() nvml.Return

	// GetInfoFunc mocks the GetInfo method.
	GetInfoFunc func() (nvml.ComputeInstanceInfo, nvml.Return)

	// calls tracks calls to the methods.
	calls struct {
		// Destroy holds details about calls to the Destroy method.
		Destroy []struct {
		}
		// GetInfo holds details about calls to the GetInfo method.
		GetInfo []struct {
		}
	}
	lockDestroy sync.RWMutex
	lockGetInfo sync.RWMutex
}

// Destroy calls DestroyFunc.
func (mock *ComputeInstance) Destroy() nvml.Return {
	if mock.DestroyFunc == nil {
		panic("ComputeInstance.DestroyFunc: method is nil but ComputeInstance.Destroy was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDestroy.Lock()
	mock.calls.Destroy = append(mock.calls.Destroy, callInfo)
	mock.lockDestroy.Unlock()
	return mock.DestroyFunc()
}

// DestroyCalls gets all the calls that were made to Destroy.
// Check the length with:
//
//	len(mockedComputeInstance.DestroyCalls())
func (mock *ComputeInstance) DestroyCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDestroy.RLock()
	calls = mock.calls.Destroy
	mock.lockDestroy.RUnlock()
	return calls
}

// GetInfo calls GetInfoFunc.
func (mock *ComputeInstance) GetInfo() (nvml.ComputeInstanceInfo, nvml.Return) {
	if mock.GetInfoFunc == nil {
		panic("ComputeInstance.GetInfoFunc: method is nil but ComputeInstance.GetInfo was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetInfo.Lock()
	mock.calls.GetInfo = append(mock.calls.GetInfo, callInfo)
	mock.lockGetInfo.Unlock()
	return mock.GetInfoFunc()
}

// GetInfoCalls gets all the calls that were made to GetInfo.
// Check the length with:
//
//	len(mockedComputeInstance.GetInfoCalls())
func (mock *ComputeInstance) GetInfoCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetInfo.RLock()
	calls = mock.calls.GetInfo
	mock.lockGetInfo.RUnlock()
	return calls
}
