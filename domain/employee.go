package domain

import "context"

type Employee struct {
	EmployeeID int64
	FirstName  string
	LastName   string
	Title      string
	WorkPhone  string
}

type EmployeeRepository interface {
	Save(ctx context.Context, e []*Employee) error
}

type EmployeeUsecase interface {
	Save(ctx context.Context, e []*Employee) error
}
