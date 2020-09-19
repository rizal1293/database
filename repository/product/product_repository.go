package repository

import (
	"context"
	"database/sql"
	"fmt"
	"rest_api/domain"
	"strings"
)

type productRepository struct {
	Conn *sql.DB
}

func NewProductRepository(Conn *sql.DB) domain.ProductRepository {
	return &productRepository{Conn}
}

func (db *productRepository) Save(ctx context.Context, p []*domain.Product) (err error) {
	valueArgs := []interface{}{}
	valueStrings := []string{}

	for _, col := range p {
		valueStrings = append(valueStrings, "(?, ?, ?, ?)")

		valueArgs = append(valueArgs, col.ProductID)
		valueArgs = append(valueArgs, col.ProductName)
		valueArgs = append(valueArgs, col.UntiPrice)
		valueArgs = append(valueArgs, col.InStock)
	}

	query := `INSERT INTO Products (ProductID, ProductName, UnitPrice, InStock) VALUES %s`
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

	p[len(p)-1].ProductID = lastID

	return nil
}
