// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// MapMetaDataFunc is an autogenerated mock type for the MapMetaDataFunc type
type MapMetaDataFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *MapMetaDataFunc) Execute(_a0 context.Context, _a1 *http.Request) metadata.MD {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) metadata.MD); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// NewMapMetaDataFunc creates a new instance of MapMetaDataFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMapMetaDataFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MapMetaDataFunc {
	mock := &MapMetaDataFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
