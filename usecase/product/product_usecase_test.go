package usecase_test

import (
	"context"
	"rest_api/domain"
	"rest_api/domain/mocks"
	usecase "rest_api/usecase/product"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSave(t *testing.T) {
	mockRepo := new(mocks.ProductRepository)

	var mockProduct []*domain.Product

	mockProduct = append(mockProduct, &domain.Product{
		ProductID: 1,
		UntiPrice: 1,
		InStock:   "1",
	})

	t.Run("success", func(t *testing.T) {
		tempMockProduct := mockProduct
		mockRepo.On("Save", mock.Anything, mock.AnythingOfType("[]*domain.Product")).Return(nil).Once()

		u := usecase.NewProductUsecase(mockRepo, time.Second*2)

		err := u.Save(context.TODO(), tempMockProduct)

		assert.NoError(t, err)
		assert.Equal(t, len(mockProduct), len(tempMockProduct))
		mockRepo.AssertExpectations(t)
	})
}
