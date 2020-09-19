package repository_test

import (
	"context"
	"database/sql/driver"
	"rest_api/domain"
	repo "rest_api/repository/product"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSave(t *testing.T) {
	var mockProduct []*domain.Product

	mockProduct = append(mockProduct, &domain.Product{
		ProductID:   1,
		ProductName: "Banana",
		UntiPrice:   1,
		InStock:     "1",
	})

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	valueArgs := []driver.Value{}

	for _, col := range mockProduct {
		valueArgs = append(valueArgs, col.ProductID)
		valueArgs = append(valueArgs, col.ProductName)
		valueArgs = append(valueArgs, col.UntiPrice)
		valueArgs = append(valueArgs, col.InStock)
	}

	mock.ExpectBegin()
	prep := mock.ExpectExec("INSERT INTO Products")
	prep.WithArgs(valueArgs...).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := repo.NewProductRepository(db)
	err = r.Save(context.TODO(), mockProduct)
	lastID := mockProduct[len(mockProduct)-1].ProductID

	assert.NoError(t, err)
	assert.Equal(t, int64(1), lastID)
}
