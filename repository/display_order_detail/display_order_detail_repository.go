package repository

import (
	"context"
	"database/sql"
	"rest_api/domain"

	"github.com/sirupsen/logrus"
)

type customerOrderRepository struct {
	Conn *sql.DB
}

func NewCustomerOrder(Conn *sql.DB) domain.CustomerOrderRepository {
	return &customerOrderRepository{Conn}
}

func (db *customerOrderRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.CustomerOrderDetail, err error) {
	rows, err := db.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(err)
		}
	}()

	result = make([]domain.CustomerOrderDetail, 0)
	for rows.Next() {
		toBeAdded := domain.CustomerOrderDetail{}
		err = rows.Scan(
			&toBeAdded.OrderID,
			&toBeAdded.EmployeeName,
			&toBeAdded.ShippingMethod,
		)

		if err != nil {
			logrus.Error(err)
		}
		result = append(result, toBeAdded)
	}

	return result, nil
}

func (db *customerOrderRepository) GetOrdersByCustomer(ctx context.Context, name string) (domain.CustomerOrder, error) {
	query := `SELECT o.OrderID, CONCAT(CONCAT(e.FisrtName, ' '), e.LastName) EmployeeName, sm.ShippingMethod from Orders o
			inner join Customers c on c.CustomerID = o.CustomerID 	
			inner join Employees e on e.EmployeeID = o.EmployeeID
			inner join Shipping_Methods sm on sm.ShippingMethodID = o.ShippingMethodID 
			where o.ShipDate is not NULL and c.FisrtName = ?
			order by e.FisrtName`

	res, err := db.fetch(ctx, query, name)
	if err != nil {
		return domain.CustomerOrder{}, err
	}

	if len(res) > 0 {
		return domain.CustomerOrder{
			CustomerName: name,
			Orders:       res,
		}, nil
	}
	return domain.CustomerOrder{}, nil
}
