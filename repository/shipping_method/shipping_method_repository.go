package repository

import (
	"context"
	"database/sql"
	"fmt"
	"rest_api/domain"
	"strings"
)

type shippingMethodRepository struct {
	Conn *sql.DB
}

func NewShippingMethodRepository(Conn *sql.DB) domain.ShippingMethodRepository {
	return &shippingMethodRepository{Conn}
}

func (db *shippingMethodRepository) Save(ctx context.Context, sm []*domain.ShippingMethod) (err error) {
	valueArgs := []interface{}{}
	valueStrings := []string{}

	for _, col := range sm {
		valueStrings = append(valueStrings, "(?, ?)")

		valueArgs = append(valueArgs, col.ShippingMethodID)
		valueArgs = append(valueArgs, col.ShippingMethod)
	}

	query := `INSERT INTO Shipping_Methods (ShippingMethodID, ShippingMethod) VALUES %s`
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

	sm[len(sm)-1].ShippingMethodID = lastID

	return nil
}
