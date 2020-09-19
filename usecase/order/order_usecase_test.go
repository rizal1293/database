package usecase_test

import (
	"context"
	"rest_api/domain"
	"rest_api/domain/mocks"
	usecase "rest_api/usecase/order"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSave(t *testing.T) {
	mockRepo := new(mocks.OrderRepository)

	var mockOrder []*domain.Order

	mockOrder = append(mockOrder, &domain.Order{
		OrderID: 1,
		Customer: domain.Customer{
			CustomerID: 1,
		},
		Employee: domain.Employee{
			EmployeeID: 1,
		},
		OrderDate:           "05/05/2020",
		PurchaseOrderNumber: "1",
		ShipDate:            "12/05/2021",
		ShippingMethod: domain.ShippingMethod{
			ShippingMethodID: 1,
		},
		FreightCharge:   1,
		Taxes:           1,
		PaymentReceived: 1,
		Comment:         "Success",
	})

	t.Run("success", func(t *testing.T) {
		tempMockOrder := mockOrder
		mockRepo.On("Save", mock.Anything, mock.AnythingOfType("[]*domain.Order")).Return(nil).Once()

		u := usecase.NewOrderUsecase(mockRepo, time.Second*2)

		err := u.Save(context.TODO(), tempMockOrder)

		assert.NoError(t, err)
		assert.Equal(t, len(mockOrder), len(tempMockOrder))
		mockRepo.AssertExpectations(t)
	})
}
