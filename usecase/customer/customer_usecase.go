package usecase

import (
	"context"
	"rest_api/domain"
	"time"
)

type customerUsecase struct {
	customerRepo domain.CustomerRepository
	timeout      time.Duration
}

func NewCustomerUsecase(r domain.CustomerRepository, t time.Duration) domain.CustomerUsecase {
	return &customerUsecase{
		customerRepo: r,
		timeout:      t,
	}
}

func (s *customerUsecase) Save(ctx context.Context, c []*domain.Customer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	err = s.customerRepo.Save(ctx, c)
	return
}
