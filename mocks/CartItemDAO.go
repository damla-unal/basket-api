// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "basket-api/internal/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CartItemDAO is an autogenerated mock type for the CartItemDAO type
type CartItemDAO struct {
	mock.Mock
}

// DeleteCartItem provides a mock function with given fields: ctx, id, cartID, cartPrice
func (_m *CartItemDAO) DeleteCartItem(ctx context.Context, id int, cartID int, cartPrice int) error {
	ret := _m.Called(ctx, id, cartID, cartPrice)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int) error); ok {
		r0 = rf(ctx, id, cartID, cartPrice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCartItemByID provides a mock function with given fields: ctx, id
func (_m *CartItemDAO) GetCartItemByID(ctx context.Context, id int) (model.CartItem, error) {
	ret := _m.Called(ctx, id)

	var r0 model.CartItem
	if rf, ok := ret.Get(0).(func(context.Context, int) model.CartItem); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.CartItem)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCartItemsByCartID provides a mock function with given fields: ctx, cartID
func (_m *CartItemDAO) GetCartItemsByCartID(ctx context.Context, cartID int) ([]model.CartItem, error) {
	ret := _m.Called(ctx, cartID)

	var r0 []model.CartItem
	if rf, ok := ret.Get(0).(func(context.Context, int) []model.CartItem); ok {
		r0 = rf(ctx, cartID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.CartItem)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, cartID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCartItem provides a mock function with given fields: ctx, cartItemToUpdate, updatedCartPrice
func (_m *CartItemDAO) UpdateCartItem(ctx context.Context, cartItemToUpdate model.CartItem, updatedCartPrice int) error {
	ret := _m.Called(ctx, cartItemToUpdate, updatedCartPrice)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CartItem, int) error); ok {
		r0 = rf(ctx, cartItemToUpdate, updatedCartPrice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpsertCartItem provides a mock function with given fields: ctx, cartItemToUpsert, updatedCartPrice
func (_m *CartItemDAO) UpsertCartItem(ctx context.Context, cartItemToUpsert model.CartItem, updatedCartPrice int) error {
	ret := _m.Called(ctx, cartItemToUpsert, updatedCartPrice)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CartItem, int) error); ok {
		r0 = rf(ctx, cartItemToUpsert, updatedCartPrice)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCartItemDAO interface {
	mock.TestingT
	Cleanup(func())
}

// NewCartItemDAO creates a new instance of CartItemDAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCartItemDAO(t mockConstructorTestingTNewCartItemDAO) *CartItemDAO {
	mock := &CartItemDAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
