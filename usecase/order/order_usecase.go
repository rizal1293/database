package usecase

import (
	"context"
	"rest_api/domain"
	"time"
)

type orderUsecase struct {
	orderRepo domain.OrderRepository
	timeout   time.Duration
}

func NewOrderUsecase(r domain.OrderRepository, t time.Duration) domain.OrderUsecase {
	return &orderUsecase{
		orderRepo: r,
		timeout:   t,
	}
}

func (s *orderUsecase) Save(ctx context.Context, o []*domain.Order) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	err = s.orderRepo.Save(ctx, o)
	return
}
