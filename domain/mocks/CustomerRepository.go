// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "rest_api/domain"

	mock "github.com/stretchr/testify/mock"
)

// CustomerRepository is an autogenerated mock type for the CustomerRepository type
type CustomerRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: ctx, c
func (_m *CustomerRepository) Save(ctx context.Context, c []*domain.Customer) error {
	ret := _m.Called(ctx, c)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*domain.Customer) error); ok {
		r0 = rf(ctx, c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
