package repository_test

import (
	"context"
	"database/sql/driver"
	"rest_api/domain"
	repo "rest_api/repository/customer"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSave(t *testing.T) {
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	valueArgs := []driver.Value{}

	for _, row := range mockCustomer {

		valueArgs = append(valueArgs, row.CustomerID)
		valueArgs = append(valueArgs, row.CompanyName)
		valueArgs = append(valueArgs, row.FisrtName)
		valueArgs = append(valueArgs, row.LastName)
		valueArgs = append(valueArgs, row.BillingAddress)
		valueArgs = append(valueArgs, row.City)
		valueArgs = append(valueArgs, row.StateOrProvince)
		valueArgs = append(valueArgs, row.ZIPCode)
		valueArgs = append(valueArgs, row.Email)
		valueArgs = append(valueArgs, row.CompanyWebsite)
		valueArgs = append(valueArgs, row.PhoneNumber)
		valueArgs = append(valueArgs, row.FaxNumber)
		valueArgs = append(valueArgs, row.ShipAddress)
		valueArgs = append(valueArgs, row.ShipCity)
		valueArgs = append(valueArgs, row.ShipStateOrProvince)
		valueArgs = append(valueArgs, row.ShipZIPCode)
		valueArgs = append(valueArgs, row.ShipPhoneNumber)
	}

	mock.ExpectBegin()
	prep := mock.ExpectExec("INSERT INTO Customers")
	prep.WithArgs(valueArgs...).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := repo.NewCustomerRepository(db)
	err = r.Save(context.TODO(), mockCustomer)
	lastID := mockCustomer[len(mockCustomer)-1].CustomerID

	assert.NoError(t, err)
	assert.Equal(t, int64(1), lastID)
}
