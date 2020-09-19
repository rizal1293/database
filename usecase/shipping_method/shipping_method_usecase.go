package usecase

import (
	"context"
	"rest_api/domain"
	"time"
)

type shippingMethodRepository struct {
	shipMethodRepo domain.ShippingMethodRepository
	timeout time.Duration
}

func NewShippingMethodRepository(r domain.ShippingMethodRepository, t time.Duration) domain.ShippingMethodUsecase {
	return &shippingMethodRepository{
		shipMethodRepo: r,
		timeout: t,
	}
}

func (s *shippingMethodRepository) Save(ctx context.Context, sm []*domain.ShippingMethod) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	err = s.shipMethodRepo.Save(ctx, sm)
	return
}
