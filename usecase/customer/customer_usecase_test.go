package usecase_test

import (
	"context"
	"rest_api/domain"
	"rest_api/domain/mocks"
	usecase "rest_api/usecase/customer"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSave(t *testing.T) {
	mockRepo := new(mocks.CustomerRepository)

	var mockCustomer []*domain.Customer

	mockCustomer = append(mockCustomer, &domain.Customer{
		CustomerID:          1,
		CompanyName:         "XYZ",
		FisrtName:           "Jhon",
		LastName:            "Dhoe",
		BillingAddress:      "Bandung",
		City:                "Bandung",
		StateOrProvince:     "Jawa Barat",
		ZIPCode:             "46344",
		Email:               "xyz@gmail.com",
		CompanyWebsite:      "xyz.com",
		PhoneNumber:         "081122333778",
		FaxNumber:           "1234",
		ShipAddress:         "Bandung",
		ShipCity:            "Bandung",
		ShipStateOrProvince: "Jawa Barat",
		ShipZIPCode:         "46344",
		ShipPhoneNumber:     "08111111111",
	})

	t.Run("success", func(t *testing.T) {
		tempMockCustomer := mockCustomer
		mockRepo.On("Save", mock.Anything, mock.AnythingOfType("[]*domain.Customer")).Return(nil).Once()

		u := usecase.NewCustomerUsecase(mockRepo, time.Second*2)

		err := u.Save(context.TODO(), tempMockCustomer)

		assert.NoError(t, err)
		assert.Equal(t, len(mockCustomer), len(tempMockCustomer))
		mockRepo.AssertExpectations(t)
	})
}
