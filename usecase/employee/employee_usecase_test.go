package usecase_test

import (
	"context"
	"rest_api/domain"
	"rest_api/domain/mocks"
	usecase "rest_api/usecase/employee"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSave(t *testing.T) {
	mockRepo := new(mocks.EmployeeRepository)

	var mockEmployee []*domain.Employee

	mockEmployee = append(mockEmployee, &domain.Employee{
		EmployeeID: 1,
		FirstName:  "Jhon",
		LastName:   "Doe",
		Title:      "Software Engineer",
		WorkPhone:  "08111111111",
	})

	t.Run("success", func(t *testing.T) {
		tempMockEmployee := mockEmployee
		mockRepo.On("Save", mock.Anything, mock.AnythingOfType("[]*domain.Employee")).Return(nil).Once()

		u := usecase.NewEmployeeUsecase(mockRepo, time.Second*2)

		err := u.Save(context.TODO(), tempMockEmployee)

		assert.NoError(t, err)
		assert.Equal(t, len(mockEmployee), len(tempMockEmployee))
		mockRepo.AssertExpectations(t)
	})
}
