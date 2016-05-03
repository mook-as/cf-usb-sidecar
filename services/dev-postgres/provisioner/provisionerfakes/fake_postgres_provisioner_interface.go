// This file was generated by counterfeiter
package provisionerfakes

import (
	"sync"

	"github.com/hpcloud/catalog-service-manager/services/dev-postgres/provisioner"
)

type FakePostgresProvisionerInterface struct {
	CreateDatabaseStub        func(string) error
	createDatabaseMutex       sync.RWMutex
	createDatabaseArgsForCall []struct {
		arg1 string
	}
	createDatabaseReturns struct {
		result1 error
	}
	DeleteDatabaseStub        func(string) error
	deleteDatabaseMutex       sync.RWMutex
	deleteDatabaseArgsForCall []struct {
		arg1 string
	}
	deleteDatabaseReturns struct {
		result1 error
	}
	DatabaseExistsStub        func(string) (bool, error)
	databaseExistsMutex       sync.RWMutex
	databaseExistsArgsForCall []struct {
		arg1 string
	}
	databaseExistsReturns struct {
		result1 bool
		result2 error
	}
	CreateUserStub        func(string, string, string) error
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createUserReturns struct {
		result1 error
	}
	DeleteUserStub        func(string, string) error
	deleteUserMutex       sync.RWMutex
	deleteUserArgsForCall []struct {
		arg1 string
		arg2 string
	}
	deleteUserReturns struct {
		result1 error
	}
	UserExistsStub        func(string) (bool, error)
	userExistsMutex       sync.RWMutex
	userExistsArgsForCall []struct {
		arg1 string
	}
	userExistsReturns struct {
		result1 bool
		result2 error
	}
}

func (fake *FakePostgresProvisionerInterface) CreateDatabase(arg1 string) error {
	fake.createDatabaseMutex.Lock()
	fake.createDatabaseArgsForCall = append(fake.createDatabaseArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.createDatabaseMutex.Unlock()
	if fake.CreateDatabaseStub != nil {
		return fake.CreateDatabaseStub(arg1)
	} else {
		return fake.createDatabaseReturns.result1
	}
}

func (fake *FakePostgresProvisionerInterface) CreateDatabaseCallCount() int {
	fake.createDatabaseMutex.RLock()
	defer fake.createDatabaseMutex.RUnlock()
	return len(fake.createDatabaseArgsForCall)
}

func (fake *FakePostgresProvisionerInterface) CreateDatabaseArgsForCall(i int) string {
	fake.createDatabaseMutex.RLock()
	defer fake.createDatabaseMutex.RUnlock()
	return fake.createDatabaseArgsForCall[i].arg1
}

func (fake *FakePostgresProvisionerInterface) CreateDatabaseReturns(result1 error) {
	fake.CreateDatabaseStub = nil
	fake.createDatabaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePostgresProvisionerInterface) DeleteDatabase(arg1 string) error {
	fake.deleteDatabaseMutex.Lock()
	fake.deleteDatabaseArgsForCall = append(fake.deleteDatabaseArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.deleteDatabaseMutex.Unlock()
	if fake.DeleteDatabaseStub != nil {
		return fake.DeleteDatabaseStub(arg1)
	} else {
		return fake.deleteDatabaseReturns.result1
	}
}

func (fake *FakePostgresProvisionerInterface) DeleteDatabaseCallCount() int {
	fake.deleteDatabaseMutex.RLock()
	defer fake.deleteDatabaseMutex.RUnlock()
	return len(fake.deleteDatabaseArgsForCall)
}

func (fake *FakePostgresProvisionerInterface) DeleteDatabaseArgsForCall(i int) string {
	fake.deleteDatabaseMutex.RLock()
	defer fake.deleteDatabaseMutex.RUnlock()
	return fake.deleteDatabaseArgsForCall[i].arg1
}

func (fake *FakePostgresProvisionerInterface) DeleteDatabaseReturns(result1 error) {
	fake.DeleteDatabaseStub = nil
	fake.deleteDatabaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePostgresProvisionerInterface) DatabaseExists(arg1 string) (bool, error) {
	fake.databaseExistsMutex.Lock()
	fake.databaseExistsArgsForCall = append(fake.databaseExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.databaseExistsMutex.Unlock()
	if fake.DatabaseExistsStub != nil {
		return fake.DatabaseExistsStub(arg1)
	} else {
		return fake.databaseExistsReturns.result1, fake.databaseExistsReturns.result2
	}
}

func (fake *FakePostgresProvisionerInterface) DatabaseExistsCallCount() int {
	fake.databaseExistsMutex.RLock()
	defer fake.databaseExistsMutex.RUnlock()
	return len(fake.databaseExistsArgsForCall)
}

func (fake *FakePostgresProvisionerInterface) DatabaseExistsArgsForCall(i int) string {
	fake.databaseExistsMutex.RLock()
	defer fake.databaseExistsMutex.RUnlock()
	return fake.databaseExistsArgsForCall[i].arg1
}

func (fake *FakePostgresProvisionerInterface) DatabaseExistsReturns(result1 bool, result2 error) {
	fake.DatabaseExistsStub = nil
	fake.databaseExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakePostgresProvisionerInterface) CreateUser(arg1 string, arg2 string, arg3 string) error {
	fake.createUserMutex.Lock()
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(arg1, arg2, arg3)
	} else {
		return fake.createUserReturns.result1
	}
}

func (fake *FakePostgresProvisionerInterface) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakePostgresProvisionerInterface) CreateUserArgsForCall(i int) (string, string, string) {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return fake.createUserArgsForCall[i].arg1, fake.createUserArgsForCall[i].arg2, fake.createUserArgsForCall[i].arg3
}

func (fake *FakePostgresProvisionerInterface) CreateUserReturns(result1 error) {
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePostgresProvisionerInterface) DeleteUser(arg1 string, arg2 string) error {
	fake.deleteUserMutex.Lock()
	fake.deleteUserArgsForCall = append(fake.deleteUserArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.deleteUserMutex.Unlock()
	if fake.DeleteUserStub != nil {
		return fake.DeleteUserStub(arg1, arg2)
	} else {
		return fake.deleteUserReturns.result1
	}
}

func (fake *FakePostgresProvisionerInterface) DeleteUserCallCount() int {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return len(fake.deleteUserArgsForCall)
}

func (fake *FakePostgresProvisionerInterface) DeleteUserArgsForCall(i int) (string, string) {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return fake.deleteUserArgsForCall[i].arg1, fake.deleteUserArgsForCall[i].arg2
}

func (fake *FakePostgresProvisionerInterface) DeleteUserReturns(result1 error) {
	fake.DeleteUserStub = nil
	fake.deleteUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePostgresProvisionerInterface) UserExists(arg1 string) (bool, error) {
	fake.userExistsMutex.Lock()
	fake.userExistsArgsForCall = append(fake.userExistsArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.userExistsMutex.Unlock()
	if fake.UserExistsStub != nil {
		return fake.UserExistsStub(arg1)
	} else {
		return fake.userExistsReturns.result1, fake.userExistsReturns.result2
	}
}

func (fake *FakePostgresProvisionerInterface) UserExistsCallCount() int {
	fake.userExistsMutex.RLock()
	defer fake.userExistsMutex.RUnlock()
	return len(fake.userExistsArgsForCall)
}

func (fake *FakePostgresProvisionerInterface) UserExistsArgsForCall(i int) string {
	fake.userExistsMutex.RLock()
	defer fake.userExistsMutex.RUnlock()
	return fake.userExistsArgsForCall[i].arg1
}

func (fake *FakePostgresProvisionerInterface) UserExistsReturns(result1 bool, result2 error) {
	fake.UserExistsStub = nil
	fake.userExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

var _ provisioner.PostgresProvisionerInterface = new(FakePostgresProvisionerInterface)
