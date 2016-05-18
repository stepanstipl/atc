// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakeTeamDB struct {
	GetAllPipelinesStub        func() ([]db.SavedPipeline, error)
	getAllPipelinesMutex       sync.RWMutex
	getAllPipelinesArgsForCall []struct{}
	getAllPipelinesReturns     struct {
		result1 []db.SavedPipeline
		result2 error
	}
	GetPipelineByIDStub        func(pipelineID int) (db.SavedPipeline, error)
	getPipelineByIDMutex       sync.RWMutex
	getPipelineByIDArgsForCall []struct {
		pipelineID int
	}
	getPipelineByIDReturns struct {
		result1 db.SavedPipeline
		result2 error
	}
	GetPipelineByTeamNameAndNameStub        func(teamName string, pipelineName string) (db.SavedPipeline, error)
	getPipelineByTeamNameAndNameMutex       sync.RWMutex
	getPipelineByTeamNameAndNameArgsForCall []struct {
		teamName     string
		pipelineName string
	}
	getPipelineByTeamNameAndNameReturns struct {
		result1 db.SavedPipeline
		result2 error
	}
	OrderPipelinesStub        func([]string) error
	orderPipelinesMutex       sync.RWMutex
	orderPipelinesArgsForCall []struct {
		arg1 []string
	}
	orderPipelinesReturns struct {
		result1 error
	}
	GetTeamByNameStub        func(teamName string) (db.SavedTeam, bool, error)
	getTeamByNameMutex       sync.RWMutex
	getTeamByNameArgsForCall []struct {
		teamName string
	}
	getTeamByNameReturns struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}
	SaveTeamStub        func(team db.Team) (db.SavedTeam, error)
	saveTeamMutex       sync.RWMutex
	saveTeamArgsForCall []struct {
		team db.Team
	}
	saveTeamReturns struct {
		result1 db.SavedTeam
		result2 error
	}
	UpdateTeamBasicAuthStub        func(team db.Team) (db.SavedTeam, error)
	updateTeamBasicAuthMutex       sync.RWMutex
	updateTeamBasicAuthArgsForCall []struct {
		team db.Team
	}
	updateTeamBasicAuthReturns struct {
		result1 db.SavedTeam
		result2 error
	}
	UpdateTeamGitHubAuthStub        func(team db.Team) (db.SavedTeam, error)
	updateTeamGitHubAuthMutex       sync.RWMutex
	updateTeamGitHubAuthArgsForCall []struct {
		team db.Team
	}
	updateTeamGitHubAuthReturns struct {
		result1 db.SavedTeam
		result2 error
	}
}

func (fake *FakeTeamDB) GetAllPipelines() ([]db.SavedPipeline, error) {
	fake.getAllPipelinesMutex.Lock()
	fake.getAllPipelinesArgsForCall = append(fake.getAllPipelinesArgsForCall, struct{}{})
	fake.getAllPipelinesMutex.Unlock()
	if fake.GetAllPipelinesStub != nil {
		return fake.GetAllPipelinesStub()
	} else {
		return fake.getAllPipelinesReturns.result1, fake.getAllPipelinesReturns.result2
	}
}

func (fake *FakeTeamDB) GetAllPipelinesCallCount() int {
	fake.getAllPipelinesMutex.RLock()
	defer fake.getAllPipelinesMutex.RUnlock()
	return len(fake.getAllPipelinesArgsForCall)
}

func (fake *FakeTeamDB) GetAllPipelinesReturns(result1 []db.SavedPipeline, result2 error) {
	fake.GetAllPipelinesStub = nil
	fake.getAllPipelinesReturns = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetPipelineByID(pipelineID int) (db.SavedPipeline, error) {
	fake.getPipelineByIDMutex.Lock()
	fake.getPipelineByIDArgsForCall = append(fake.getPipelineByIDArgsForCall, struct {
		pipelineID int
	}{pipelineID})
	fake.getPipelineByIDMutex.Unlock()
	if fake.GetPipelineByIDStub != nil {
		return fake.GetPipelineByIDStub(pipelineID)
	} else {
		return fake.getPipelineByIDReturns.result1, fake.getPipelineByIDReturns.result2
	}
}

func (fake *FakeTeamDB) GetPipelineByIDCallCount() int {
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	return len(fake.getPipelineByIDArgsForCall)
}

func (fake *FakeTeamDB) GetPipelineByIDArgsForCall(i int) int {
	fake.getPipelineByIDMutex.RLock()
	defer fake.getPipelineByIDMutex.RUnlock()
	return fake.getPipelineByIDArgsForCall[i].pipelineID
}

func (fake *FakeTeamDB) GetPipelineByIDReturns(result1 db.SavedPipeline, result2 error) {
	fake.GetPipelineByIDStub = nil
	fake.getPipelineByIDReturns = struct {
		result1 db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetPipelineByTeamNameAndName(teamName string, pipelineName string) (db.SavedPipeline, error) {
	fake.getPipelineByTeamNameAndNameMutex.Lock()
	fake.getPipelineByTeamNameAndNameArgsForCall = append(fake.getPipelineByTeamNameAndNameArgsForCall, struct {
		teamName     string
		pipelineName string
	}{teamName, pipelineName})
	fake.getPipelineByTeamNameAndNameMutex.Unlock()
	if fake.GetPipelineByTeamNameAndNameStub != nil {
		return fake.GetPipelineByTeamNameAndNameStub(teamName, pipelineName)
	} else {
		return fake.getPipelineByTeamNameAndNameReturns.result1, fake.getPipelineByTeamNameAndNameReturns.result2
	}
}

func (fake *FakeTeamDB) GetPipelineByTeamNameAndNameCallCount() int {
	fake.getPipelineByTeamNameAndNameMutex.RLock()
	defer fake.getPipelineByTeamNameAndNameMutex.RUnlock()
	return len(fake.getPipelineByTeamNameAndNameArgsForCall)
}

func (fake *FakeTeamDB) GetPipelineByTeamNameAndNameArgsForCall(i int) (string, string) {
	fake.getPipelineByTeamNameAndNameMutex.RLock()
	defer fake.getPipelineByTeamNameAndNameMutex.RUnlock()
	return fake.getPipelineByTeamNameAndNameArgsForCall[i].teamName, fake.getPipelineByTeamNameAndNameArgsForCall[i].pipelineName
}

func (fake *FakeTeamDB) GetPipelineByTeamNameAndNameReturns(result1 db.SavedPipeline, result2 error) {
	fake.GetPipelineByTeamNameAndNameStub = nil
	fake.getPipelineByTeamNameAndNameReturns = struct {
		result1 db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) OrderPipelines(arg1 []string) error {
	fake.orderPipelinesMutex.Lock()
	fake.orderPipelinesArgsForCall = append(fake.orderPipelinesArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.orderPipelinesMutex.Unlock()
	if fake.OrderPipelinesStub != nil {
		return fake.OrderPipelinesStub(arg1)
	} else {
		return fake.orderPipelinesReturns.result1
	}
}

func (fake *FakeTeamDB) OrderPipelinesCallCount() int {
	fake.orderPipelinesMutex.RLock()
	defer fake.orderPipelinesMutex.RUnlock()
	return len(fake.orderPipelinesArgsForCall)
}

func (fake *FakeTeamDB) OrderPipelinesArgsForCall(i int) []string {
	fake.orderPipelinesMutex.RLock()
	defer fake.orderPipelinesMutex.RUnlock()
	return fake.orderPipelinesArgsForCall[i].arg1
}

func (fake *FakeTeamDB) OrderPipelinesReturns(result1 error) {
	fake.OrderPipelinesStub = nil
	fake.orderPipelinesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTeamDB) GetTeamByName(teamName string) (db.SavedTeam, bool, error) {
	fake.getTeamByNameMutex.Lock()
	fake.getTeamByNameArgsForCall = append(fake.getTeamByNameArgsForCall, struct {
		teamName string
	}{teamName})
	fake.getTeamByNameMutex.Unlock()
	if fake.GetTeamByNameStub != nil {
		return fake.GetTeamByNameStub(teamName)
	} else {
		return fake.getTeamByNameReturns.result1, fake.getTeamByNameReturns.result2, fake.getTeamByNameReturns.result3
	}
}

func (fake *FakeTeamDB) GetTeamByNameCallCount() int {
	fake.getTeamByNameMutex.RLock()
	defer fake.getTeamByNameMutex.RUnlock()
	return len(fake.getTeamByNameArgsForCall)
}

func (fake *FakeTeamDB) GetTeamByNameArgsForCall(i int) string {
	fake.getTeamByNameMutex.RLock()
	defer fake.getTeamByNameMutex.RUnlock()
	return fake.getTeamByNameArgsForCall[i].teamName
}

func (fake *FakeTeamDB) GetTeamByNameReturns(result1 db.SavedTeam, result2 bool, result3 error) {
	fake.GetTeamByNameStub = nil
	fake.getTeamByNameReturns = struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) SaveTeam(team db.Team) (db.SavedTeam, error) {
	fake.saveTeamMutex.Lock()
	fake.saveTeamArgsForCall = append(fake.saveTeamArgsForCall, struct {
		team db.Team
	}{team})
	fake.saveTeamMutex.Unlock()
	if fake.SaveTeamStub != nil {
		return fake.SaveTeamStub(team)
	} else {
		return fake.saveTeamReturns.result1, fake.saveTeamReturns.result2
	}
}

func (fake *FakeTeamDB) SaveTeamCallCount() int {
	fake.saveTeamMutex.RLock()
	defer fake.saveTeamMutex.RUnlock()
	return len(fake.saveTeamArgsForCall)
}

func (fake *FakeTeamDB) SaveTeamArgsForCall(i int) db.Team {
	fake.saveTeamMutex.RLock()
	defer fake.saveTeamMutex.RUnlock()
	return fake.saveTeamArgsForCall[i].team
}

func (fake *FakeTeamDB) SaveTeamReturns(result1 db.SavedTeam, result2 error) {
	fake.SaveTeamStub = nil
	fake.saveTeamReturns = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) UpdateTeamBasicAuth(team db.Team) (db.SavedTeam, error) {
	fake.updateTeamBasicAuthMutex.Lock()
	fake.updateTeamBasicAuthArgsForCall = append(fake.updateTeamBasicAuthArgsForCall, struct {
		team db.Team
	}{team})
	fake.updateTeamBasicAuthMutex.Unlock()
	if fake.UpdateTeamBasicAuthStub != nil {
		return fake.UpdateTeamBasicAuthStub(team)
	} else {
		return fake.updateTeamBasicAuthReturns.result1, fake.updateTeamBasicAuthReturns.result2
	}
}

func (fake *FakeTeamDB) UpdateTeamBasicAuthCallCount() int {
	fake.updateTeamBasicAuthMutex.RLock()
	defer fake.updateTeamBasicAuthMutex.RUnlock()
	return len(fake.updateTeamBasicAuthArgsForCall)
}

func (fake *FakeTeamDB) UpdateTeamBasicAuthArgsForCall(i int) db.Team {
	fake.updateTeamBasicAuthMutex.RLock()
	defer fake.updateTeamBasicAuthMutex.RUnlock()
	return fake.updateTeamBasicAuthArgsForCall[i].team
}

func (fake *FakeTeamDB) UpdateTeamBasicAuthReturns(result1 db.SavedTeam, result2 error) {
	fake.UpdateTeamBasicAuthStub = nil
	fake.updateTeamBasicAuthReturns = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) UpdateTeamGitHubAuth(team db.Team) (db.SavedTeam, error) {
	fake.updateTeamGitHubAuthMutex.Lock()
	fake.updateTeamGitHubAuthArgsForCall = append(fake.updateTeamGitHubAuthArgsForCall, struct {
		team db.Team
	}{team})
	fake.updateTeamGitHubAuthMutex.Unlock()
	if fake.UpdateTeamGitHubAuthStub != nil {
		return fake.UpdateTeamGitHubAuthStub(team)
	} else {
		return fake.updateTeamGitHubAuthReturns.result1, fake.updateTeamGitHubAuthReturns.result2
	}
}

func (fake *FakeTeamDB) UpdateTeamGitHubAuthCallCount() int {
	fake.updateTeamGitHubAuthMutex.RLock()
	defer fake.updateTeamGitHubAuthMutex.RUnlock()
	return len(fake.updateTeamGitHubAuthArgsForCall)
}

func (fake *FakeTeamDB) UpdateTeamGitHubAuthArgsForCall(i int) db.Team {
	fake.updateTeamGitHubAuthMutex.RLock()
	defer fake.updateTeamGitHubAuthMutex.RUnlock()
	return fake.updateTeamGitHubAuthArgsForCall[i].team
}

func (fake *FakeTeamDB) UpdateTeamGitHubAuthReturns(result1 db.SavedTeam, result2 error) {
	fake.UpdateTeamGitHubAuthStub = nil
	fake.updateTeamGitHubAuthReturns = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

var _ db.TeamDB = new(FakeTeamDB)