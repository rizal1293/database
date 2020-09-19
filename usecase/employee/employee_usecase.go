package usecase

import (
	"context"
	"rest_api/domain"
	"time"
)

type employeeUsecase struct {
	employeRepo domain.EmployeeRepository
	timeout     time.Duration
}

func NewEmployeeUsecase(r domain.EmployeeRepository, t time.Duration) domain.EmployeeUsecase {
	return &employeeUsecase{
		employeRepo: r,
		timeout:     t,
	}
}

func (s *employeeUsecase) Save(ctx context.Context, e []*domain.Employee) (err error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	err = s.employeRepo.Save(ctx, e)
	return
}
