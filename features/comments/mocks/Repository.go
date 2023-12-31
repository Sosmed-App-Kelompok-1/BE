// Code generated by mockery v2.37.1. DO NOT EDIT.

package mocks

import (
	context "context"
	comments "sosmed/features/comments"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, data
func (_m *Repository) Create(ctx context.Context, data comments.Comment) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, comments.Comment) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, commentId
func (_m *Repository) Delete(ctx context.Context, commentId uint) error {
	ret := _m.Called(ctx, commentId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, commentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByPostId provides a mock function with given fields: ctx, postId
func (_m *Repository) DeleteByPostId(ctx context.Context, postId uint) error {
	ret := _m.Called(ctx, postId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, postId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByUserId provides a mock function with given fields: ctx, userId
func (_m *Repository) DeleteByUserId(ctx context.Context, userId uint) error {
	ret := _m.Called(ctx, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
