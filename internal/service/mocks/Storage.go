// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// GetLink provides a mock function with given fields: ctx, shortenedLink
func (_m *Storage) GetLink(ctx context.Context, shortenedLink string) (string, error) {
	ret := _m.Called(ctx, shortenedLink)

	if len(ret) == 0 {
		panic("no return value specified for GetLink")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, shortenedLink)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, shortenedLink)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortenedLink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveLink provides a mock function with given fields: ctx, originalLink
func (_m *Storage) SaveLink(ctx context.Context, originalLink string) (string, error) {
	ret := _m.Called(ctx, originalLink)

	if len(ret) == 0 {
		panic("no return value specified for SaveLink")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, originalLink)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, originalLink)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, originalLink)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
