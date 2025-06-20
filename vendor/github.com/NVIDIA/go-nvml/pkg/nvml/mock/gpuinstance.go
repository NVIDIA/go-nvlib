// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"sync"
)

// Ensure, that GpuInstance does implement nvml.GpuInstance.
// If this is not the case, regenerate this file with moq.
var _ nvml.GpuInstance = &GpuInstance{}

// GpuInstance is a mock implementation of nvml.GpuInstance.
//
//	func TestSomethingThatUsesGpuInstance(t *testing.T) {
//
//		// make and configure a mocked nvml.GpuInstance
//		mockedGpuInstance := &GpuInstance{
//			CreateComputeInstanceFunc: func(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) (nvml.ComputeInstance, nvml.Return) {
//				panic("mock out the CreateComputeInstance method")
//			},
//			CreateComputeInstanceWithPlacementFunc: func(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo, computeInstancePlacement *nvml.ComputeInstancePlacement) (nvml.ComputeInstance, nvml.Return) {
//				panic("mock out the CreateComputeInstanceWithPlacement method")
//			},
//			DestroyFunc: func() nvml.Return {
//				panic("mock out the Destroy method")
//			},
//			GetActiveVgpusFunc: func() (nvml.ActiveVgpuInstanceInfo, nvml.Return) {
//				panic("mock out the GetActiveVgpus method")
//			},
//			GetComputeInstanceByIdFunc: func(n int) (nvml.ComputeInstance, nvml.Return) {
//				panic("mock out the GetComputeInstanceById method")
//			},
//			GetComputeInstancePossiblePlacementsFunc: func(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstancePlacement, nvml.Return) {
//				panic("mock out the GetComputeInstancePossiblePlacements method")
//			},
//			GetComputeInstanceProfileInfoFunc: func(n1 int, n2 int) (nvml.ComputeInstanceProfileInfo, nvml.Return) {
//				panic("mock out the GetComputeInstanceProfileInfo method")
//			},
//			GetComputeInstanceProfileInfoVFunc: func(n1 int, n2 int) nvml.ComputeInstanceProfileInfoHandler {
//				panic("mock out the GetComputeInstanceProfileInfoV method")
//			},
//			GetComputeInstanceRemainingCapacityFunc: func(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) (int, nvml.Return) {
//				panic("mock out the GetComputeInstanceRemainingCapacity method")
//			},
//			GetComputeInstancesFunc: func(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstance, nvml.Return) {
//				panic("mock out the GetComputeInstances method")
//			},
//			GetCreatableVgpusFunc: func() (nvml.VgpuTypeIdInfo, nvml.Return) {
//				panic("mock out the GetCreatableVgpus method")
//			},
//			GetInfoFunc: func() (nvml.GpuInstanceInfo, nvml.Return) {
//				panic("mock out the GetInfo method")
//			},
//			GetVgpuHeterogeneousModeFunc: func() (nvml.VgpuHeterogeneousMode, nvml.Return) {
//				panic("mock out the GetVgpuHeterogeneousMode method")
//			},
//			GetVgpuSchedulerLogFunc: func() (nvml.VgpuSchedulerLogInfo, nvml.Return) {
//				panic("mock out the GetVgpuSchedulerLog method")
//			},
//			GetVgpuSchedulerStateFunc: func() (nvml.VgpuSchedulerStateInfo, nvml.Return) {
//				panic("mock out the GetVgpuSchedulerState method")
//			},
//			GetVgpuTypeCreatablePlacementsFunc: func() (nvml.VgpuCreatablePlacementInfo, nvml.Return) {
//				panic("mock out the GetVgpuTypeCreatablePlacements method")
//			},
//			SetVgpuHeterogeneousModeFunc: func(vgpuHeterogeneousMode *nvml.VgpuHeterogeneousMode) nvml.Return {
//				panic("mock out the SetVgpuHeterogeneousMode method")
//			},
//			SetVgpuSchedulerStateFunc: func(vgpuSchedulerState *nvml.VgpuSchedulerState) nvml.Return {
//				panic("mock out the SetVgpuSchedulerState method")
//			},
//		}
//
//		// use mockedGpuInstance in code that requires nvml.GpuInstance
//		// and then make assertions.
//
//	}
type GpuInstance struct {
	// CreateComputeInstanceFunc mocks the CreateComputeInstance method.
	CreateComputeInstanceFunc func(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) (nvml.ComputeInstance, nvml.Return)

	// CreateComputeInstanceWithPlacementFunc mocks the CreateComputeInstanceWithPlacement method.
	CreateComputeInstanceWithPlacementFunc func(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo, computeInstancePlacement *nvml.ComputeInstancePlacement) (nvml.ComputeInstance, nvml.Return)

	// DestroyFunc mocks the Destroy method.
	DestroyFunc func() nvml.Return

	// GetActiveVgpusFunc mocks the GetActiveVgpus method.
	GetActiveVgpusFunc func() (nvml.ActiveVgpuInstanceInfo, nvml.Return)

	// GetComputeInstanceByIdFunc mocks the GetComputeInstanceById method.
	GetComputeInstanceByIdFunc func(n int) (nvml.ComputeInstance, nvml.Return)

	// GetComputeInstancePossiblePlacementsFunc mocks the GetComputeInstancePossiblePlacements method.
	GetComputeInstancePossiblePlacementsFunc func(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstancePlacement, nvml.Return)

	// GetComputeInstanceProfileInfoFunc mocks the GetComputeInstanceProfileInfo method.
	GetComputeInstanceProfileInfoFunc func(n1 int, n2 int) (nvml.ComputeInstanceProfileInfo, nvml.Return)

	// GetComputeInstanceProfileInfoVFunc mocks the GetComputeInstanceProfileInfoV method.
	GetComputeInstanceProfileInfoVFunc func(n1 int, n2 int) nvml.ComputeInstanceProfileInfoHandler

	// GetComputeInstanceRemainingCapacityFunc mocks the GetComputeInstanceRemainingCapacity method.
	GetComputeInstanceRemainingCapacityFunc func(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) (int, nvml.Return)

	// GetComputeInstancesFunc mocks the GetComputeInstances method.
	GetComputeInstancesFunc func(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstance, nvml.Return)

	// GetCreatableVgpusFunc mocks the GetCreatableVgpus method.
	GetCreatableVgpusFunc func() (nvml.VgpuTypeIdInfo, nvml.Return)

	// GetInfoFunc mocks the GetInfo method.
	GetInfoFunc func() (nvml.GpuInstanceInfo, nvml.Return)

	// GetVgpuHeterogeneousModeFunc mocks the GetVgpuHeterogeneousMode method.
	GetVgpuHeterogeneousModeFunc func() (nvml.VgpuHeterogeneousMode, nvml.Return)

	// GetVgpuSchedulerLogFunc mocks the GetVgpuSchedulerLog method.
	GetVgpuSchedulerLogFunc func() (nvml.VgpuSchedulerLogInfo, nvml.Return)

	// GetVgpuSchedulerStateFunc mocks the GetVgpuSchedulerState method.
	GetVgpuSchedulerStateFunc func() (nvml.VgpuSchedulerStateInfo, nvml.Return)

	// GetVgpuTypeCreatablePlacementsFunc mocks the GetVgpuTypeCreatablePlacements method.
	GetVgpuTypeCreatablePlacementsFunc func() (nvml.VgpuCreatablePlacementInfo, nvml.Return)

	// SetVgpuHeterogeneousModeFunc mocks the SetVgpuHeterogeneousMode method.
	SetVgpuHeterogeneousModeFunc func(vgpuHeterogeneousMode *nvml.VgpuHeterogeneousMode) nvml.Return

	// SetVgpuSchedulerStateFunc mocks the SetVgpuSchedulerState method.
	SetVgpuSchedulerStateFunc func(vgpuSchedulerState *nvml.VgpuSchedulerState) nvml.Return

	// calls tracks calls to the methods.
	calls struct {
		// CreateComputeInstance holds details about calls to the CreateComputeInstance method.
		CreateComputeInstance []struct {
			// ComputeInstanceProfileInfo is the computeInstanceProfileInfo argument value.
			ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
		}
		// CreateComputeInstanceWithPlacement holds details about calls to the CreateComputeInstanceWithPlacement method.
		CreateComputeInstanceWithPlacement []struct {
			// ComputeInstanceProfileInfo is the computeInstanceProfileInfo argument value.
			ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
			// ComputeInstancePlacement is the computeInstancePlacement argument value.
			ComputeInstancePlacement *nvml.ComputeInstancePlacement
		}
		// Destroy holds details about calls to the Destroy method.
		Destroy []struct {
		}
		// GetActiveVgpus holds details about calls to the GetActiveVgpus method.
		GetActiveVgpus []struct {
		}
		// GetComputeInstanceById holds details about calls to the GetComputeInstanceById method.
		GetComputeInstanceById []struct {
			// N is the n argument value.
			N int
		}
		// GetComputeInstancePossiblePlacements holds details about calls to the GetComputeInstancePossiblePlacements method.
		GetComputeInstancePossiblePlacements []struct {
			// ComputeInstanceProfileInfo is the computeInstanceProfileInfo argument value.
			ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
		}
		// GetComputeInstanceProfileInfo holds details about calls to the GetComputeInstanceProfileInfo method.
		GetComputeInstanceProfileInfo []struct {
			// N1 is the n1 argument value.
			N1 int
			// N2 is the n2 argument value.
			N2 int
		}
		// GetComputeInstanceProfileInfoV holds details about calls to the GetComputeInstanceProfileInfoV method.
		GetComputeInstanceProfileInfoV []struct {
			// N1 is the n1 argument value.
			N1 int
			// N2 is the n2 argument value.
			N2 int
		}
		// GetComputeInstanceRemainingCapacity holds details about calls to the GetComputeInstanceRemainingCapacity method.
		GetComputeInstanceRemainingCapacity []struct {
			// ComputeInstanceProfileInfo is the computeInstanceProfileInfo argument value.
			ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
		}
		// GetComputeInstances holds details about calls to the GetComputeInstances method.
		GetComputeInstances []struct {
			// ComputeInstanceProfileInfo is the computeInstanceProfileInfo argument value.
			ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
		}
		// GetCreatableVgpus holds details about calls to the GetCreatableVgpus method.
		GetCreatableVgpus []struct {
		}
		// GetInfo holds details about calls to the GetInfo method.
		GetInfo []struct {
		}
		// GetVgpuHeterogeneousMode holds details about calls to the GetVgpuHeterogeneousMode method.
		GetVgpuHeterogeneousMode []struct {
		}
		// GetVgpuSchedulerLog holds details about calls to the GetVgpuSchedulerLog method.
		GetVgpuSchedulerLog []struct {
		}
		// GetVgpuSchedulerState holds details about calls to the GetVgpuSchedulerState method.
		GetVgpuSchedulerState []struct {
		}
		// GetVgpuTypeCreatablePlacements holds details about calls to the GetVgpuTypeCreatablePlacements method.
		GetVgpuTypeCreatablePlacements []struct {
		}
		// SetVgpuHeterogeneousMode holds details about calls to the SetVgpuHeterogeneousMode method.
		SetVgpuHeterogeneousMode []struct {
			// VgpuHeterogeneousMode is the vgpuHeterogeneousMode argument value.
			VgpuHeterogeneousMode *nvml.VgpuHeterogeneousMode
		}
		// SetVgpuSchedulerState holds details about calls to the SetVgpuSchedulerState method.
		SetVgpuSchedulerState []struct {
			// VgpuSchedulerState is the vgpuSchedulerState argument value.
			VgpuSchedulerState *nvml.VgpuSchedulerState
		}
	}
	lockCreateComputeInstance                sync.RWMutex
	lockCreateComputeInstanceWithPlacement   sync.RWMutex
	lockDestroy                              sync.RWMutex
	lockGetActiveVgpus                       sync.RWMutex
	lockGetComputeInstanceById               sync.RWMutex
	lockGetComputeInstancePossiblePlacements sync.RWMutex
	lockGetComputeInstanceProfileInfo        sync.RWMutex
	lockGetComputeInstanceProfileInfoV       sync.RWMutex
	lockGetComputeInstanceRemainingCapacity  sync.RWMutex
	lockGetComputeInstances                  sync.RWMutex
	lockGetCreatableVgpus                    sync.RWMutex
	lockGetInfo                              sync.RWMutex
	lockGetVgpuHeterogeneousMode             sync.RWMutex
	lockGetVgpuSchedulerLog                  sync.RWMutex
	lockGetVgpuSchedulerState                sync.RWMutex
	lockGetVgpuTypeCreatablePlacements       sync.RWMutex
	lockSetVgpuHeterogeneousMode             sync.RWMutex
	lockSetVgpuSchedulerState                sync.RWMutex
}

// CreateComputeInstance calls CreateComputeInstanceFunc.
func (mock *GpuInstance) CreateComputeInstance(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) (nvml.ComputeInstance, nvml.Return) {
	if mock.CreateComputeInstanceFunc == nil {
		panic("GpuInstance.CreateComputeInstanceFunc: method is nil but GpuInstance.CreateComputeInstance was just called")
	}
	callInfo := struct {
		ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
	}{
		ComputeInstanceProfileInfo: computeInstanceProfileInfo,
	}
	mock.lockCreateComputeInstance.Lock()
	mock.calls.CreateComputeInstance = append(mock.calls.CreateComputeInstance, callInfo)
	mock.lockCreateComputeInstance.Unlock()
	return mock.CreateComputeInstanceFunc(computeInstanceProfileInfo)
}

// CreateComputeInstanceCalls gets all the calls that were made to CreateComputeInstance.
// Check the length with:
//
//	len(mockedGpuInstance.CreateComputeInstanceCalls())
func (mock *GpuInstance) CreateComputeInstanceCalls() []struct {
	ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
} {
	var calls []struct {
		ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
	}
	mock.lockCreateComputeInstance.RLock()
	calls = mock.calls.CreateComputeInstance
	mock.lockCreateComputeInstance.RUnlock()
	return calls
}

// CreateComputeInstanceWithPlacement calls CreateComputeInstanceWithPlacementFunc.
func (mock *GpuInstance) CreateComputeInstanceWithPlacement(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo, computeInstancePlacement *nvml.ComputeInstancePlacement) (nvml.ComputeInstance, nvml.Return) {
	if mock.CreateComputeInstanceWithPlacementFunc == nil {
		panic("GpuInstance.CreateComputeInstanceWithPlacementFunc: method is nil but GpuInstance.CreateComputeInstanceWithPlacement was just called")
	}
	callInfo := struct {
		ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
		ComputeInstancePlacement   *nvml.ComputeInstancePlacement
	}{
		ComputeInstanceProfileInfo: computeInstanceProfileInfo,
		ComputeInstancePlacement:   computeInstancePlacement,
	}
	mock.lockCreateComputeInstanceWithPlacement.Lock()
	mock.calls.CreateComputeInstanceWithPlacement = append(mock.calls.CreateComputeInstanceWithPlacement, callInfo)
	mock.lockCreateComputeInstanceWithPlacement.Unlock()
	return mock.CreateComputeInstanceWithPlacementFunc(computeInstanceProfileInfo, computeInstancePlacement)
}

// CreateComputeInstanceWithPlacementCalls gets all the calls that were made to CreateComputeInstanceWithPlacement.
// Check the length with:
//
//	len(mockedGpuInstance.CreateComputeInstanceWithPlacementCalls())
func (mock *GpuInstance) CreateComputeInstanceWithPlacementCalls() []struct {
	ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
	ComputeInstancePlacement   *nvml.ComputeInstancePlacement
} {
	var calls []struct {
		ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
		ComputeInstancePlacement   *nvml.ComputeInstancePlacement
	}
	mock.lockCreateComputeInstanceWithPlacement.RLock()
	calls = mock.calls.CreateComputeInstanceWithPlacement
	mock.lockCreateComputeInstanceWithPlacement.RUnlock()
	return calls
}

// Destroy calls DestroyFunc.
func (mock *GpuInstance) Destroy() nvml.Return {
	if mock.DestroyFunc == nil {
		panic("GpuInstance.DestroyFunc: method is nil but GpuInstance.Destroy was just called")
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
//	len(mockedGpuInstance.DestroyCalls())
func (mock *GpuInstance) DestroyCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDestroy.RLock()
	calls = mock.calls.Destroy
	mock.lockDestroy.RUnlock()
	return calls
}

// GetActiveVgpus calls GetActiveVgpusFunc.
func (mock *GpuInstance) GetActiveVgpus() (nvml.ActiveVgpuInstanceInfo, nvml.Return) {
	if mock.GetActiveVgpusFunc == nil {
		panic("GpuInstance.GetActiveVgpusFunc: method is nil but GpuInstance.GetActiveVgpus was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetActiveVgpus.Lock()
	mock.calls.GetActiveVgpus = append(mock.calls.GetActiveVgpus, callInfo)
	mock.lockGetActiveVgpus.Unlock()
	return mock.GetActiveVgpusFunc()
}

// GetActiveVgpusCalls gets all the calls that were made to GetActiveVgpus.
// Check the length with:
//
//	len(mockedGpuInstance.GetActiveVgpusCalls())
func (mock *GpuInstance) GetActiveVgpusCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetActiveVgpus.RLock()
	calls = mock.calls.GetActiveVgpus
	mock.lockGetActiveVgpus.RUnlock()
	return calls
}

// GetComputeInstanceById calls GetComputeInstanceByIdFunc.
func (mock *GpuInstance) GetComputeInstanceById(n int) (nvml.ComputeInstance, nvml.Return) {
	if mock.GetComputeInstanceByIdFunc == nil {
		panic("GpuInstance.GetComputeInstanceByIdFunc: method is nil but GpuInstance.GetComputeInstanceById was just called")
	}
	callInfo := struct {
		N int
	}{
		N: n,
	}
	mock.lockGetComputeInstanceById.Lock()
	mock.calls.GetComputeInstanceById = append(mock.calls.GetComputeInstanceById, callInfo)
	mock.lockGetComputeInstanceById.Unlock()
	return mock.GetComputeInstanceByIdFunc(n)
}

// GetComputeInstanceByIdCalls gets all the calls that were made to GetComputeInstanceById.
// Check the length with:
//
//	len(mockedGpuInstance.GetComputeInstanceByIdCalls())
func (mock *GpuInstance) GetComputeInstanceByIdCalls() []struct {
	N int
} {
	var calls []struct {
		N int
	}
	mock.lockGetComputeInstanceById.RLock()
	calls = mock.calls.GetComputeInstanceById
	mock.lockGetComputeInstanceById.RUnlock()
	return calls
}

// GetComputeInstancePossiblePlacements calls GetComputeInstancePossiblePlacementsFunc.
func (mock *GpuInstance) GetComputeInstancePossiblePlacements(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstancePlacement, nvml.Return) {
	if mock.GetComputeInstancePossiblePlacementsFunc == nil {
		panic("GpuInstance.GetComputeInstancePossiblePlacementsFunc: method is nil but GpuInstance.GetComputeInstancePossiblePlacements was just called")
	}
	callInfo := struct {
		ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
	}{
		ComputeInstanceProfileInfo: computeInstanceProfileInfo,
	}
	mock.lockGetComputeInstancePossiblePlacements.Lock()
	mock.calls.GetComputeInstancePossiblePlacements = append(mock.calls.GetComputeInstancePossiblePlacements, callInfo)
	mock.lockGetComputeInstancePossiblePlacements.Unlock()
	return mock.GetComputeInstancePossiblePlacementsFunc(computeInstanceProfileInfo)
}

// GetComputeInstancePossiblePlacementsCalls gets all the calls that were made to GetComputeInstancePossiblePlacements.
// Check the length with:
//
//	len(mockedGpuInstance.GetComputeInstancePossiblePlacementsCalls())
func (mock *GpuInstance) GetComputeInstancePossiblePlacementsCalls() []struct {
	ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
} {
	var calls []struct {
		ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
	}
	mock.lockGetComputeInstancePossiblePlacements.RLock()
	calls = mock.calls.GetComputeInstancePossiblePlacements
	mock.lockGetComputeInstancePossiblePlacements.RUnlock()
	return calls
}

// GetComputeInstanceProfileInfo calls GetComputeInstanceProfileInfoFunc.
func (mock *GpuInstance) GetComputeInstanceProfileInfo(n1 int, n2 int) (nvml.ComputeInstanceProfileInfo, nvml.Return) {
	if mock.GetComputeInstanceProfileInfoFunc == nil {
		panic("GpuInstance.GetComputeInstanceProfileInfoFunc: method is nil but GpuInstance.GetComputeInstanceProfileInfo was just called")
	}
	callInfo := struct {
		N1 int
		N2 int
	}{
		N1: n1,
		N2: n2,
	}
	mock.lockGetComputeInstanceProfileInfo.Lock()
	mock.calls.GetComputeInstanceProfileInfo = append(mock.calls.GetComputeInstanceProfileInfo, callInfo)
	mock.lockGetComputeInstanceProfileInfo.Unlock()
	return mock.GetComputeInstanceProfileInfoFunc(n1, n2)
}

// GetComputeInstanceProfileInfoCalls gets all the calls that were made to GetComputeInstanceProfileInfo.
// Check the length with:
//
//	len(mockedGpuInstance.GetComputeInstanceProfileInfoCalls())
func (mock *GpuInstance) GetComputeInstanceProfileInfoCalls() []struct {
	N1 int
	N2 int
} {
	var calls []struct {
		N1 int
		N2 int
	}
	mock.lockGetComputeInstanceProfileInfo.RLock()
	calls = mock.calls.GetComputeInstanceProfileInfo
	mock.lockGetComputeInstanceProfileInfo.RUnlock()
	return calls
}

// GetComputeInstanceProfileInfoV calls GetComputeInstanceProfileInfoVFunc.
func (mock *GpuInstance) GetComputeInstanceProfileInfoV(n1 int, n2 int) nvml.ComputeInstanceProfileInfoHandler {
	if mock.GetComputeInstanceProfileInfoVFunc == nil {
		panic("GpuInstance.GetComputeInstanceProfileInfoVFunc: method is nil but GpuInstance.GetComputeInstanceProfileInfoV was just called")
	}
	callInfo := struct {
		N1 int
		N2 int
	}{
		N1: n1,
		N2: n2,
	}
	mock.lockGetComputeInstanceProfileInfoV.Lock()
	mock.calls.GetComputeInstanceProfileInfoV = append(mock.calls.GetComputeInstanceProfileInfoV, callInfo)
	mock.lockGetComputeInstanceProfileInfoV.Unlock()
	return mock.GetComputeInstanceProfileInfoVFunc(n1, n2)
}

// GetComputeInstanceProfileInfoVCalls gets all the calls that were made to GetComputeInstanceProfileInfoV.
// Check the length with:
//
//	len(mockedGpuInstance.GetComputeInstanceProfileInfoVCalls())
func (mock *GpuInstance) GetComputeInstanceProfileInfoVCalls() []struct {
	N1 int
	N2 int
} {
	var calls []struct {
		N1 int
		N2 int
	}
	mock.lockGetComputeInstanceProfileInfoV.RLock()
	calls = mock.calls.GetComputeInstanceProfileInfoV
	mock.lockGetComputeInstanceProfileInfoV.RUnlock()
	return calls
}

// GetComputeInstanceRemainingCapacity calls GetComputeInstanceRemainingCapacityFunc.
func (mock *GpuInstance) GetComputeInstanceRemainingCapacity(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) (int, nvml.Return) {
	if mock.GetComputeInstanceRemainingCapacityFunc == nil {
		panic("GpuInstance.GetComputeInstanceRemainingCapacityFunc: method is nil but GpuInstance.GetComputeInstanceRemainingCapacity was just called")
	}
	callInfo := struct {
		ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
	}{
		ComputeInstanceProfileInfo: computeInstanceProfileInfo,
	}
	mock.lockGetComputeInstanceRemainingCapacity.Lock()
	mock.calls.GetComputeInstanceRemainingCapacity = append(mock.calls.GetComputeInstanceRemainingCapacity, callInfo)
	mock.lockGetComputeInstanceRemainingCapacity.Unlock()
	return mock.GetComputeInstanceRemainingCapacityFunc(computeInstanceProfileInfo)
}

// GetComputeInstanceRemainingCapacityCalls gets all the calls that were made to GetComputeInstanceRemainingCapacity.
// Check the length with:
//
//	len(mockedGpuInstance.GetComputeInstanceRemainingCapacityCalls())
func (mock *GpuInstance) GetComputeInstanceRemainingCapacityCalls() []struct {
	ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
} {
	var calls []struct {
		ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
	}
	mock.lockGetComputeInstanceRemainingCapacity.RLock()
	calls = mock.calls.GetComputeInstanceRemainingCapacity
	mock.lockGetComputeInstanceRemainingCapacity.RUnlock()
	return calls
}

// GetComputeInstances calls GetComputeInstancesFunc.
func (mock *GpuInstance) GetComputeInstances(computeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo) ([]nvml.ComputeInstance, nvml.Return) {
	if mock.GetComputeInstancesFunc == nil {
		panic("GpuInstance.GetComputeInstancesFunc: method is nil but GpuInstance.GetComputeInstances was just called")
	}
	callInfo := struct {
		ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
	}{
		ComputeInstanceProfileInfo: computeInstanceProfileInfo,
	}
	mock.lockGetComputeInstances.Lock()
	mock.calls.GetComputeInstances = append(mock.calls.GetComputeInstances, callInfo)
	mock.lockGetComputeInstances.Unlock()
	return mock.GetComputeInstancesFunc(computeInstanceProfileInfo)
}

// GetComputeInstancesCalls gets all the calls that were made to GetComputeInstances.
// Check the length with:
//
//	len(mockedGpuInstance.GetComputeInstancesCalls())
func (mock *GpuInstance) GetComputeInstancesCalls() []struct {
	ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
} {
	var calls []struct {
		ComputeInstanceProfileInfo *nvml.ComputeInstanceProfileInfo
	}
	mock.lockGetComputeInstances.RLock()
	calls = mock.calls.GetComputeInstances
	mock.lockGetComputeInstances.RUnlock()
	return calls
}

// GetCreatableVgpus calls GetCreatableVgpusFunc.
func (mock *GpuInstance) GetCreatableVgpus() (nvml.VgpuTypeIdInfo, nvml.Return) {
	if mock.GetCreatableVgpusFunc == nil {
		panic("GpuInstance.GetCreatableVgpusFunc: method is nil but GpuInstance.GetCreatableVgpus was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetCreatableVgpus.Lock()
	mock.calls.GetCreatableVgpus = append(mock.calls.GetCreatableVgpus, callInfo)
	mock.lockGetCreatableVgpus.Unlock()
	return mock.GetCreatableVgpusFunc()
}

// GetCreatableVgpusCalls gets all the calls that were made to GetCreatableVgpus.
// Check the length with:
//
//	len(mockedGpuInstance.GetCreatableVgpusCalls())
func (mock *GpuInstance) GetCreatableVgpusCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetCreatableVgpus.RLock()
	calls = mock.calls.GetCreatableVgpus
	mock.lockGetCreatableVgpus.RUnlock()
	return calls
}

// GetInfo calls GetInfoFunc.
func (mock *GpuInstance) GetInfo() (nvml.GpuInstanceInfo, nvml.Return) {
	if mock.GetInfoFunc == nil {
		panic("GpuInstance.GetInfoFunc: method is nil but GpuInstance.GetInfo was just called")
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
//	len(mockedGpuInstance.GetInfoCalls())
func (mock *GpuInstance) GetInfoCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetInfo.RLock()
	calls = mock.calls.GetInfo
	mock.lockGetInfo.RUnlock()
	return calls
}

// GetVgpuHeterogeneousMode calls GetVgpuHeterogeneousModeFunc.
func (mock *GpuInstance) GetVgpuHeterogeneousMode() (nvml.VgpuHeterogeneousMode, nvml.Return) {
	if mock.GetVgpuHeterogeneousModeFunc == nil {
		panic("GpuInstance.GetVgpuHeterogeneousModeFunc: method is nil but GpuInstance.GetVgpuHeterogeneousMode was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetVgpuHeterogeneousMode.Lock()
	mock.calls.GetVgpuHeterogeneousMode = append(mock.calls.GetVgpuHeterogeneousMode, callInfo)
	mock.lockGetVgpuHeterogeneousMode.Unlock()
	return mock.GetVgpuHeterogeneousModeFunc()
}

// GetVgpuHeterogeneousModeCalls gets all the calls that were made to GetVgpuHeterogeneousMode.
// Check the length with:
//
//	len(mockedGpuInstance.GetVgpuHeterogeneousModeCalls())
func (mock *GpuInstance) GetVgpuHeterogeneousModeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetVgpuHeterogeneousMode.RLock()
	calls = mock.calls.GetVgpuHeterogeneousMode
	mock.lockGetVgpuHeterogeneousMode.RUnlock()
	return calls
}

// GetVgpuSchedulerLog calls GetVgpuSchedulerLogFunc.
func (mock *GpuInstance) GetVgpuSchedulerLog() (nvml.VgpuSchedulerLogInfo, nvml.Return) {
	if mock.GetVgpuSchedulerLogFunc == nil {
		panic("GpuInstance.GetVgpuSchedulerLogFunc: method is nil but GpuInstance.GetVgpuSchedulerLog was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetVgpuSchedulerLog.Lock()
	mock.calls.GetVgpuSchedulerLog = append(mock.calls.GetVgpuSchedulerLog, callInfo)
	mock.lockGetVgpuSchedulerLog.Unlock()
	return mock.GetVgpuSchedulerLogFunc()
}

// GetVgpuSchedulerLogCalls gets all the calls that were made to GetVgpuSchedulerLog.
// Check the length with:
//
//	len(mockedGpuInstance.GetVgpuSchedulerLogCalls())
func (mock *GpuInstance) GetVgpuSchedulerLogCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetVgpuSchedulerLog.RLock()
	calls = mock.calls.GetVgpuSchedulerLog
	mock.lockGetVgpuSchedulerLog.RUnlock()
	return calls
}

// GetVgpuSchedulerState calls GetVgpuSchedulerStateFunc.
func (mock *GpuInstance) GetVgpuSchedulerState() (nvml.VgpuSchedulerStateInfo, nvml.Return) {
	if mock.GetVgpuSchedulerStateFunc == nil {
		panic("GpuInstance.GetVgpuSchedulerStateFunc: method is nil but GpuInstance.GetVgpuSchedulerState was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetVgpuSchedulerState.Lock()
	mock.calls.GetVgpuSchedulerState = append(mock.calls.GetVgpuSchedulerState, callInfo)
	mock.lockGetVgpuSchedulerState.Unlock()
	return mock.GetVgpuSchedulerStateFunc()
}

// GetVgpuSchedulerStateCalls gets all the calls that were made to GetVgpuSchedulerState.
// Check the length with:
//
//	len(mockedGpuInstance.GetVgpuSchedulerStateCalls())
func (mock *GpuInstance) GetVgpuSchedulerStateCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetVgpuSchedulerState.RLock()
	calls = mock.calls.GetVgpuSchedulerState
	mock.lockGetVgpuSchedulerState.RUnlock()
	return calls
}

// GetVgpuTypeCreatablePlacements calls GetVgpuTypeCreatablePlacementsFunc.
func (mock *GpuInstance) GetVgpuTypeCreatablePlacements() (nvml.VgpuCreatablePlacementInfo, nvml.Return) {
	if mock.GetVgpuTypeCreatablePlacementsFunc == nil {
		panic("GpuInstance.GetVgpuTypeCreatablePlacementsFunc: method is nil but GpuInstance.GetVgpuTypeCreatablePlacements was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetVgpuTypeCreatablePlacements.Lock()
	mock.calls.GetVgpuTypeCreatablePlacements = append(mock.calls.GetVgpuTypeCreatablePlacements, callInfo)
	mock.lockGetVgpuTypeCreatablePlacements.Unlock()
	return mock.GetVgpuTypeCreatablePlacementsFunc()
}

// GetVgpuTypeCreatablePlacementsCalls gets all the calls that were made to GetVgpuTypeCreatablePlacements.
// Check the length with:
//
//	len(mockedGpuInstance.GetVgpuTypeCreatablePlacementsCalls())
func (mock *GpuInstance) GetVgpuTypeCreatablePlacementsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetVgpuTypeCreatablePlacements.RLock()
	calls = mock.calls.GetVgpuTypeCreatablePlacements
	mock.lockGetVgpuTypeCreatablePlacements.RUnlock()
	return calls
}

// SetVgpuHeterogeneousMode calls SetVgpuHeterogeneousModeFunc.
func (mock *GpuInstance) SetVgpuHeterogeneousMode(vgpuHeterogeneousMode *nvml.VgpuHeterogeneousMode) nvml.Return {
	if mock.SetVgpuHeterogeneousModeFunc == nil {
		panic("GpuInstance.SetVgpuHeterogeneousModeFunc: method is nil but GpuInstance.SetVgpuHeterogeneousMode was just called")
	}
	callInfo := struct {
		VgpuHeterogeneousMode *nvml.VgpuHeterogeneousMode
	}{
		VgpuHeterogeneousMode: vgpuHeterogeneousMode,
	}
	mock.lockSetVgpuHeterogeneousMode.Lock()
	mock.calls.SetVgpuHeterogeneousMode = append(mock.calls.SetVgpuHeterogeneousMode, callInfo)
	mock.lockSetVgpuHeterogeneousMode.Unlock()
	return mock.SetVgpuHeterogeneousModeFunc(vgpuHeterogeneousMode)
}

// SetVgpuHeterogeneousModeCalls gets all the calls that were made to SetVgpuHeterogeneousMode.
// Check the length with:
//
//	len(mockedGpuInstance.SetVgpuHeterogeneousModeCalls())
func (mock *GpuInstance) SetVgpuHeterogeneousModeCalls() []struct {
	VgpuHeterogeneousMode *nvml.VgpuHeterogeneousMode
} {
	var calls []struct {
		VgpuHeterogeneousMode *nvml.VgpuHeterogeneousMode
	}
	mock.lockSetVgpuHeterogeneousMode.RLock()
	calls = mock.calls.SetVgpuHeterogeneousMode
	mock.lockSetVgpuHeterogeneousMode.RUnlock()
	return calls
}

// SetVgpuSchedulerState calls SetVgpuSchedulerStateFunc.
func (mock *GpuInstance) SetVgpuSchedulerState(vgpuSchedulerState *nvml.VgpuSchedulerState) nvml.Return {
	if mock.SetVgpuSchedulerStateFunc == nil {
		panic("GpuInstance.SetVgpuSchedulerStateFunc: method is nil but GpuInstance.SetVgpuSchedulerState was just called")
	}
	callInfo := struct {
		VgpuSchedulerState *nvml.VgpuSchedulerState
	}{
		VgpuSchedulerState: vgpuSchedulerState,
	}
	mock.lockSetVgpuSchedulerState.Lock()
	mock.calls.SetVgpuSchedulerState = append(mock.calls.SetVgpuSchedulerState, callInfo)
	mock.lockSetVgpuSchedulerState.Unlock()
	return mock.SetVgpuSchedulerStateFunc(vgpuSchedulerState)
}

// SetVgpuSchedulerStateCalls gets all the calls that were made to SetVgpuSchedulerState.
// Check the length with:
//
//	len(mockedGpuInstance.SetVgpuSchedulerStateCalls())
func (mock *GpuInstance) SetVgpuSchedulerStateCalls() []struct {
	VgpuSchedulerState *nvml.VgpuSchedulerState
} {
	var calls []struct {
		VgpuSchedulerState *nvml.VgpuSchedulerState
	}
	mock.lockSetVgpuSchedulerState.RLock()
	calls = mock.calls.SetVgpuSchedulerState
	mock.lockSetVgpuSchedulerState.RUnlock()
	return calls
}
