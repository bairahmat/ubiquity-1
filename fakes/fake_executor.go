/**
 * Copyright 2017 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"os"
	"sync"

	"github.com/IBM/ubiquity/utils"
)

type FakeExecutor struct {
	ExecuteStub        func(command string, args []string) ([]byte, error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		command string
		args    []string
	}
	executeReturns struct {
		result1 []byte
		result2 error
	}
	executeReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	StatStub        func(string) (os.FileInfo, error)
	statMutex       sync.RWMutex
	statArgsForCall []struct {
		arg1 string
	}
	statReturns struct {
		result1 os.FileInfo
		result2 error
	}
	statReturnsOnCall map[int]struct {
		result1 os.FileInfo
		result2 error
	}
	MkdirStub        func(string, os.FileMode) error
	mkdirMutex       sync.RWMutex
	mkdirArgsForCall []struct {
		arg1 string
		arg2 os.FileMode
	}
	mkdirReturns struct {
		result1 error
	}
	mkdirReturnsOnCall map[int]struct {
		result1 error
	}
	MkdirAllStub        func(string, os.FileMode) error
	mkdirAllMutex       sync.RWMutex
	mkdirAllArgsForCall []struct {
		arg1 string
		arg2 os.FileMode
	}
	mkdirAllReturns struct {
		result1 error
	}
	mkdirAllReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveAllStub        func(string) error
	removeAllMutex       sync.RWMutex
	removeAllArgsForCall []struct {
		arg1 string
	}
	removeAllReturns struct {
		result1 error
	}
	removeAllReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveStub        func(string) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 string
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	HostnameStub        func() (string, error)
	hostnameMutex       sync.RWMutex
	hostnameArgsForCall []struct{}
	hostnameReturns     struct {
		result1 string
		result2 error
	}
	hostnameReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	IsExecutableStub        func(string) error
	isExecutableMutex       sync.RWMutex
	isExecutableArgsForCall []struct {
		arg1 string
	}
	isExecutableReturns struct {
		result1 error
	}
	isExecutableReturnsOnCall map[int]struct {
		result1 error
	}
	EvalSymlinksStub        func(path string) (string, error)
	evalSymlinksMutex       sync.RWMutex
	evalSymlinksArgsForCall []struct {
		path string
	}
	evalSymlinksReturns struct {
		result1 string
		result2 error
	}
	evalSymlinksReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeExecutor) Execute(command string, args []string) ([]byte, error) {
	var argsCopy []string
	if args != nil {
		argsCopy = make([]string, len(args))
		copy(argsCopy, args)
	}
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		command string
		args    []string
	}{command, argsCopy})
	fake.recordInvocation("Execute", []interface{}{command, argsCopy})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(command, args)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.executeReturns.result1, fake.executeReturns.result2
}

func (fake *FakeExecutor) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeExecutor) ExecuteArgsForCall(i int) (string, []string) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].command, fake.executeArgsForCall[i].args
}

func (fake *FakeExecutor) ExecuteReturns(result1 []byte, result2 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeExecutor) ExecuteReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeExecutor) Stat(arg1 string) (os.FileInfo, error) {
	fake.statMutex.Lock()
	ret, specificReturn := fake.statReturnsOnCall[len(fake.statArgsForCall)]
	fake.statArgsForCall = append(fake.statArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Stat", []interface{}{arg1})
	fake.statMutex.Unlock()
	if fake.StatStub != nil {
		return fake.StatStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.statReturns.result1, fake.statReturns.result2
}

func (fake *FakeExecutor) StatCallCount() int {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return len(fake.statArgsForCall)
}

func (fake *FakeExecutor) StatArgsForCall(i int) string {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return fake.statArgsForCall[i].arg1
}

func (fake *FakeExecutor) StatReturns(result1 os.FileInfo, result2 error) {
	fake.StatStub = nil
	fake.statReturns = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeExecutor) StatReturnsOnCall(i int, result1 os.FileInfo, result2 error) {
	fake.StatStub = nil
	if fake.statReturnsOnCall == nil {
		fake.statReturnsOnCall = make(map[int]struct {
			result1 os.FileInfo
			result2 error
		})
	}
	fake.statReturnsOnCall[i] = struct {
		result1 os.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeExecutor) Mkdir(arg1 string, arg2 os.FileMode) error {
	fake.mkdirMutex.Lock()
	ret, specificReturn := fake.mkdirReturnsOnCall[len(fake.mkdirArgsForCall)]
	fake.mkdirArgsForCall = append(fake.mkdirArgsForCall, struct {
		arg1 string
		arg2 os.FileMode
	}{arg1, arg2})
	fake.recordInvocation("Mkdir", []interface{}{arg1, arg2})
	fake.mkdirMutex.Unlock()
	if fake.MkdirStub != nil {
		return fake.MkdirStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.mkdirReturns.result1
}

func (fake *FakeExecutor) MkdirCallCount() int {
	fake.mkdirMutex.RLock()
	defer fake.mkdirMutex.RUnlock()
	return len(fake.mkdirArgsForCall)
}

func (fake *FakeExecutor) MkdirArgsForCall(i int) (string, os.FileMode) {
	fake.mkdirMutex.RLock()
	defer fake.mkdirMutex.RUnlock()
	return fake.mkdirArgsForCall[i].arg1, fake.mkdirArgsForCall[i].arg2
}

func (fake *FakeExecutor) MkdirReturns(result1 error) {
	fake.MkdirStub = nil
	fake.mkdirReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) MkdirReturnsOnCall(i int, result1 error) {
	fake.MkdirStub = nil
	if fake.mkdirReturnsOnCall == nil {
		fake.mkdirReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mkdirReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) MkdirAll(arg1 string, arg2 os.FileMode) error {
	fake.mkdirAllMutex.Lock()
	ret, specificReturn := fake.mkdirAllReturnsOnCall[len(fake.mkdirAllArgsForCall)]
	fake.mkdirAllArgsForCall = append(fake.mkdirAllArgsForCall, struct {
		arg1 string
		arg2 os.FileMode
	}{arg1, arg2})
	fake.recordInvocation("MkdirAll", []interface{}{arg1, arg2})
	fake.mkdirAllMutex.Unlock()
	if fake.MkdirAllStub != nil {
		return fake.MkdirAllStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.mkdirAllReturns.result1
}

func (fake *FakeExecutor) MkdirAllCallCount() int {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	return len(fake.mkdirAllArgsForCall)
}

func (fake *FakeExecutor) MkdirAllArgsForCall(i int) (string, os.FileMode) {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	return fake.mkdirAllArgsForCall[i].arg1, fake.mkdirAllArgsForCall[i].arg2
}

func (fake *FakeExecutor) MkdirAllReturns(result1 error) {
	fake.MkdirAllStub = nil
	fake.mkdirAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) MkdirAllReturnsOnCall(i int, result1 error) {
	fake.MkdirAllStub = nil
	if fake.mkdirAllReturnsOnCall == nil {
		fake.mkdirAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mkdirAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) RemoveAll(arg1 string) error {
	fake.removeAllMutex.Lock()
	ret, specificReturn := fake.removeAllReturnsOnCall[len(fake.removeAllArgsForCall)]
	fake.removeAllArgsForCall = append(fake.removeAllArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RemoveAll", []interface{}{arg1})
	fake.removeAllMutex.Unlock()
	if fake.RemoveAllStub != nil {
		return fake.RemoveAllStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeAllReturns.result1
}

func (fake *FakeExecutor) RemoveAllCallCount() int {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	return len(fake.removeAllArgsForCall)
}

func (fake *FakeExecutor) RemoveAllArgsForCall(i int) string {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	return fake.removeAllArgsForCall[i].arg1
}

func (fake *FakeExecutor) RemoveAllReturns(result1 error) {
	fake.RemoveAllStub = nil
	fake.removeAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) RemoveAllReturnsOnCall(i int, result1 error) {
	fake.RemoveAllStub = nil
	if fake.removeAllReturnsOnCall == nil {
		fake.removeAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) Remove(arg1 string) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Remove", []interface{}{arg1})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.removeReturns.result1
}

func (fake *FakeExecutor) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeExecutor) RemoveArgsForCall(i int) string {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].arg1
}

func (fake *FakeExecutor) RemoveReturns(result1 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) RemoveReturnsOnCall(i int, result1 error) {
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) Hostname() (string, error) {
	fake.hostnameMutex.Lock()
	ret, specificReturn := fake.hostnameReturnsOnCall[len(fake.hostnameArgsForCall)]
	fake.hostnameArgsForCall = append(fake.hostnameArgsForCall, struct{}{})
	fake.recordInvocation("Hostname", []interface{}{})
	fake.hostnameMutex.Unlock()
	if fake.HostnameStub != nil {
		return fake.HostnameStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.hostnameReturns.result1, fake.hostnameReturns.result2
}

func (fake *FakeExecutor) HostnameCallCount() int {
	fake.hostnameMutex.RLock()
	defer fake.hostnameMutex.RUnlock()
	return len(fake.hostnameArgsForCall)
}

func (fake *FakeExecutor) HostnameReturns(result1 string, result2 error) {
	fake.HostnameStub = nil
	fake.hostnameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeExecutor) HostnameReturnsOnCall(i int, result1 string, result2 error) {
	fake.HostnameStub = nil
	if fake.hostnameReturnsOnCall == nil {
		fake.hostnameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.hostnameReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeExecutor) IsExecutable(arg1 string) error {
	fake.isExecutableMutex.Lock()
	ret, specificReturn := fake.isExecutableReturnsOnCall[len(fake.isExecutableArgsForCall)]
	fake.isExecutableArgsForCall = append(fake.isExecutableArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("IsExecutable", []interface{}{arg1})
	fake.isExecutableMutex.Unlock()
	if fake.IsExecutableStub != nil {
		return fake.IsExecutableStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isExecutableReturns.result1
}

func (fake *FakeExecutor) IsExecutableCallCount() int {
	fake.isExecutableMutex.RLock()
	defer fake.isExecutableMutex.RUnlock()
	return len(fake.isExecutableArgsForCall)
}

func (fake *FakeExecutor) IsExecutableArgsForCall(i int) string {
	fake.isExecutableMutex.RLock()
	defer fake.isExecutableMutex.RUnlock()
	return fake.isExecutableArgsForCall[i].arg1
}

func (fake *FakeExecutor) IsExecutableReturns(result1 error) {
	fake.IsExecutableStub = nil
	fake.isExecutableReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) IsExecutableReturnsOnCall(i int, result1 error) {
	fake.IsExecutableStub = nil
	if fake.isExecutableReturnsOnCall == nil {
		fake.isExecutableReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.isExecutableReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeExecutor) EvalSymlinks(path string) (string, error) {
	fake.evalSymlinksMutex.Lock()
	ret, specificReturn := fake.evalSymlinksReturnsOnCall[len(fake.evalSymlinksArgsForCall)]
	fake.evalSymlinksArgsForCall = append(fake.evalSymlinksArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("EvalSymlinks", []interface{}{path})
	fake.evalSymlinksMutex.Unlock()
	if fake.EvalSymlinksStub != nil {
		return fake.EvalSymlinksStub(path)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.evalSymlinksReturns.result1, fake.evalSymlinksReturns.result2
}

func (fake *FakeExecutor) EvalSymlinksCallCount() int {
	fake.evalSymlinksMutex.RLock()
	defer fake.evalSymlinksMutex.RUnlock()
	return len(fake.evalSymlinksArgsForCall)
}

func (fake *FakeExecutor) EvalSymlinksArgsForCall(i int) string {
	fake.evalSymlinksMutex.RLock()
	defer fake.evalSymlinksMutex.RUnlock()
	return fake.evalSymlinksArgsForCall[i].path
}

func (fake *FakeExecutor) EvalSymlinksReturns(result1 string, result2 error) {
	fake.EvalSymlinksStub = nil
	fake.evalSymlinksReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeExecutor) EvalSymlinksReturnsOnCall(i int, result1 string, result2 error) {
	fake.EvalSymlinksStub = nil
	if fake.evalSymlinksReturnsOnCall == nil {
		fake.evalSymlinksReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.evalSymlinksReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeExecutor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	fake.mkdirMutex.RLock()
	defer fake.mkdirMutex.RUnlock()
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.hostnameMutex.RLock()
	defer fake.hostnameMutex.RUnlock()
	fake.isExecutableMutex.RLock()
	defer fake.isExecutableMutex.RUnlock()
	fake.evalSymlinksMutex.RLock()
	defer fake.evalSymlinksMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeExecutor) recordInvocation(key string, args []interface{}) {
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

var _ utils.Executor = new(FakeExecutor)
