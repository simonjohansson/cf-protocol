// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeShareServiceActor struct {
	ShareServiceInstanceInSpaceByOrganizationAndSpaceNameStub        func(serviceInstanceName string, sourceSpaceGUID string, sharedToOrgGUID string, sharedToSpaceName string) (v3action.Warnings, error)
	shareServiceInstanceInSpaceByOrganizationAndSpaceNameMutex       sync.RWMutex
	shareServiceInstanceInSpaceByOrganizationAndSpaceNameArgsForCall []struct {
		serviceInstanceName string
		sourceSpaceGUID     string
		sharedToOrgGUID     string
		sharedToSpaceName   string
	}
	shareServiceInstanceInSpaceByOrganizationAndSpaceNameReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	shareServiceInstanceInSpaceByOrganizationAndSpaceNameReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	ShareServiceInstanceInSpaceByOrganizationNameAndSpaceNameStub        func(serviceInstanceName string, sourceSpaceGUID string, sharedToOrgName string, sharedToSpaceName string) (v3action.Warnings, error)
	shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameMutex       sync.RWMutex
	shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameArgsForCall []struct {
		serviceInstanceName string
		sourceSpaceGUID     string
		sharedToOrgName     string
		sharedToSpaceName   string
	}
	shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeShareServiceActor) ShareServiceInstanceInSpaceByOrganizationAndSpaceName(serviceInstanceName string, sourceSpaceGUID string, sharedToOrgGUID string, sharedToSpaceName string) (v3action.Warnings, error) {
	fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameMutex.Lock()
	ret, specificReturn := fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameReturnsOnCall[len(fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameArgsForCall)]
	fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameArgsForCall = append(fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameArgsForCall, struct {
		serviceInstanceName string
		sourceSpaceGUID     string
		sharedToOrgGUID     string
		sharedToSpaceName   string
	}{serviceInstanceName, sourceSpaceGUID, sharedToOrgGUID, sharedToSpaceName})
	fake.recordInvocation("ShareServiceInstanceInSpaceByOrganizationAndSpaceName", []interface{}{serviceInstanceName, sourceSpaceGUID, sharedToOrgGUID, sharedToSpaceName})
	fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameMutex.Unlock()
	if fake.ShareServiceInstanceInSpaceByOrganizationAndSpaceNameStub != nil {
		return fake.ShareServiceInstanceInSpaceByOrganizationAndSpaceNameStub(serviceInstanceName, sourceSpaceGUID, sharedToOrgGUID, sharedToSpaceName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameReturns.result1, fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameReturns.result2
}

func (fake *FakeShareServiceActor) ShareServiceInstanceInSpaceByOrganizationAndSpaceNameCallCount() int {
	fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameMutex.RLock()
	defer fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameMutex.RUnlock()
	return len(fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameArgsForCall)
}

func (fake *FakeShareServiceActor) ShareServiceInstanceInSpaceByOrganizationAndSpaceNameArgsForCall(i int) (string, string, string, string) {
	fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameMutex.RLock()
	defer fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameMutex.RUnlock()
	return fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameArgsForCall[i].serviceInstanceName, fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameArgsForCall[i].sourceSpaceGUID, fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameArgsForCall[i].sharedToOrgGUID, fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameArgsForCall[i].sharedToSpaceName
}

func (fake *FakeShareServiceActor) ShareServiceInstanceInSpaceByOrganizationAndSpaceNameReturns(result1 v3action.Warnings, result2 error) {
	fake.ShareServiceInstanceInSpaceByOrganizationAndSpaceNameStub = nil
	fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) ShareServiceInstanceInSpaceByOrganizationAndSpaceNameReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.ShareServiceInstanceInSpaceByOrganizationAndSpaceNameStub = nil
	if fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameReturnsOnCall == nil {
		fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) ShareServiceInstanceInSpaceByOrganizationNameAndSpaceName(serviceInstanceName string, sourceSpaceGUID string, sharedToOrgName string, sharedToSpaceName string) (v3action.Warnings, error) {
	fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameMutex.Lock()
	ret, specificReturn := fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturnsOnCall[len(fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameArgsForCall)]
	fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameArgsForCall = append(fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameArgsForCall, struct {
		serviceInstanceName string
		sourceSpaceGUID     string
		sharedToOrgName     string
		sharedToSpaceName   string
	}{serviceInstanceName, sourceSpaceGUID, sharedToOrgName, sharedToSpaceName})
	fake.recordInvocation("ShareServiceInstanceInSpaceByOrganizationNameAndSpaceName", []interface{}{serviceInstanceName, sourceSpaceGUID, sharedToOrgName, sharedToSpaceName})
	fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameMutex.Unlock()
	if fake.ShareServiceInstanceInSpaceByOrganizationNameAndSpaceNameStub != nil {
		return fake.ShareServiceInstanceInSpaceByOrganizationNameAndSpaceNameStub(serviceInstanceName, sourceSpaceGUID, sharedToOrgName, sharedToSpaceName)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturns.result1, fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturns.result2
}

func (fake *FakeShareServiceActor) ShareServiceInstanceInSpaceByOrganizationNameAndSpaceNameCallCount() int {
	fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameMutex.RLock()
	defer fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameMutex.RUnlock()
	return len(fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameArgsForCall)
}

func (fake *FakeShareServiceActor) ShareServiceInstanceInSpaceByOrganizationNameAndSpaceNameArgsForCall(i int) (string, string, string, string) {
	fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameMutex.RLock()
	defer fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameMutex.RUnlock()
	return fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameArgsForCall[i].serviceInstanceName, fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameArgsForCall[i].sourceSpaceGUID, fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameArgsForCall[i].sharedToOrgName, fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameArgsForCall[i].sharedToSpaceName
}

func (fake *FakeShareServiceActor) ShareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturns(result1 v3action.Warnings, result2 error) {
	fake.ShareServiceInstanceInSpaceByOrganizationNameAndSpaceNameStub = nil
	fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) ShareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.ShareServiceInstanceInSpaceByOrganizationNameAndSpaceNameStub = nil
	if fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturnsOnCall == nil {
		fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeShareServiceActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeShareServiceActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeShareServiceActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeShareServiceActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeShareServiceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameMutex.RLock()
	defer fake.shareServiceInstanceInSpaceByOrganizationAndSpaceNameMutex.RUnlock()
	fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameMutex.RLock()
	defer fake.shareServiceInstanceInSpaceByOrganizationNameAndSpaceNameMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeShareServiceActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v3.ShareServiceActor = new(FakeShareServiceActor)
