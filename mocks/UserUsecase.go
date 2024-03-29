// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	mock "github.com/stretchr/testify/mock"

	model "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0, _a1
func (_m *UserUsecase) CreateUser(_a0 context.Context, _a1 dto.RegisterReq) (*dto.RegisterRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dto.RegisterRes
	if rf, ok := ret.Get(0).(func(context.Context, dto.RegisterReq) *dto.RegisterRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.RegisterRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.RegisterReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields: _a0
func (_m *UserUsecase) GetAllUsers(_a0 context.Context) ([]model.User, error) {
	ret := _m.Called(_a0)

	var r0 []model.User
	if rf, ok := ret.Get(0).(func(context.Context) []model.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserDetails provides a mock function with given fields: _a0, _a1
func (_m *UserUsecase) GetUserDetails(_a0 context.Context, _a1 uint) (*dto.UserDetails, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dto.UserDetails
	if rf, ok := ret.Get(0).(func(context.Context, uint) *dto.UserDetails); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UserDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserLogin provides a mock function with given fields: _a0, _a1
func (_m *UserUsecase) UserLogin(_a0 context.Context, _a1 dto.LoginReq) (*dto.LoginRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *dto.LoginRes
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginReq) *dto.LoginRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.LoginRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, dto.LoginReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
