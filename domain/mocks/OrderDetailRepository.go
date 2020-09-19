// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "rest_api/domain"

	mock "github.com/stretchr/testify/mock"
)

// OrderDetailRepository is an autogenerated mock type for the OrderDetailRepository type
type OrderDetailRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: ctx, od
func (_m *OrderDetailRepository) Save(ctx context.Context, od []*domain.OrderDetail) error {
	ret := _m.Called(ctx, od)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*domain.OrderDetail) error); ok {
		r0 = rf(ctx, od)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
