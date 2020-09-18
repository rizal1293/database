package domain

type Order struct {
	OrderID             int64
	Costumer            Costumer
	Employee            Employee
	PurchaseOrderNumber string
	ShipDate            string
	ShippingMethod      ShippingMethod
	FreightCharge       int
	Taxes               int
	PaymentReceived     int
	Comment             string
}
