package repository_test

import (
	"context"
	"database/sql/driver"
	"rest_api/domain"
	repo "rest_api/repository/shipping_method"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSave(t *testing.T) {
	var mockShippingMethod []*domain.ShippingMethod

	mockShippingMethod = append(mockShippingMethod, &domain.ShippingMethod{
		ShippingMethodID: 1,
		ShippingMethod:   "COD",
	})

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	valueArgs := []driver.Value{}

	for _, col := range mockShippingMethod {
		valueArgs = append(valueArgs, col.ShippingMethodID)
		valueArgs = append(valueArgs, col.ShippingMethod)
	}

	mock.ExpectBegin()
	prep := mock.ExpectExec("INSERT INTO Shipping_Methods")
	prep.WithArgs(valueArgs...).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	r := repo.NewShippingMethodRepository(db)
	err = r.Save(context.TODO(), mockShippingMethod)
	lastID := mockShippingMethod[len(mockShippingMethod)-1].ShippingMethodID

	assert.NoError(t, err)
	assert.Equal(t, int64(1), lastID)
}
