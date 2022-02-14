/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package signfakes

import (
	"context"
	"sync"

	"github.com/sigstore/cosign/cmd/cosign/cli/options"
	signa "github.com/sigstore/cosign/cmd/cosign/cli/sign"
	"sigs.k8s.io/release-sdk/sign"
)

type FakeImpl struct {
	EnvDefaultStub        func(string, string) string
	envDefaultMutex       sync.RWMutex
	envDefaultArgsForCall []struct {
		arg1 string
		arg2 string
	}
	envDefaultReturns struct {
		result1 string
	}
	envDefaultReturnsOnCall map[int]struct {
		result1 string
	}
	IsImageSignedInternalStub        func(context.Context, string) (bool, error)
	isImageSignedInternalMutex       sync.RWMutex
	isImageSignedInternalArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	isImageSignedInternalReturns struct {
		result1 bool
		result2 error
	}
	isImageSignedInternalReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	SetenvStub        func(string, string) error
	setenvMutex       sync.RWMutex
	setenvArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setenvReturns struct {
		result1 error
	}
	setenvReturnsOnCall map[int]struct {
		result1 error
	}
	SignImageInternalStub        func(context.Context, signa.KeyOpts, options.RegistryOptions, map[string]interface{}, []string, string, bool, string, string, string, bool, bool, string) error
	signImageInternalMutex       sync.RWMutex
	signImageInternalArgsForCall []struct {
		arg1  context.Context
		arg2  signa.KeyOpts
		arg3  options.RegistryOptions
		arg4  map[string]interface{}
		arg5  []string
		arg6  string
		arg7  bool
		arg8  string
		arg9  string
		arg10 string
		arg11 bool
		arg12 bool
		arg13 string
	}
	signImageInternalReturns struct {
		result1 error
	}
	signImageInternalReturnsOnCall map[int]struct {
		result1 error
	}
	VerifyFileInternalStub        func(*sign.Signer, string) (*sign.SignedObject, error)
	verifyFileInternalMutex       sync.RWMutex
	verifyFileInternalArgsForCall []struct {
		arg1 *sign.Signer
		arg2 string
	}
	verifyFileInternalReturns struct {
		result1 *sign.SignedObject
		result2 error
	}
	verifyFileInternalReturnsOnCall map[int]struct {
		result1 *sign.SignedObject
		result2 error
	}
	VerifyImageInternalStub        func(context.Context, string, []string) (*sign.SignedObject, error)
	verifyImageInternalMutex       sync.RWMutex
	verifyImageInternalArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []string
	}
	verifyImageInternalReturns struct {
		result1 *sign.SignedObject
		result2 error
	}
	verifyImageInternalReturnsOnCall map[int]struct {
		result1 *sign.SignedObject
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) EnvDefault(arg1 string, arg2 string) string {
	fake.envDefaultMutex.Lock()
	ret, specificReturn := fake.envDefaultReturnsOnCall[len(fake.envDefaultArgsForCall)]
	fake.envDefaultArgsForCall = append(fake.envDefaultArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.EnvDefaultStub
	fakeReturns := fake.envDefaultReturns
	fake.recordInvocation("EnvDefault", []interface{}{arg1, arg2})
	fake.envDefaultMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) EnvDefaultCallCount() int {
	fake.envDefaultMutex.RLock()
	defer fake.envDefaultMutex.RUnlock()
	return len(fake.envDefaultArgsForCall)
}

func (fake *FakeImpl) EnvDefaultCalls(stub func(string, string) string) {
	fake.envDefaultMutex.Lock()
	defer fake.envDefaultMutex.Unlock()
	fake.EnvDefaultStub = stub
}

func (fake *FakeImpl) EnvDefaultArgsForCall(i int) (string, string) {
	fake.envDefaultMutex.RLock()
	defer fake.envDefaultMutex.RUnlock()
	argsForCall := fake.envDefaultArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) EnvDefaultReturns(result1 string) {
	fake.envDefaultMutex.Lock()
	defer fake.envDefaultMutex.Unlock()
	fake.EnvDefaultStub = nil
	fake.envDefaultReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeImpl) EnvDefaultReturnsOnCall(i int, result1 string) {
	fake.envDefaultMutex.Lock()
	defer fake.envDefaultMutex.Unlock()
	fake.EnvDefaultStub = nil
	if fake.envDefaultReturnsOnCall == nil {
		fake.envDefaultReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.envDefaultReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeImpl) IsImageSignedInternal(arg1 context.Context, arg2 string) (bool, error) {
	fake.isImageSignedInternalMutex.Lock()
	ret, specificReturn := fake.isImageSignedInternalReturnsOnCall[len(fake.isImageSignedInternalArgsForCall)]
	fake.isImageSignedInternalArgsForCall = append(fake.isImageSignedInternalArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.IsImageSignedInternalStub
	fakeReturns := fake.isImageSignedInternalReturns
	fake.recordInvocation("IsImageSignedInternal", []interface{}{arg1, arg2})
	fake.isImageSignedInternalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) IsImageSignedInternalCallCount() int {
	fake.isImageSignedInternalMutex.RLock()
	defer fake.isImageSignedInternalMutex.RUnlock()
	return len(fake.isImageSignedInternalArgsForCall)
}

func (fake *FakeImpl) IsImageSignedInternalCalls(stub func(context.Context, string) (bool, error)) {
	fake.isImageSignedInternalMutex.Lock()
	defer fake.isImageSignedInternalMutex.Unlock()
	fake.IsImageSignedInternalStub = stub
}

func (fake *FakeImpl) IsImageSignedInternalArgsForCall(i int) (context.Context, string) {
	fake.isImageSignedInternalMutex.RLock()
	defer fake.isImageSignedInternalMutex.RUnlock()
	argsForCall := fake.isImageSignedInternalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) IsImageSignedInternalReturns(result1 bool, result2 error) {
	fake.isImageSignedInternalMutex.Lock()
	defer fake.isImageSignedInternalMutex.Unlock()
	fake.IsImageSignedInternalStub = nil
	fake.isImageSignedInternalReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) IsImageSignedInternalReturnsOnCall(i int, result1 bool, result2 error) {
	fake.isImageSignedInternalMutex.Lock()
	defer fake.isImageSignedInternalMutex.Unlock()
	fake.IsImageSignedInternalStub = nil
	if fake.isImageSignedInternalReturnsOnCall == nil {
		fake.isImageSignedInternalReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isImageSignedInternalReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Setenv(arg1 string, arg2 string) error {
	fake.setenvMutex.Lock()
	ret, specificReturn := fake.setenvReturnsOnCall[len(fake.setenvArgsForCall)]
	fake.setenvArgsForCall = append(fake.setenvArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.SetenvStub
	fakeReturns := fake.setenvReturns
	fake.recordInvocation("Setenv", []interface{}{arg1, arg2})
	fake.setenvMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) SetenvCallCount() int {
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	return len(fake.setenvArgsForCall)
}

func (fake *FakeImpl) SetenvCalls(stub func(string, string) error) {
	fake.setenvMutex.Lock()
	defer fake.setenvMutex.Unlock()
	fake.SetenvStub = stub
}

func (fake *FakeImpl) SetenvArgsForCall(i int) (string, string) {
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	argsForCall := fake.setenvArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) SetenvReturns(result1 error) {
	fake.setenvMutex.Lock()
	defer fake.setenvMutex.Unlock()
	fake.SetenvStub = nil
	fake.setenvReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SetenvReturnsOnCall(i int, result1 error) {
	fake.setenvMutex.Lock()
	defer fake.setenvMutex.Unlock()
	fake.SetenvStub = nil
	if fake.setenvReturnsOnCall == nil {
		fake.setenvReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setenvReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SignImageInternal(arg1 context.Context, arg2 signa.KeyOpts, arg3 options.RegistryOptions, arg4 map[string]interface{}, arg5 []string, arg6 string, arg7 bool, arg8 string, arg9 string, arg10 string, arg11 bool, arg12 bool, arg13 string) error {
	var arg5Copy []string
	if arg5 != nil {
		arg5Copy = make([]string, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.signImageInternalMutex.Lock()
	ret, specificReturn := fake.signImageInternalReturnsOnCall[len(fake.signImageInternalArgsForCall)]
	fake.signImageInternalArgsForCall = append(fake.signImageInternalArgsForCall, struct {
		arg1  context.Context
		arg2  signa.KeyOpts
		arg3  options.RegistryOptions
		arg4  map[string]interface{}
		arg5  []string
		arg6  string
		arg7  bool
		arg8  string
		arg9  string
		arg10 string
		arg11 bool
		arg12 bool
		arg13 string
	}{arg1, arg2, arg3, arg4, arg5Copy, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13})
	stub := fake.SignImageInternalStub
	fakeReturns := fake.signImageInternalReturns
	fake.recordInvocation("SignImageInternal", []interface{}{arg1, arg2, arg3, arg4, arg5Copy, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13})
	fake.signImageInternalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) SignImageInternalCallCount() int {
	fake.signImageInternalMutex.RLock()
	defer fake.signImageInternalMutex.RUnlock()
	return len(fake.signImageInternalArgsForCall)
}

func (fake *FakeImpl) SignImageInternalCalls(stub func(context.Context, signa.KeyOpts, options.RegistryOptions, map[string]interface{}, []string, string, bool, string, string, string, bool, bool, string) error) {
	fake.signImageInternalMutex.Lock()
	defer fake.signImageInternalMutex.Unlock()
	fake.SignImageInternalStub = stub
}

func (fake *FakeImpl) SignImageInternalArgsForCall(i int) (context.Context, signa.KeyOpts, options.RegistryOptions, map[string]interface{}, []string, string, bool, string, string, string, bool, bool, string) {
	fake.signImageInternalMutex.RLock()
	defer fake.signImageInternalMutex.RUnlock()
	argsForCall := fake.signImageInternalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8, argsForCall.arg9, argsForCall.arg10, argsForCall.arg11, argsForCall.arg12, argsForCall.arg13
}

func (fake *FakeImpl) SignImageInternalReturns(result1 error) {
	fake.signImageInternalMutex.Lock()
	defer fake.signImageInternalMutex.Unlock()
	fake.SignImageInternalStub = nil
	fake.signImageInternalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SignImageInternalReturnsOnCall(i int, result1 error) {
	fake.signImageInternalMutex.Lock()
	defer fake.signImageInternalMutex.Unlock()
	fake.SignImageInternalStub = nil
	if fake.signImageInternalReturnsOnCall == nil {
		fake.signImageInternalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.signImageInternalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) VerifyFileInternal(arg1 *sign.Signer, arg2 string) (*sign.SignedObject, error) {
	fake.verifyFileInternalMutex.Lock()
	ret, specificReturn := fake.verifyFileInternalReturnsOnCall[len(fake.verifyFileInternalArgsForCall)]
	fake.verifyFileInternalArgsForCall = append(fake.verifyFileInternalArgsForCall, struct {
		arg1 *sign.Signer
		arg2 string
	}{arg1, arg2})
	stub := fake.VerifyFileInternalStub
	fakeReturns := fake.verifyFileInternalReturns
	fake.recordInvocation("VerifyFileInternal", []interface{}{arg1, arg2})
	fake.verifyFileInternalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) VerifyFileInternalCallCount() int {
	fake.verifyFileInternalMutex.RLock()
	defer fake.verifyFileInternalMutex.RUnlock()
	return len(fake.verifyFileInternalArgsForCall)
}

func (fake *FakeImpl) VerifyFileInternalCalls(stub func(*sign.Signer, string) (*sign.SignedObject, error)) {
	fake.verifyFileInternalMutex.Lock()
	defer fake.verifyFileInternalMutex.Unlock()
	fake.VerifyFileInternalStub = stub
}

func (fake *FakeImpl) VerifyFileInternalArgsForCall(i int) (*sign.Signer, string) {
	fake.verifyFileInternalMutex.RLock()
	defer fake.verifyFileInternalMutex.RUnlock()
	argsForCall := fake.verifyFileInternalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) VerifyFileInternalReturns(result1 *sign.SignedObject, result2 error) {
	fake.verifyFileInternalMutex.Lock()
	defer fake.verifyFileInternalMutex.Unlock()
	fake.VerifyFileInternalStub = nil
	fake.verifyFileInternalReturns = struct {
		result1 *sign.SignedObject
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) VerifyFileInternalReturnsOnCall(i int, result1 *sign.SignedObject, result2 error) {
	fake.verifyFileInternalMutex.Lock()
	defer fake.verifyFileInternalMutex.Unlock()
	fake.VerifyFileInternalStub = nil
	if fake.verifyFileInternalReturnsOnCall == nil {
		fake.verifyFileInternalReturnsOnCall = make(map[int]struct {
			result1 *sign.SignedObject
			result2 error
		})
	}
	fake.verifyFileInternalReturnsOnCall[i] = struct {
		result1 *sign.SignedObject
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) VerifyImageInternal(arg1 context.Context, arg2 string, arg3 []string) (*sign.SignedObject, error) {
	var arg3Copy []string
	if arg3 != nil {
		arg3Copy = make([]string, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.verifyImageInternalMutex.Lock()
	ret, specificReturn := fake.verifyImageInternalReturnsOnCall[len(fake.verifyImageInternalArgsForCall)]
	fake.verifyImageInternalArgsForCall = append(fake.verifyImageInternalArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []string
	}{arg1, arg2, arg3Copy})
	stub := fake.VerifyImageInternalStub
	fakeReturns := fake.verifyImageInternalReturns
	fake.recordInvocation("VerifyImageInternal", []interface{}{arg1, arg2, arg3Copy})
	fake.verifyImageInternalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) VerifyImageInternalCallCount() int {
	fake.verifyImageInternalMutex.RLock()
	defer fake.verifyImageInternalMutex.RUnlock()
	return len(fake.verifyImageInternalArgsForCall)
}

func (fake *FakeImpl) VerifyImageInternalCalls(stub func(context.Context, string, []string) (*sign.SignedObject, error)) {
	fake.verifyImageInternalMutex.Lock()
	defer fake.verifyImageInternalMutex.Unlock()
	fake.VerifyImageInternalStub = stub
}

func (fake *FakeImpl) VerifyImageInternalArgsForCall(i int) (context.Context, string, []string) {
	fake.verifyImageInternalMutex.RLock()
	defer fake.verifyImageInternalMutex.RUnlock()
	argsForCall := fake.verifyImageInternalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) VerifyImageInternalReturns(result1 *sign.SignedObject, result2 error) {
	fake.verifyImageInternalMutex.Lock()
	defer fake.verifyImageInternalMutex.Unlock()
	fake.VerifyImageInternalStub = nil
	fake.verifyImageInternalReturns = struct {
		result1 *sign.SignedObject
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) VerifyImageInternalReturnsOnCall(i int, result1 *sign.SignedObject, result2 error) {
	fake.verifyImageInternalMutex.Lock()
	defer fake.verifyImageInternalMutex.Unlock()
	fake.VerifyImageInternalStub = nil
	if fake.verifyImageInternalReturnsOnCall == nil {
		fake.verifyImageInternalReturnsOnCall = make(map[int]struct {
			result1 *sign.SignedObject
			result2 error
		})
	}
	fake.verifyImageInternalReturnsOnCall[i] = struct {
		result1 *sign.SignedObject
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.envDefaultMutex.RLock()
	defer fake.envDefaultMutex.RUnlock()
	fake.isImageSignedInternalMutex.RLock()
	defer fake.isImageSignedInternalMutex.RUnlock()
	fake.setenvMutex.RLock()
	defer fake.setenvMutex.RUnlock()
	fake.signImageInternalMutex.RLock()
	defer fake.signImageInternalMutex.RUnlock()
	fake.verifyFileInternalMutex.RLock()
	defer fake.verifyFileInternalMutex.RUnlock()
	fake.verifyImageInternalMutex.RLock()
	defer fake.verifyImageInternalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
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
