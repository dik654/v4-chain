// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	bridgetypes "github.com/dydxprotocol/v4/x/bridge/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/cosmos/cosmos-sdk/types"
)

// BridgeKeeper is an autogenerated mock type for the BridgeKeeper type
type BridgeKeeper struct {
	mock.Mock
}

// AcknowledgeBridges provides a mock function with given fields: ctx, bridges
func (_m *BridgeKeeper) AcknowledgeBridges(ctx types.Context, bridges []bridgetypes.BridgeEvent) error {
	ret := _m.Called(ctx, bridges)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, []bridgetypes.BridgeEvent) error); ok {
		r0 = rf(ctx, bridges)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompleteBridge provides a mock function with given fields: ctx, bridges
func (_m *BridgeKeeper) CompleteBridge(ctx types.Context, bridges bridgetypes.BridgeEvent) error {
	ret := _m.Called(ctx, bridges)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, bridgetypes.BridgeEvent) error); ok {
		r0 = rf(ctx, bridges)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAcknowledgedEventInfo provides a mock function with given fields: ctx
func (_m *BridgeKeeper) GetAcknowledgedEventInfo(ctx types.Context) bridgetypes.BridgeEventInfo {
	ret := _m.Called(ctx)

	var r0 bridgetypes.BridgeEventInfo
	if rf, ok := ret.Get(0).(func(types.Context) bridgetypes.BridgeEventInfo); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bridgetypes.BridgeEventInfo)
	}

	return r0
}

// GetEventParams provides a mock function with given fields: ctx
func (_m *BridgeKeeper) GetEventParams(ctx types.Context) bridgetypes.EventParams {
	ret := _m.Called(ctx)

	var r0 bridgetypes.EventParams
	if rf, ok := ret.Get(0).(func(types.Context) bridgetypes.EventParams); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bridgetypes.EventParams)
	}

	return r0
}

// GetProposeParams provides a mock function with given fields: ctx
func (_m *BridgeKeeper) GetProposeParams(ctx types.Context) bridgetypes.ProposeParams {
	ret := _m.Called(ctx)

	var r0 bridgetypes.ProposeParams
	if rf, ok := ret.Get(0).(func(types.Context) bridgetypes.ProposeParams); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bridgetypes.ProposeParams)
	}

	return r0
}

// GetRecognizedEventInfo provides a mock function with given fields: ctx
func (_m *BridgeKeeper) GetRecognizedEventInfo(ctx types.Context) bridgetypes.BridgeEventInfo {
	ret := _m.Called(ctx)

	var r0 bridgetypes.BridgeEventInfo
	if rf, ok := ret.Get(0).(func(types.Context) bridgetypes.BridgeEventInfo); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bridgetypes.BridgeEventInfo)
	}

	return r0
}

// GetSafetyParams provides a mock function with given fields: ctx
func (_m *BridgeKeeper) GetSafetyParams(ctx types.Context) bridgetypes.SafetyParams {
	ret := _m.Called(ctx)

	var r0 bridgetypes.SafetyParams
	if rf, ok := ret.Get(0).(func(types.Context) bridgetypes.SafetyParams); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bridgetypes.SafetyParams)
	}

	return r0
}

// UpdateEventParams provides a mock function with given fields: ctx, params
func (_m *BridgeKeeper) UpdateEventParams(ctx types.Context, params bridgetypes.EventParams) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, bridgetypes.EventParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProposeParams provides a mock function with given fields: ctx, params
func (_m *BridgeKeeper) UpdateProposeParams(ctx types.Context, params bridgetypes.ProposeParams) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, bridgetypes.ProposeParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSafetyParams provides a mock function with given fields: ctx, params
func (_m *BridgeKeeper) UpdateSafetyParams(ctx types.Context, params bridgetypes.SafetyParams) error {
	ret := _m.Called(ctx, params)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, bridgetypes.SafetyParams) error); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBridgeKeeper interface {
	mock.TestingT
	Cleanup(func())
}

// NewBridgeKeeper creates a new instance of BridgeKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBridgeKeeper(t mockConstructorTestingTNewBridgeKeeper) *BridgeKeeper {
	mock := &BridgeKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
