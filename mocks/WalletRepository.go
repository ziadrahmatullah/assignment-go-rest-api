// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	mock "github.com/stretchr/testify/mock"
)

// WalletRepository is an autogenerated mock type for the WalletRepository type
type WalletRepository struct {
	mock.Mock
}

// FindWalletByUserId provides a mock function with given fields: _a0, _a1
func (_m *WalletRepository) FindWalletByUserId(_a0 context.Context, _a1 uint) (*model.Wallet, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, uint) *model.Wallet); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
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

// NewWallet provides a mock function with given fields: _a0, _a1
func (_m *WalletRepository) NewWallet(_a0 context.Context, _a1 uint) (*model.Wallet, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, uint) *model.Wallet); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Wallet)
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

type mockConstructorTestingTNewWalletRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewWalletRepository creates a new instance of WalletRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWalletRepository(t mockConstructorTestingTNewWalletRepository) *WalletRepository {
	mock := &WalletRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
