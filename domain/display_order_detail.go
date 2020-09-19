package domain

import "context"

type CustomerOrder struct {
	CustomerName string
	Orders []CustomerOrderDetail
}

type CustomerOrderDetail struct {
	OrderID int64
	EmployeeName string
	ShippingMethod string
}

type CustomerOrderRepository interface {
	GetOrdersByCustomer(ctx context.Context, name string) (CustomerOrder, error)
}

type CustomerOrderUsecase interface {
	GetOrdersByCustomer(ctx context.Context, name string) (CustomerOrder, error)
}