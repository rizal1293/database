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
	Product []DetailOrderProduct
	Total int
}

type DetailOrderProduct struct {
	Product string
	UnitPrice int
	Quantity int
	Discount int
	Subtotal int
}

type CustomerOrderRepository interface {
	GetOrdersByCustomer(ctx context.Context, name string) (CustomerOrder, error)
	// GetDetailOrder(ctx context.Context, orderID int64) ([]DetailOrderProduct, error)
}

type CustomerOrderUsecase interface {
	GetOrdersByCustomer(ctx context.Context, name string) (CustomerOrder, error)
	// GetDetailOrder(ctx context.Context, orderID int64) ([]DetailOrderProduct, error)
}