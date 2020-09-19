package domain

import "context"

type ShippingMethod struct {
	ShippingMethodID int64
	ShippingMethod   string
}

type ShippingMethodRepository interface {
	Save(ctx context.Context, sm []*ShippingMethod) error
}

type ShippingMethodUsecase interface {
	Save(ctx context.Context, sm []*ShippingMethod) error
}
