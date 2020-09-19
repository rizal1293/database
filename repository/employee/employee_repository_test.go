package repository_test

import (
	"context"
	"database/sql/driver"
	"rest_api/domain"
	repo "rest_api/repository/employee"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSave(t *testing.T) {
	var mockEmployee []*domain.Employee

	mockEmployee = append(mockEmployee, &domain.Employee{
		EmployeeID: 1,
		FirstName:  "Jhon",
		LastName:   "Doe",
		Title:      "Software Engineer",
		WorkPhone:  "08111111111",
	})

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	valueArgs := []driver.Value{}

	for _, row := range mockEmployee {
		valueArgs = append(valueArgs, row.EmployeeID)
		valueArgs = append(valueArgs, row.FirstName)
		valueArgs = append(valueArgs, row.LastName)
		valueArgs = append(valueArgs, row.Title)
		valueArgs = append(valueArgs, row.WorkPhone)
	}

	mock.ExpectBegin()
	prep := mock.ExpectExec("INSERT INTO Employees")
	prep.WithArgs(valueArgs...).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := repo.NewEmployeeRepository(db)
	err = r.Save(context.TODO(), mockEmployee)
	lastID := mockEmployee[len(mockEmployee)-1].EmployeeID

	assert.NoError(t, err)
	assert.Equal(t, int64(1), lastID)
}
