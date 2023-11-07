// Code generated by mockery v1.0.0. DO NOT EDIT.

package integration

import (
	flowkit "github.com/onflow/flow-cli/flowkit"
	mock "github.com/stretchr/testify/mock"
)

// mockFlowState is an autogenerated mock type for the flowState type
type mockFlowState struct {
	mock.Mock
}

// GetCodeByName provides a mock function with given fields: name
func (_m *mockFlowState) GetCodeByName(name string) (string, error) {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsLoaded provides a mock function with given fields:
func (_m *mockFlowState) IsLoaded() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Load provides a mock function with given fields: configPath
func (_m *mockFlowState) Load(configPath string) error {
	ret := _m.Called(configPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(configPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reload provides a mock function with given fields:
func (_m *mockFlowState) Reload() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// getConfigPath provides a mock function with given fields:
func (_m *mockFlowState) getConfigPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// getState provides a mock function with given fields:
func (_m *mockFlowState) getState() *flowkit.State {
	ret := _m.Called()

	var r0 *flowkit.State
	if rf, ok := ret.Get(0).(func() *flowkit.State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flowkit.State)
		}
	}

	return r0
}
