package repository_test

import (
	"context"
	"database/sql/driver"
	"rest_api/domain"
	repo "rest_api/repository/order_detail"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSave(t *testing.T) {
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

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	valueArgs := []driver.Value{}

	for _, col := range mockOrderDetail {
		valueArgs = append(valueArgs, col.OrderDetailID)
		valueArgs = append(valueArgs, col.Order.OrderID)
		valueArgs = append(valueArgs, col.Product.ProductID)
		valueArgs = append(valueArgs, col.Quantity)
		valueArgs = append(valueArgs, col.UnitPrice)
		valueArgs = append(valueArgs, col.Discount)
	}

	mock.ExpectBegin()
	prep := mock.ExpectExec("INSERT INTO Order_Details")
	prep.WithArgs(valueArgs...).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := repo.NewOrderDetailRepository(db)
	err = r.Save(context.TODO(), mockOrderDetail)
	lastID := mockOrderDetail[len(mockOrderDetail)-1].OrderDetailID

	assert.NoError(t, err)
	assert.Equal(t, int64(1), lastID)
}
