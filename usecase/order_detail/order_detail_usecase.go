package usecase

import (
	"context"
	"rest_api/domain"
	"time"
)

type orderDetailUsecase struct {
	orderDetailRepo domain.OrderDetailRepository
	timeout         time.Duration
}

func NewOrderDetailUsecase(r domain.OrderDetailRepository, t time.Duration) domain.OrderDetailUsecase {
	return &orderDetailUsecase{
		orderDetailRepo: r,
		timeout:         t,
	}
}

func (s *orderDetailUsecase) Save(ctx context.Context, od []*domain.OrderDetail) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	err = s.orderDetailRepo.Save(ctx, od)
	return
}
