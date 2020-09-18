package domain

type OrderDetail struct {
	OrderDetailID int64
	Order         Order
	Product       Product
	Quantity      int
	UnitPrice     int
	Discount      int
}
