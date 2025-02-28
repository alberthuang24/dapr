/*
Copyright 2021 The Dapr Authors
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

package testing

import "github.com/dapr/components-contrib/state"

type TransactionalStoreMock struct {
	MockStateStore
}

// Code generated by mockery v2.3.0. DO NOT EDIT.
func (_m *TransactionalStoreMock) Multi(request *state.TransactionalStateRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*state.TransactionalStateRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *TransactionalStoreMock) Features() []state.Feature {
	return []state.Feature{state.FeatureTransactional}
}

func (_m *TransactionalStoreMock) Close() error {
	return nil
}
