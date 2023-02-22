// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EventConsumer is an autogenerated mock type for the EventConsumer type
type EventConsumer struct {
	mock.Mock
}

// Consume provides a mock function with given fields: groupID
func (_m *EventConsumer) Consume(groupID string) ([]byte, error) {
	ret := _m.Called(groupID)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(groupID)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewEventConsumer interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventConsumer creates a new instance of EventConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventConsumer(t mockConstructorTestingTNewEventConsumer) *EventConsumer {
	mock := &EventConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
