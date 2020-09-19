package usecase

import (
	"context"
	"rest_api/domain"
	"time"
)

type productUsecase struct {
	productRepo domain.ProductRepository
	timeout     time.Duration
}

func NewProductUsecase(r domain.ProductRepository, t time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepo: r,
		timeout:     t,
	}
}

func (s *productUsecase) Save(ctx context.Context, p []*domain.Product) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	err = s.productRepo.Save(ctx, p)
	return
}
