package usecase_test

import (
	"context"
	"rest_api/domain"
	"rest_api/domain/mocks"
	usecase "rest_api/usecase/shipping_method"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSave(t *testing.T) {
	mockRepo := new(mocks.ShippingMethodRepository)

	var mockShippingMethod []*domain.ShippingMethod

	mockShippingMethod = append(mockShippingMethod, &domain.ShippingMethod{
		ShippingMethodID: 1,
		ShippingMethod:   "COD",
	})

	t.Run("success", func(t *testing.T) {
		tempMockShippingMethod := mockShippingMethod
		mockRepo.On("Save", mock.Anything, mock.AnythingOfType("[]*domain.ShippingMethod")).Return(nil).Once()

		u := usecase.NewShippingMethodRepository(mockRepo, time.Second*2)

		err := u.Save(context.TODO(), tempMockShippingMethod)

		assert.NoError(t, err)
		assert.Equal(t, len(mockShippingMethod), len(tempMockShippingMethod))
		mockRepo.AssertExpectations(t)
	})
}
