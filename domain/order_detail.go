package domain

import "context"

type OrderDetail struct {
	OrderDetailID int64   
	Order         Order   
	Product       Product 
	Quantity      int     
	UnitPrice     int     
	Discount      int     
}

type OrderDetailRepository interface {
	Save(ctx context.Context, od []*OrderDetail) error
}

type OrderDetailUsecase interface {
	Save(ctx context.Context, od []*OrderDetail) error
}
