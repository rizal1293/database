package repository

import (
	"context"
	"database/sql"
	"fmt"
	"rest_api/domain"
	"strings"

	"github.com/sirupsen/logrus"
)

type orderRepository struct {
	Conn *sql.DB
}

func NewOrderRepository(Conn *sql.DB) domain.OrderRepository {
	return &orderRepository{Conn}
}

func (db *orderRepository) Save(ctx context.Context, o []*domain.Order) (err error) {
	valueArgs := []interface{}{}
	valueStrings := []string{}

	for _, row := range o {

		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?,?, ?, ?, ?, ?, ?)")

		valueArgs = append(valueArgs, row.OrderID)
		valueArgs = append(valueArgs, row.Customer.CustomerID)
		valueArgs = append(valueArgs, row.Employee.EmployeeID)
		valueArgs = append(valueArgs, row.OrderDate)
		valueArgs = append(valueArgs, row.PurchaseOrderNumber)
		valueArgs = append(valueArgs, row.ShipDate)
		valueArgs = append(valueArgs, row.ShippingMethod.ShippingMethodID)
		valueArgs = append(valueArgs, row.FreightCharge)
		valueArgs = append(valueArgs, row.Taxes)
		valueArgs = append(valueArgs, row.PaymentReceived)
		valueArgs = append(valueArgs, row.Comment)
	}

	query := `INSERT INTO Orders (OrderID, CustomerID, EmployeeID, OrderDate, PurchaseOrderNumber, ShipDate, ShippingMethodID, FreightCharge, Taxes, PaymentReceived, Comment) VALUES %s`
	query = fmt.Sprintf(query, strings.Join(valueStrings, ","))

	tx, err := db.Conn.Begin()
	if err != nil {
		return err
	}

	res, err := tx.Exec(query, valueArgs...)

	if err != nil {
		tx.Rollback()
		return err
	}

	lastID, err := res.LastInsertId()

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	o[len(o)-1].OrderID = lastID

	return nil
}
