// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"autoscaler/db"
	"autoscaler/models"
	"database/sql"
	"sync"
)

type FakeBindingDB struct {
	CheckServiceBindingStub        func(string) bool
	checkServiceBindingMutex       sync.RWMutex
	checkServiceBindingArgsForCall []struct {
		arg1 string
	}
	checkServiceBindingReturns struct {
		result1 bool
	}
	checkServiceBindingReturnsOnCall map[int]struct {
		result1 bool
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	CountServiceInstancesInOrgStub        func(string) (int, error)
	countServiceInstancesInOrgMutex       sync.RWMutex
	countServiceInstancesInOrgArgsForCall []struct {
		arg1 string
	}
	countServiceInstancesInOrgReturns struct {
		result1 int
		result2 error
	}
	countServiceInstancesInOrgReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	CreateServiceBindingStub        func(string, string, string) error
	createServiceBindingMutex       sync.RWMutex
	createServiceBindingArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createServiceBindingReturns struct {
		result1 error
	}
	createServiceBindingReturnsOnCall map[int]struct {
		result1 error
	}
	CreateServiceInstanceStub        func(models.ServiceInstance) error
	createServiceInstanceMutex       sync.RWMutex
	createServiceInstanceArgsForCall []struct {
		arg1 models.ServiceInstance
	}
	createServiceInstanceReturns struct {
		result1 error
	}
	createServiceInstanceReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteServiceBindingStub        func(string) error
	deleteServiceBindingMutex       sync.RWMutex
	deleteServiceBindingArgsForCall []struct {
		arg1 string
	}
	deleteServiceBindingReturns struct {
		result1 error
	}
	deleteServiceBindingReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteServiceBindingByAppIdStub        func(string) error
	deleteServiceBindingByAppIdMutex       sync.RWMutex
	deleteServiceBindingByAppIdArgsForCall []struct {
		arg1 string
	}
	deleteServiceBindingByAppIdReturns struct {
		result1 error
	}
	deleteServiceBindingByAppIdReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteServiceInstanceStub        func(string) error
	deleteServiceInstanceMutex       sync.RWMutex
	deleteServiceInstanceArgsForCall []struct {
		arg1 string
	}
	deleteServiceInstanceReturns struct {
		result1 error
	}
	deleteServiceInstanceReturnsOnCall map[int]struct {
		result1 error
	}
	GetAppIdByBindingIdStub        func(string) (string, error)
	getAppIdByBindingIdMutex       sync.RWMutex
	getAppIdByBindingIdArgsForCall []struct {
		arg1 string
	}
	getAppIdByBindingIdReturns struct {
		result1 string
		result2 error
	}
	getAppIdByBindingIdReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetAppIdsByInstanceIdStub        func(string) ([]string, error)
	getAppIdsByInstanceIdMutex       sync.RWMutex
	getAppIdsByInstanceIdArgsForCall []struct {
		arg1 string
	}
	getAppIdsByInstanceIdReturns struct {
		result1 []string
		result2 error
	}
	getAppIdsByInstanceIdReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetBindingIdsByInstanceIdStub        func(string) ([]string, error)
	getBindingIdsByInstanceIdMutex       sync.RWMutex
	getBindingIdsByInstanceIdArgsForCall []struct {
		arg1 string
	}
	getBindingIdsByInstanceIdReturns struct {
		result1 []string
		result2 error
	}
	getBindingIdsByInstanceIdReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetDBStatusStub        func() sql.DBStats
	getDBStatusMutex       sync.RWMutex
	getDBStatusArgsForCall []struct {
	}
	getDBStatusReturns struct {
		result1 sql.DBStats
	}
	getDBStatusReturnsOnCall map[int]struct {
		result1 sql.DBStats
	}
	GetServiceInstanceStub        func(string) (*models.ServiceInstance, error)
	getServiceInstanceMutex       sync.RWMutex
	getServiceInstanceArgsForCall []struct {
		arg1 string
	}
	getServiceInstanceReturns struct {
		result1 *models.ServiceInstance
		result2 error
	}
	getServiceInstanceReturnsOnCall map[int]struct {
		result1 *models.ServiceInstance
		result2 error
	}
	GetServiceInstanceByAppIdStub        func(string) (*models.ServiceInstance, error)
	getServiceInstanceByAppIdMutex       sync.RWMutex
	getServiceInstanceByAppIdArgsForCall []struct {
		arg1 string
	}
	getServiceInstanceByAppIdReturns struct {
		result1 *models.ServiceInstance
		result2 error
	}
	getServiceInstanceByAppIdReturnsOnCall map[int]struct {
		result1 *models.ServiceInstance
		result2 error
	}
	UpdateServiceInstanceStub        func(models.ServiceInstance) error
	updateServiceInstanceMutex       sync.RWMutex
	updateServiceInstanceArgsForCall []struct {
		arg1 models.ServiceInstance
	}
	updateServiceInstanceReturns struct {
		result1 error
	}
	updateServiceInstanceReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBindingDB) CheckServiceBinding(arg1 string) bool {
	fake.checkServiceBindingMutex.Lock()
	ret, specificReturn := fake.checkServiceBindingReturnsOnCall[len(fake.checkServiceBindingArgsForCall)]
	fake.checkServiceBindingArgsForCall = append(fake.checkServiceBindingArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CheckServiceBindingStub
	fakeReturns := fake.checkServiceBindingReturns
	fake.recordInvocation("CheckServiceBinding", []interface{}{arg1})
	fake.checkServiceBindingMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBindingDB) CheckServiceBindingCallCount() int {
	fake.checkServiceBindingMutex.RLock()
	defer fake.checkServiceBindingMutex.RUnlock()
	return len(fake.checkServiceBindingArgsForCall)
}

func (fake *FakeBindingDB) CheckServiceBindingCalls(stub func(string) bool) {
	fake.checkServiceBindingMutex.Lock()
	defer fake.checkServiceBindingMutex.Unlock()
	fake.CheckServiceBindingStub = stub
}

func (fake *FakeBindingDB) CheckServiceBindingArgsForCall(i int) string {
	fake.checkServiceBindingMutex.RLock()
	defer fake.checkServiceBindingMutex.RUnlock()
	argsForCall := fake.checkServiceBindingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) CheckServiceBindingReturns(result1 bool) {
	fake.checkServiceBindingMutex.Lock()
	defer fake.checkServiceBindingMutex.Unlock()
	fake.CheckServiceBindingStub = nil
	fake.checkServiceBindingReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBindingDB) CheckServiceBindingReturnsOnCall(i int, result1 bool) {
	fake.checkServiceBindingMutex.Lock()
	defer fake.checkServiceBindingMutex.Unlock()
	fake.CheckServiceBindingStub = nil
	if fake.checkServiceBindingReturnsOnCall == nil {
		fake.checkServiceBindingReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.checkServiceBindingReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBindingDB) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBindingDB) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeBindingDB) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeBindingDB) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) CountServiceInstancesInOrg(arg1 string) (int, error) {
	fake.countServiceInstancesInOrgMutex.Lock()
	ret, specificReturn := fake.countServiceInstancesInOrgReturnsOnCall[len(fake.countServiceInstancesInOrgArgsForCall)]
	fake.countServiceInstancesInOrgArgsForCall = append(fake.countServiceInstancesInOrgArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CountServiceInstancesInOrgStub
	fakeReturns := fake.countServiceInstancesInOrgReturns
	fake.recordInvocation("CountServiceInstancesInOrg", []interface{}{arg1})
	fake.countServiceInstancesInOrgMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBindingDB) CountServiceInstancesInOrgCallCount() int {
	fake.countServiceInstancesInOrgMutex.RLock()
	defer fake.countServiceInstancesInOrgMutex.RUnlock()
	return len(fake.countServiceInstancesInOrgArgsForCall)
}

func (fake *FakeBindingDB) CountServiceInstancesInOrgCalls(stub func(string) (int, error)) {
	fake.countServiceInstancesInOrgMutex.Lock()
	defer fake.countServiceInstancesInOrgMutex.Unlock()
	fake.CountServiceInstancesInOrgStub = stub
}

func (fake *FakeBindingDB) CountServiceInstancesInOrgArgsForCall(i int) string {
	fake.countServiceInstancesInOrgMutex.RLock()
	defer fake.countServiceInstancesInOrgMutex.RUnlock()
	argsForCall := fake.countServiceInstancesInOrgArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) CountServiceInstancesInOrgReturns(result1 int, result2 error) {
	fake.countServiceInstancesInOrgMutex.Lock()
	defer fake.countServiceInstancesInOrgMutex.Unlock()
	fake.CountServiceInstancesInOrgStub = nil
	fake.countServiceInstancesInOrgReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) CountServiceInstancesInOrgReturnsOnCall(i int, result1 int, result2 error) {
	fake.countServiceInstancesInOrgMutex.Lock()
	defer fake.countServiceInstancesInOrgMutex.Unlock()
	fake.CountServiceInstancesInOrgStub = nil
	if fake.countServiceInstancesInOrgReturnsOnCall == nil {
		fake.countServiceInstancesInOrgReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.countServiceInstancesInOrgReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) CreateServiceBinding(arg1 string, arg2 string, arg3 string) error {
	fake.createServiceBindingMutex.Lock()
	ret, specificReturn := fake.createServiceBindingReturnsOnCall[len(fake.createServiceBindingArgsForCall)]
	fake.createServiceBindingArgsForCall = append(fake.createServiceBindingArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CreateServiceBindingStub
	fakeReturns := fake.createServiceBindingReturns
	fake.recordInvocation("CreateServiceBinding", []interface{}{arg1, arg2, arg3})
	fake.createServiceBindingMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBindingDB) CreateServiceBindingCallCount() int {
	fake.createServiceBindingMutex.RLock()
	defer fake.createServiceBindingMutex.RUnlock()
	return len(fake.createServiceBindingArgsForCall)
}

func (fake *FakeBindingDB) CreateServiceBindingCalls(stub func(string, string, string) error) {
	fake.createServiceBindingMutex.Lock()
	defer fake.createServiceBindingMutex.Unlock()
	fake.CreateServiceBindingStub = stub
}

func (fake *FakeBindingDB) CreateServiceBindingArgsForCall(i int) (string, string, string) {
	fake.createServiceBindingMutex.RLock()
	defer fake.createServiceBindingMutex.RUnlock()
	argsForCall := fake.createServiceBindingArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeBindingDB) CreateServiceBindingReturns(result1 error) {
	fake.createServiceBindingMutex.Lock()
	defer fake.createServiceBindingMutex.Unlock()
	fake.CreateServiceBindingStub = nil
	fake.createServiceBindingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) CreateServiceBindingReturnsOnCall(i int, result1 error) {
	fake.createServiceBindingMutex.Lock()
	defer fake.createServiceBindingMutex.Unlock()
	fake.CreateServiceBindingStub = nil
	if fake.createServiceBindingReturnsOnCall == nil {
		fake.createServiceBindingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createServiceBindingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) CreateServiceInstance(arg1 models.ServiceInstance) error {
	fake.createServiceInstanceMutex.Lock()
	ret, specificReturn := fake.createServiceInstanceReturnsOnCall[len(fake.createServiceInstanceArgsForCall)]
	fake.createServiceInstanceArgsForCall = append(fake.createServiceInstanceArgsForCall, struct {
		arg1 models.ServiceInstance
	}{arg1})
	stub := fake.CreateServiceInstanceStub
	fakeReturns := fake.createServiceInstanceReturns
	fake.recordInvocation("CreateServiceInstance", []interface{}{arg1})
	fake.createServiceInstanceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBindingDB) CreateServiceInstanceCallCount() int {
	fake.createServiceInstanceMutex.RLock()
	defer fake.createServiceInstanceMutex.RUnlock()
	return len(fake.createServiceInstanceArgsForCall)
}

func (fake *FakeBindingDB) CreateServiceInstanceCalls(stub func(models.ServiceInstance) error) {
	fake.createServiceInstanceMutex.Lock()
	defer fake.createServiceInstanceMutex.Unlock()
	fake.CreateServiceInstanceStub = stub
}

func (fake *FakeBindingDB) CreateServiceInstanceArgsForCall(i int) models.ServiceInstance {
	fake.createServiceInstanceMutex.RLock()
	defer fake.createServiceInstanceMutex.RUnlock()
	argsForCall := fake.createServiceInstanceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) CreateServiceInstanceReturns(result1 error) {
	fake.createServiceInstanceMutex.Lock()
	defer fake.createServiceInstanceMutex.Unlock()
	fake.CreateServiceInstanceStub = nil
	fake.createServiceInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) CreateServiceInstanceReturnsOnCall(i int, result1 error) {
	fake.createServiceInstanceMutex.Lock()
	defer fake.createServiceInstanceMutex.Unlock()
	fake.CreateServiceInstanceStub = nil
	if fake.createServiceInstanceReturnsOnCall == nil {
		fake.createServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createServiceInstanceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) DeleteServiceBinding(arg1 string) error {
	fake.deleteServiceBindingMutex.Lock()
	ret, specificReturn := fake.deleteServiceBindingReturnsOnCall[len(fake.deleteServiceBindingArgsForCall)]
	fake.deleteServiceBindingArgsForCall = append(fake.deleteServiceBindingArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteServiceBindingStub
	fakeReturns := fake.deleteServiceBindingReturns
	fake.recordInvocation("DeleteServiceBinding", []interface{}{arg1})
	fake.deleteServiceBindingMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBindingDB) DeleteServiceBindingCallCount() int {
	fake.deleteServiceBindingMutex.RLock()
	defer fake.deleteServiceBindingMutex.RUnlock()
	return len(fake.deleteServiceBindingArgsForCall)
}

func (fake *FakeBindingDB) DeleteServiceBindingCalls(stub func(string) error) {
	fake.deleteServiceBindingMutex.Lock()
	defer fake.deleteServiceBindingMutex.Unlock()
	fake.DeleteServiceBindingStub = stub
}

func (fake *FakeBindingDB) DeleteServiceBindingArgsForCall(i int) string {
	fake.deleteServiceBindingMutex.RLock()
	defer fake.deleteServiceBindingMutex.RUnlock()
	argsForCall := fake.deleteServiceBindingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) DeleteServiceBindingReturns(result1 error) {
	fake.deleteServiceBindingMutex.Lock()
	defer fake.deleteServiceBindingMutex.Unlock()
	fake.DeleteServiceBindingStub = nil
	fake.deleteServiceBindingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) DeleteServiceBindingReturnsOnCall(i int, result1 error) {
	fake.deleteServiceBindingMutex.Lock()
	defer fake.deleteServiceBindingMutex.Unlock()
	fake.DeleteServiceBindingStub = nil
	if fake.deleteServiceBindingReturnsOnCall == nil {
		fake.deleteServiceBindingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteServiceBindingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) DeleteServiceBindingByAppId(arg1 string) error {
	fake.deleteServiceBindingByAppIdMutex.Lock()
	ret, specificReturn := fake.deleteServiceBindingByAppIdReturnsOnCall[len(fake.deleteServiceBindingByAppIdArgsForCall)]
	fake.deleteServiceBindingByAppIdArgsForCall = append(fake.deleteServiceBindingByAppIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteServiceBindingByAppIdStub
	fakeReturns := fake.deleteServiceBindingByAppIdReturns
	fake.recordInvocation("DeleteServiceBindingByAppId", []interface{}{arg1})
	fake.deleteServiceBindingByAppIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBindingDB) DeleteServiceBindingByAppIdCallCount() int {
	fake.deleteServiceBindingByAppIdMutex.RLock()
	defer fake.deleteServiceBindingByAppIdMutex.RUnlock()
	return len(fake.deleteServiceBindingByAppIdArgsForCall)
}

func (fake *FakeBindingDB) DeleteServiceBindingByAppIdCalls(stub func(string) error) {
	fake.deleteServiceBindingByAppIdMutex.Lock()
	defer fake.deleteServiceBindingByAppIdMutex.Unlock()
	fake.DeleteServiceBindingByAppIdStub = stub
}

func (fake *FakeBindingDB) DeleteServiceBindingByAppIdArgsForCall(i int) string {
	fake.deleteServiceBindingByAppIdMutex.RLock()
	defer fake.deleteServiceBindingByAppIdMutex.RUnlock()
	argsForCall := fake.deleteServiceBindingByAppIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) DeleteServiceBindingByAppIdReturns(result1 error) {
	fake.deleteServiceBindingByAppIdMutex.Lock()
	defer fake.deleteServiceBindingByAppIdMutex.Unlock()
	fake.DeleteServiceBindingByAppIdStub = nil
	fake.deleteServiceBindingByAppIdReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) DeleteServiceBindingByAppIdReturnsOnCall(i int, result1 error) {
	fake.deleteServiceBindingByAppIdMutex.Lock()
	defer fake.deleteServiceBindingByAppIdMutex.Unlock()
	fake.DeleteServiceBindingByAppIdStub = nil
	if fake.deleteServiceBindingByAppIdReturnsOnCall == nil {
		fake.deleteServiceBindingByAppIdReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteServiceBindingByAppIdReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) DeleteServiceInstance(arg1 string) error {
	fake.deleteServiceInstanceMutex.Lock()
	ret, specificReturn := fake.deleteServiceInstanceReturnsOnCall[len(fake.deleteServiceInstanceArgsForCall)]
	fake.deleteServiceInstanceArgsForCall = append(fake.deleteServiceInstanceArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteServiceInstanceStub
	fakeReturns := fake.deleteServiceInstanceReturns
	fake.recordInvocation("DeleteServiceInstance", []interface{}{arg1})
	fake.deleteServiceInstanceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBindingDB) DeleteServiceInstanceCallCount() int {
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	return len(fake.deleteServiceInstanceArgsForCall)
}

func (fake *FakeBindingDB) DeleteServiceInstanceCalls(stub func(string) error) {
	fake.deleteServiceInstanceMutex.Lock()
	defer fake.deleteServiceInstanceMutex.Unlock()
	fake.DeleteServiceInstanceStub = stub
}

func (fake *FakeBindingDB) DeleteServiceInstanceArgsForCall(i int) string {
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	argsForCall := fake.deleteServiceInstanceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) DeleteServiceInstanceReturns(result1 error) {
	fake.deleteServiceInstanceMutex.Lock()
	defer fake.deleteServiceInstanceMutex.Unlock()
	fake.DeleteServiceInstanceStub = nil
	fake.deleteServiceInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) DeleteServiceInstanceReturnsOnCall(i int, result1 error) {
	fake.deleteServiceInstanceMutex.Lock()
	defer fake.deleteServiceInstanceMutex.Unlock()
	fake.DeleteServiceInstanceStub = nil
	if fake.deleteServiceInstanceReturnsOnCall == nil {
		fake.deleteServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteServiceInstanceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) GetAppIdByBindingId(arg1 string) (string, error) {
	fake.getAppIdByBindingIdMutex.Lock()
	ret, specificReturn := fake.getAppIdByBindingIdReturnsOnCall[len(fake.getAppIdByBindingIdArgsForCall)]
	fake.getAppIdByBindingIdArgsForCall = append(fake.getAppIdByBindingIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetAppIdByBindingIdStub
	fakeReturns := fake.getAppIdByBindingIdReturns
	fake.recordInvocation("GetAppIdByBindingId", []interface{}{arg1})
	fake.getAppIdByBindingIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBindingDB) GetAppIdByBindingIdCallCount() int {
	fake.getAppIdByBindingIdMutex.RLock()
	defer fake.getAppIdByBindingIdMutex.RUnlock()
	return len(fake.getAppIdByBindingIdArgsForCall)
}

func (fake *FakeBindingDB) GetAppIdByBindingIdCalls(stub func(string) (string, error)) {
	fake.getAppIdByBindingIdMutex.Lock()
	defer fake.getAppIdByBindingIdMutex.Unlock()
	fake.GetAppIdByBindingIdStub = stub
}

func (fake *FakeBindingDB) GetAppIdByBindingIdArgsForCall(i int) string {
	fake.getAppIdByBindingIdMutex.RLock()
	defer fake.getAppIdByBindingIdMutex.RUnlock()
	argsForCall := fake.getAppIdByBindingIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) GetAppIdByBindingIdReturns(result1 string, result2 error) {
	fake.getAppIdByBindingIdMutex.Lock()
	defer fake.getAppIdByBindingIdMutex.Unlock()
	fake.GetAppIdByBindingIdStub = nil
	fake.getAppIdByBindingIdReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) GetAppIdByBindingIdReturnsOnCall(i int, result1 string, result2 error) {
	fake.getAppIdByBindingIdMutex.Lock()
	defer fake.getAppIdByBindingIdMutex.Unlock()
	fake.GetAppIdByBindingIdStub = nil
	if fake.getAppIdByBindingIdReturnsOnCall == nil {
		fake.getAppIdByBindingIdReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getAppIdByBindingIdReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) GetAppIdsByInstanceId(arg1 string) ([]string, error) {
	fake.getAppIdsByInstanceIdMutex.Lock()
	ret, specificReturn := fake.getAppIdsByInstanceIdReturnsOnCall[len(fake.getAppIdsByInstanceIdArgsForCall)]
	fake.getAppIdsByInstanceIdArgsForCall = append(fake.getAppIdsByInstanceIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetAppIdsByInstanceIdStub
	fakeReturns := fake.getAppIdsByInstanceIdReturns
	fake.recordInvocation("GetAppIdsByInstanceId", []interface{}{arg1})
	fake.getAppIdsByInstanceIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBindingDB) GetAppIdsByInstanceIdCallCount() int {
	fake.getAppIdsByInstanceIdMutex.RLock()
	defer fake.getAppIdsByInstanceIdMutex.RUnlock()
	return len(fake.getAppIdsByInstanceIdArgsForCall)
}

func (fake *FakeBindingDB) GetAppIdsByInstanceIdCalls(stub func(string) ([]string, error)) {
	fake.getAppIdsByInstanceIdMutex.Lock()
	defer fake.getAppIdsByInstanceIdMutex.Unlock()
	fake.GetAppIdsByInstanceIdStub = stub
}

func (fake *FakeBindingDB) GetAppIdsByInstanceIdArgsForCall(i int) string {
	fake.getAppIdsByInstanceIdMutex.RLock()
	defer fake.getAppIdsByInstanceIdMutex.RUnlock()
	argsForCall := fake.getAppIdsByInstanceIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) GetAppIdsByInstanceIdReturns(result1 []string, result2 error) {
	fake.getAppIdsByInstanceIdMutex.Lock()
	defer fake.getAppIdsByInstanceIdMutex.Unlock()
	fake.GetAppIdsByInstanceIdStub = nil
	fake.getAppIdsByInstanceIdReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) GetAppIdsByInstanceIdReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getAppIdsByInstanceIdMutex.Lock()
	defer fake.getAppIdsByInstanceIdMutex.Unlock()
	fake.GetAppIdsByInstanceIdStub = nil
	if fake.getAppIdsByInstanceIdReturnsOnCall == nil {
		fake.getAppIdsByInstanceIdReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getAppIdsByInstanceIdReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) GetBindingIdsByInstanceId(arg1 string) ([]string, error) {
	fake.getBindingIdsByInstanceIdMutex.Lock()
	ret, specificReturn := fake.getBindingIdsByInstanceIdReturnsOnCall[len(fake.getBindingIdsByInstanceIdArgsForCall)]
	fake.getBindingIdsByInstanceIdArgsForCall = append(fake.getBindingIdsByInstanceIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetBindingIdsByInstanceIdStub
	fakeReturns := fake.getBindingIdsByInstanceIdReturns
	fake.recordInvocation("GetBindingIdsByInstanceId", []interface{}{arg1})
	fake.getBindingIdsByInstanceIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBindingDB) GetBindingIdsByInstanceIdCallCount() int {
	fake.getBindingIdsByInstanceIdMutex.RLock()
	defer fake.getBindingIdsByInstanceIdMutex.RUnlock()
	return len(fake.getBindingIdsByInstanceIdArgsForCall)
}

func (fake *FakeBindingDB) GetBindingIdsByInstanceIdCalls(stub func(string) ([]string, error)) {
	fake.getBindingIdsByInstanceIdMutex.Lock()
	defer fake.getBindingIdsByInstanceIdMutex.Unlock()
	fake.GetBindingIdsByInstanceIdStub = stub
}

func (fake *FakeBindingDB) GetBindingIdsByInstanceIdArgsForCall(i int) string {
	fake.getBindingIdsByInstanceIdMutex.RLock()
	defer fake.getBindingIdsByInstanceIdMutex.RUnlock()
	argsForCall := fake.getBindingIdsByInstanceIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) GetBindingIdsByInstanceIdReturns(result1 []string, result2 error) {
	fake.getBindingIdsByInstanceIdMutex.Lock()
	defer fake.getBindingIdsByInstanceIdMutex.Unlock()
	fake.GetBindingIdsByInstanceIdStub = nil
	fake.getBindingIdsByInstanceIdReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) GetBindingIdsByInstanceIdReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getBindingIdsByInstanceIdMutex.Lock()
	defer fake.getBindingIdsByInstanceIdMutex.Unlock()
	fake.GetBindingIdsByInstanceIdStub = nil
	if fake.getBindingIdsByInstanceIdReturnsOnCall == nil {
		fake.getBindingIdsByInstanceIdReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getBindingIdsByInstanceIdReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) GetDBStatus() sql.DBStats {
	fake.getDBStatusMutex.Lock()
	ret, specificReturn := fake.getDBStatusReturnsOnCall[len(fake.getDBStatusArgsForCall)]
	fake.getDBStatusArgsForCall = append(fake.getDBStatusArgsForCall, struct {
	}{})
	stub := fake.GetDBStatusStub
	fakeReturns := fake.getDBStatusReturns
	fake.recordInvocation("GetDBStatus", []interface{}{})
	fake.getDBStatusMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBindingDB) GetDBStatusCallCount() int {
	fake.getDBStatusMutex.RLock()
	defer fake.getDBStatusMutex.RUnlock()
	return len(fake.getDBStatusArgsForCall)
}

func (fake *FakeBindingDB) GetDBStatusCalls(stub func() sql.DBStats) {
	fake.getDBStatusMutex.Lock()
	defer fake.getDBStatusMutex.Unlock()
	fake.GetDBStatusStub = stub
}

func (fake *FakeBindingDB) GetDBStatusReturns(result1 sql.DBStats) {
	fake.getDBStatusMutex.Lock()
	defer fake.getDBStatusMutex.Unlock()
	fake.GetDBStatusStub = nil
	fake.getDBStatusReturns = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakeBindingDB) GetDBStatusReturnsOnCall(i int, result1 sql.DBStats) {
	fake.getDBStatusMutex.Lock()
	defer fake.getDBStatusMutex.Unlock()
	fake.GetDBStatusStub = nil
	if fake.getDBStatusReturnsOnCall == nil {
		fake.getDBStatusReturnsOnCall = make(map[int]struct {
			result1 sql.DBStats
		})
	}
	fake.getDBStatusReturnsOnCall[i] = struct {
		result1 sql.DBStats
	}{result1}
}

func (fake *FakeBindingDB) GetServiceInstance(arg1 string) (*models.ServiceInstance, error) {
	fake.getServiceInstanceMutex.Lock()
	ret, specificReturn := fake.getServiceInstanceReturnsOnCall[len(fake.getServiceInstanceArgsForCall)]
	fake.getServiceInstanceArgsForCall = append(fake.getServiceInstanceArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetServiceInstanceStub
	fakeReturns := fake.getServiceInstanceReturns
	fake.recordInvocation("GetServiceInstance", []interface{}{arg1})
	fake.getServiceInstanceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBindingDB) GetServiceInstanceCallCount() int {
	fake.getServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceMutex.RUnlock()
	return len(fake.getServiceInstanceArgsForCall)
}

func (fake *FakeBindingDB) GetServiceInstanceCalls(stub func(string) (*models.ServiceInstance, error)) {
	fake.getServiceInstanceMutex.Lock()
	defer fake.getServiceInstanceMutex.Unlock()
	fake.GetServiceInstanceStub = stub
}

func (fake *FakeBindingDB) GetServiceInstanceArgsForCall(i int) string {
	fake.getServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceMutex.RUnlock()
	argsForCall := fake.getServiceInstanceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) GetServiceInstanceReturns(result1 *models.ServiceInstance, result2 error) {
	fake.getServiceInstanceMutex.Lock()
	defer fake.getServiceInstanceMutex.Unlock()
	fake.GetServiceInstanceStub = nil
	fake.getServiceInstanceReturns = struct {
		result1 *models.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) GetServiceInstanceReturnsOnCall(i int, result1 *models.ServiceInstance, result2 error) {
	fake.getServiceInstanceMutex.Lock()
	defer fake.getServiceInstanceMutex.Unlock()
	fake.GetServiceInstanceStub = nil
	if fake.getServiceInstanceReturnsOnCall == nil {
		fake.getServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 *models.ServiceInstance
			result2 error
		})
	}
	fake.getServiceInstanceReturnsOnCall[i] = struct {
		result1 *models.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) GetServiceInstanceByAppId(arg1 string) (*models.ServiceInstance, error) {
	fake.getServiceInstanceByAppIdMutex.Lock()
	ret, specificReturn := fake.getServiceInstanceByAppIdReturnsOnCall[len(fake.getServiceInstanceByAppIdArgsForCall)]
	fake.getServiceInstanceByAppIdArgsForCall = append(fake.getServiceInstanceByAppIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetServiceInstanceByAppIdStub
	fakeReturns := fake.getServiceInstanceByAppIdReturns
	fake.recordInvocation("GetServiceInstanceByAppId", []interface{}{arg1})
	fake.getServiceInstanceByAppIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBindingDB) GetServiceInstanceByAppIdCallCount() int {
	fake.getServiceInstanceByAppIdMutex.RLock()
	defer fake.getServiceInstanceByAppIdMutex.RUnlock()
	return len(fake.getServiceInstanceByAppIdArgsForCall)
}

func (fake *FakeBindingDB) GetServiceInstanceByAppIdCalls(stub func(string) (*models.ServiceInstance, error)) {
	fake.getServiceInstanceByAppIdMutex.Lock()
	defer fake.getServiceInstanceByAppIdMutex.Unlock()
	fake.GetServiceInstanceByAppIdStub = stub
}

func (fake *FakeBindingDB) GetServiceInstanceByAppIdArgsForCall(i int) string {
	fake.getServiceInstanceByAppIdMutex.RLock()
	defer fake.getServiceInstanceByAppIdMutex.RUnlock()
	argsForCall := fake.getServiceInstanceByAppIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) GetServiceInstanceByAppIdReturns(result1 *models.ServiceInstance, result2 error) {
	fake.getServiceInstanceByAppIdMutex.Lock()
	defer fake.getServiceInstanceByAppIdMutex.Unlock()
	fake.GetServiceInstanceByAppIdStub = nil
	fake.getServiceInstanceByAppIdReturns = struct {
		result1 *models.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) GetServiceInstanceByAppIdReturnsOnCall(i int, result1 *models.ServiceInstance, result2 error) {
	fake.getServiceInstanceByAppIdMutex.Lock()
	defer fake.getServiceInstanceByAppIdMutex.Unlock()
	fake.GetServiceInstanceByAppIdStub = nil
	if fake.getServiceInstanceByAppIdReturnsOnCall == nil {
		fake.getServiceInstanceByAppIdReturnsOnCall = make(map[int]struct {
			result1 *models.ServiceInstance
			result2 error
		})
	}
	fake.getServiceInstanceByAppIdReturnsOnCall[i] = struct {
		result1 *models.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeBindingDB) UpdateServiceInstance(arg1 models.ServiceInstance) error {
	fake.updateServiceInstanceMutex.Lock()
	ret, specificReturn := fake.updateServiceInstanceReturnsOnCall[len(fake.updateServiceInstanceArgsForCall)]
	fake.updateServiceInstanceArgsForCall = append(fake.updateServiceInstanceArgsForCall, struct {
		arg1 models.ServiceInstance
	}{arg1})
	stub := fake.UpdateServiceInstanceStub
	fakeReturns := fake.updateServiceInstanceReturns
	fake.recordInvocation("UpdateServiceInstance", []interface{}{arg1})
	fake.updateServiceInstanceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeBindingDB) UpdateServiceInstanceCallCount() int {
	fake.updateServiceInstanceMutex.RLock()
	defer fake.updateServiceInstanceMutex.RUnlock()
	return len(fake.updateServiceInstanceArgsForCall)
}

func (fake *FakeBindingDB) UpdateServiceInstanceCalls(stub func(models.ServiceInstance) error) {
	fake.updateServiceInstanceMutex.Lock()
	defer fake.updateServiceInstanceMutex.Unlock()
	fake.UpdateServiceInstanceStub = stub
}

func (fake *FakeBindingDB) UpdateServiceInstanceArgsForCall(i int) models.ServiceInstance {
	fake.updateServiceInstanceMutex.RLock()
	defer fake.updateServiceInstanceMutex.RUnlock()
	argsForCall := fake.updateServiceInstanceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBindingDB) UpdateServiceInstanceReturns(result1 error) {
	fake.updateServiceInstanceMutex.Lock()
	defer fake.updateServiceInstanceMutex.Unlock()
	fake.UpdateServiceInstanceStub = nil
	fake.updateServiceInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) UpdateServiceInstanceReturnsOnCall(i int, result1 error) {
	fake.updateServiceInstanceMutex.Lock()
	defer fake.updateServiceInstanceMutex.Unlock()
	fake.UpdateServiceInstanceStub = nil
	if fake.updateServiceInstanceReturnsOnCall == nil {
		fake.updateServiceInstanceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateServiceInstanceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBindingDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkServiceBindingMutex.RLock()
	defer fake.checkServiceBindingMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.countServiceInstancesInOrgMutex.RLock()
	defer fake.countServiceInstancesInOrgMutex.RUnlock()
	fake.createServiceBindingMutex.RLock()
	defer fake.createServiceBindingMutex.RUnlock()
	fake.createServiceInstanceMutex.RLock()
	defer fake.createServiceInstanceMutex.RUnlock()
	fake.deleteServiceBindingMutex.RLock()
	defer fake.deleteServiceBindingMutex.RUnlock()
	fake.deleteServiceBindingByAppIdMutex.RLock()
	defer fake.deleteServiceBindingByAppIdMutex.RUnlock()
	fake.deleteServiceInstanceMutex.RLock()
	defer fake.deleteServiceInstanceMutex.RUnlock()
	fake.getAppIdByBindingIdMutex.RLock()
	defer fake.getAppIdByBindingIdMutex.RUnlock()
	fake.getAppIdsByInstanceIdMutex.RLock()
	defer fake.getAppIdsByInstanceIdMutex.RUnlock()
	fake.getBindingIdsByInstanceIdMutex.RLock()
	defer fake.getBindingIdsByInstanceIdMutex.RUnlock()
	fake.getDBStatusMutex.RLock()
	defer fake.getDBStatusMutex.RUnlock()
	fake.getServiceInstanceMutex.RLock()
	defer fake.getServiceInstanceMutex.RUnlock()
	fake.getServiceInstanceByAppIdMutex.RLock()
	defer fake.getServiceInstanceByAppIdMutex.RUnlock()
	fake.updateServiceInstanceMutex.RLock()
	defer fake.updateServiceInstanceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBindingDB) recordInvocation(key string, args []interface{}) {
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

var _ db.BindingDB = new(FakeBindingDB)
