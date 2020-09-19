package repository

import (
	"context"
	"database/sql"
	"fmt"
	"rest_api/domain"
	"strings"
)

type customerRepository struct {
	Conn *sql.DB
}

func NewCustomerRepository(Conn *sql.DB) domain.CustomerRepository {
	return &customerRepository{Conn}
}

func (db *customerRepository) Save(ctx context.Context, c []*domain.Customer) (err error) {
	valueArgs := []interface{}{}
	valueStrings := []string{}

	for _, row := range c {

		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

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

	query := `INSERT INTO Customers (CustomerID, CompanyName, FisrtName, 
		LastName, BillingAddress, City, StateOrProvince, ZIPCode, Email, 
		CompanyWebsite, PhoneNumber, FaxNumber, ShipAddress, ShipCity, 
		ShipStateOrProvince, ShipZIPCode, ShipPhoneNumber) VALUES %s`

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

	c[len(c)-1].CustomerID = lastID

	return nil
}
