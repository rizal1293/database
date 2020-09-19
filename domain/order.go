package domain

import "context"

type Order struct {
	OrderID             int64
	Customer            Customer
	Employee            Employee
	OrderDate           string
	PurchaseOrderNumber string
	ShipDate            string
	ShippingMethod      ShippingMethod
	FreightCharge       int
	Taxes               int
	PaymentReceived     int
	Comment             string
}

type OrderRepository interface {
	Save(ctx context.Context, o []*Order) error
}

type OrderUsecase interface {
	Save(ctx context.Context, o []*Order) error
}
