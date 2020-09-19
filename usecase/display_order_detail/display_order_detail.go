package usecase

import (
	"context"
	"rest_api/domain"
	"time"
)

type customerOrder struct {
	repo    domain.CustomerOrderRepository
	timeout time.Duration
}

func NewCustomerOrder(r domain.CustomerOrderRepository, t time.Duration) domain.CustomerOrderUsecase {
	return &customerOrder{
		repo:    r,
		timeout: t,
	}
}

func (s *customerOrder) GetOrdersByCustomer(ctx context.Context, name string) (res domain.CustomerOrder, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err = s.repo.GetOrdersByCustomer(ctx, name)
	if err != nil {
		return domain.CustomerOrder{}, err
	}
	return
}
