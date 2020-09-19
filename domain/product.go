package domain

import "context"

type Product struct {
	ProductID int64
	ProductName string
	UntiPrice int
	InStock   string
}

type ProductRepository interface {
	Save(ctx context.Context, p []*Product) error
}

type ProductUsecase interface {
	Save(ctx context.Context, p []*Product) error
}
