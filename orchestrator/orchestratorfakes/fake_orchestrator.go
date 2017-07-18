// Code generated by counterfeiter. DO NOT EDIT.
package orchestratorfakes

import (
	"sync"

	"github.com/cloudfoundry/uptimer/orchestrator"
)

type FakeOrchestrator struct {
	SetupStub        func() error
	setupMutex       sync.RWMutex
	setupArgsForCall []struct{}
	setupReturns     struct {
		result1 error
	}
	setupReturnsOnCall map[int]struct {
		result1 error
	}
	RunStub        func() (int, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct{}
	runReturns     struct {
		result1 int
		result2 error
	}
	runReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	TearDownStub        func() error
	tearDownMutex       sync.RWMutex
	tearDownArgsForCall []struct{}
	tearDownReturns     struct {
		result1 error
	}
	tearDownReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrchestrator) Setup() error {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct{}{})
	fake.recordInvocation("Setup", []interface{}{})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setupReturns.result1
}

func (fake *FakeOrchestrator) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeOrchestrator) SetupReturns(result1 error) {
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrchestrator) SetupReturnsOnCall(i int, result1 error) {
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrchestrator) Run() (int, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct{}{})
	fake.recordInvocation("Run", []interface{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.runReturns.result1, fake.runReturns.result2
}

func (fake *FakeOrchestrator) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeOrchestrator) RunReturns(result1 int, result2 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeOrchestrator) RunReturnsOnCall(i int, result1 int, result2 error) {
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeOrchestrator) TearDown() error {
	fake.tearDownMutex.Lock()
	ret, specificReturn := fake.tearDownReturnsOnCall[len(fake.tearDownArgsForCall)]
	fake.tearDownArgsForCall = append(fake.tearDownArgsForCall, struct{}{})
	fake.recordInvocation("TearDown", []interface{}{})
	fake.tearDownMutex.Unlock()
	if fake.TearDownStub != nil {
		return fake.TearDownStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.tearDownReturns.result1
}

func (fake *FakeOrchestrator) TearDownCallCount() int {
	fake.tearDownMutex.RLock()
	defer fake.tearDownMutex.RUnlock()
	return len(fake.tearDownArgsForCall)
}

func (fake *FakeOrchestrator) TearDownReturns(result1 error) {
	fake.TearDownStub = nil
	fake.tearDownReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrchestrator) TearDownReturnsOnCall(i int, result1 error) {
	fake.TearDownStub = nil
	if fake.tearDownReturnsOnCall == nil {
		fake.tearDownReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.tearDownReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrchestrator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.tearDownMutex.RLock()
	defer fake.tearDownMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOrchestrator) recordInvocation(key string, args []interface{}) {
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

var _ orchestrator.Orchestrator = new(FakeOrchestrator)
