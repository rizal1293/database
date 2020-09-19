package repository

import (
	"context"
	"database/sql"
	"fmt"
	"rest_api/domain"
	"strings"
)

type employeeRepository struct {
	Conn *sql.DB
}

func NewEmployeeRepository(Conn *sql.DB) domain.EmployeeRepository {
	return &employeeRepository{Conn}
}

func (db *employeeRepository) Save(ctx context.Context, e []*domain.Employee) (err error) {
	valueArgs := []interface{}{}
	valueStrings := []string{}

	for _, row := range e {

		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")

		valueArgs = append(valueArgs, row.EmployeeID)
		valueArgs = append(valueArgs, row.FirstName)
		valueArgs = append(valueArgs, row.LastName)
		valueArgs = append(valueArgs, row.Title)
		valueArgs = append(valueArgs, row.WorkPhone)
	}

	query := `INSERT INTO Employees (EmployeeID, FisrtName, LastName, Title, WorkPhone) VALUES %s`
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

	e[len(e)-1].EmployeeID = lastID

	return nil
}
