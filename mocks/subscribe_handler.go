// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"
	pubsub "openmyth/blockchain/pkg/iface/pubsub"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// SubscribeHandler is an autogenerated mock type for the SubscribeHandler type
type SubscribeHandler struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, topic, msg, tt
func (_m *SubscribeHandler) Execute(ctx context.Context, topic string, msg *pubsub.Pack, tt time.Time) {
	_m.Called(ctx, topic, msg, tt)
}

// NewSubscribeHandler creates a new instance of SubscribeHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSubscribeHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *SubscribeHandler {
	mock := &SubscribeHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}