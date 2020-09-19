package repository_test

import (
	"context"
	"database/sql/driver"
	"rest_api/domain"
	repo "rest_api/repository/order"
	util "rest_api/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSave(t *testing.T) {
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	valueArgs := []driver.Value{}

	for _, row := range mockOrder {
		valueArgs = append(valueArgs, row.OrderID)
		valueArgs = append(valueArgs, row.Customer.CustomerID)
		valueArgs = append(valueArgs, row.Employee.EmployeeID)
		valueArgs = append(valueArgs, util.ParserToDateYYYMMDD(row.OrderDate))
		valueArgs = append(valueArgs, row.PurchaseOrderNumber)
		valueArgs = append(valueArgs, util.ParserToDateYYYMMDD(row.ShipDate))
		valueArgs = append(valueArgs, row.ShippingMethod.ShippingMethodID)
		valueArgs = append(valueArgs, row.FreightCharge)
		valueArgs = append(valueArgs, row.Taxes)
		valueArgs = append(valueArgs, row.PaymentReceived)
		valueArgs = append(valueArgs, row.Comment)
	}

	mock.ExpectBegin()
	prep := mock.ExpectExec("INSERT INTO Orders")
	prep.WithArgs(valueArgs...).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := repo.NewOrderRepository(db)
	err = r.Save(context.TODO(), mockOrder)
	lastID := mockOrder[len(mockOrder)-1].OrderID

	assert.NoError(t, err)
	assert.Equal(t, int64(1), lastID)
}
