// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeTeamFactory struct {
	CreateTeamStub        func(name string) (*dbng.Team, error)
	createTeamMutex       sync.RWMutex
	createTeamArgsForCall []struct {
		name string
	}
	createTeamReturns struct {
		result1 *dbng.Team
		result2 error
	}
	FindTeamStub        func(name string) (*dbng.Team, bool, error)
	findTeamMutex       sync.RWMutex
	findTeamArgsForCall []struct {
		name string
	}
	findTeamReturns struct {
		result1 *dbng.Team
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTeamFactory) CreateTeam(name string) (*dbng.Team, error) {
	fake.createTeamMutex.Lock()
	fake.createTeamArgsForCall = append(fake.createTeamArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("CreateTeam", []interface{}{name})
	fake.createTeamMutex.Unlock()
	if fake.CreateTeamStub != nil {
		return fake.CreateTeamStub(name)
	} else {
		return fake.createTeamReturns.result1, fake.createTeamReturns.result2
	}
}

func (fake *FakeTeamFactory) CreateTeamCallCount() int {
	fake.createTeamMutex.RLock()
	defer fake.createTeamMutex.RUnlock()
	return len(fake.createTeamArgsForCall)
}

func (fake *FakeTeamFactory) CreateTeamArgsForCall(i int) string {
	fake.createTeamMutex.RLock()
	defer fake.createTeamMutex.RUnlock()
	return fake.createTeamArgsForCall[i].name
}

func (fake *FakeTeamFactory) CreateTeamReturns(result1 *dbng.Team, result2 error) {
	fake.CreateTeamStub = nil
	fake.createTeamReturns = struct {
		result1 *dbng.Team
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamFactory) FindTeam(name string) (*dbng.Team, bool, error) {
	fake.findTeamMutex.Lock()
	fake.findTeamArgsForCall = append(fake.findTeamArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("FindTeam", []interface{}{name})
	fake.findTeamMutex.Unlock()
	if fake.FindTeamStub != nil {
		return fake.FindTeamStub(name)
	} else {
		return fake.findTeamReturns.result1, fake.findTeamReturns.result2, fake.findTeamReturns.result3
	}
}

func (fake *FakeTeamFactory) FindTeamCallCount() int {
	fake.findTeamMutex.RLock()
	defer fake.findTeamMutex.RUnlock()
	return len(fake.findTeamArgsForCall)
}

func (fake *FakeTeamFactory) FindTeamArgsForCall(i int) string {
	fake.findTeamMutex.RLock()
	defer fake.findTeamMutex.RUnlock()
	return fake.findTeamArgsForCall[i].name
}

func (fake *FakeTeamFactory) FindTeamReturns(result1 *dbng.Team, result2 bool, result3 error) {
	fake.FindTeamStub = nil
	fake.findTeamReturns = struct {
		result1 *dbng.Team
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createTeamMutex.RLock()
	defer fake.createTeamMutex.RUnlock()
	fake.findTeamMutex.RLock()
	defer fake.findTeamMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTeamFactory) recordInvocation(key string, args []interface{}) {
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

var _ dbng.TeamFactory = new(FakeTeamFactory)