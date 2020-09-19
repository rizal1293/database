package repository

import (
	"context"
	"database/sql"
	"fmt"
	"rest_api/domain"
	"strings"
)

type orderDetailRepository struct {
	Conn *sql.DB
}

func NewOrderDetailRepository(Conn *sql.DB) domain.OrderDetailRepository {
	return &orderDetailRepository{Conn}
}

func (db *orderDetailRepository) Save(ctx context.Context, od []*domain.OrderDetail) (er error) {
	valueArgs := []interface{}{}
	valueStrings := []string{}

	for _, col := range od {

		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?)")

		valueArgs = append(valueArgs, col.OrderDetailID)
		valueArgs = append(valueArgs, col.Order.OrderID)
		valueArgs = append(valueArgs, col.Product.ProductID)
		valueArgs = append(valueArgs, col.Quantity)
		valueArgs = append(valueArgs, col.UnitPrice)
		valueArgs = append(valueArgs, col.Discount)
	}

	query := "INSERT INTO Order_Details (OrderDetailID, OrderID, ProductID, Quantity, UnitPrice, Discount) VALUES %s"
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

	od[len(od)-1].OrderDetailID = lastID

	return nil
}
