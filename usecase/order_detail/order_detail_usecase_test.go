package usecase_test

import (
	"context"
	"rest_api/domain"
	"rest_api/domain/mocks"
	usecase "rest_api/usecase/order_detail"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSave(t *testing.T) {
	mockRepo := new(mocks.OrderDetailRepository)

	var mockOrderDetail []*domain.OrderDetail

	mockOrderDetail = append(mockOrderDetail, &domain.OrderDetail{
		OrderDetailID: 1,
		Order: domain.Order{
			OrderID: 1,
		},
		Product: domain.Product{
			ProductID: 1,
		},
		Quantity:  1,
		UnitPrice: 1,
		Discount:  2,
	})

	t.Run("success", func(t *testing.T) {
		tempMockOrderDetail := mockOrderDetail
		mockRepo.On("Save", mock.Anything, mock.AnythingOfType("[]*domain.OrderDetail")).Return(nil).Once()

		u := usecase.NewOrderDetailUsecase(mockRepo, time.Second*2)

		err := u.Save(context.TODO(), tempMockOrderDetail)

		assert.NoError(t, err)
		assert.Equal(t, len(mockOrderDetail), len(tempMockOrderDetail))
		mockRepo.AssertExpectations(t)
	})
}
