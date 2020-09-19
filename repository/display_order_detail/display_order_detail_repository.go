package repository

import (
	"context"
	"database/sql"
	"rest_api/domain"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

type orderDetails struct {
	Product []domain.DetailOrderProduct
	Total   int
}

type customerOrderRepository struct {
	Conn *sql.DB
}

func NewCustomerOrder(Conn *sql.DB) domain.CustomerOrderRepository {
	return &customerOrderRepository{Conn}
}

func (db *customerOrderRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.CustomerOrderDetail, err error) {

	g, ctx := errgroup.WithContext(ctx)

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
	channelProduct := make(chan orderDetails, 0)

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

		// using gorountine to fetch detail order
		g.Go(func() error {
			res, total := db.fetchDetailOrder(ctx, &toBeAdded.OrderID)
			channelProduct <- orderDetails{
				Product: res,
				Total:   total,
			}
			return nil
		})

		result = append(result, toBeAdded)
	}

	go func() {
		err := g.Wait()
		if err != nil {
			logrus.Error(err)
			return
		}
		close(channelProduct)
	}()

	index := 0
	for data := range channelProduct {
		result[index].Product = data.Product
		result[index].Total = data.Total

		index++
	}

	return result, nil
}

func (db *customerOrderRepository) fetchDetailOrder(ctx context.Context, args ...interface{}) (result []domain.DetailOrderProduct, total int) {
	query := `SELECT 
				p.ProductName,
				p.UnitPrice,
				sum(od.Quantity) Quantity,
				sum(od.Discount) Discount,
				sum(
					(p.UnitPrice * od.Quantity)
				) Subtotal
				FROM 
				Order_Details od 
				inner join Products p ON p.ProductID = od.ProductID 
				inner join Orders o on o.OrderID = od.OrderID
				where o.OrderID = ?
				group by 
				p.ProductID, p.ProductName, p.UnitPrice`

	result = make([]domain.DetailOrderProduct, 0)
	total = 0

	rows, err := db.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return result, 0
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(err)
		}
	}()

	for rows.Next() {
		toBeAdded := domain.DetailOrderProduct{}
		err = rows.Scan(
			&toBeAdded.Product,
			&toBeAdded.UnitPrice,
			&toBeAdded.Quantity,
			&toBeAdded.Discount,
			&toBeAdded.Subtotal,
		)

		if err != nil {
			logrus.Error(err)
		}
		total += toBeAdded.Subtotal
		result = append(result, toBeAdded)
	}

	return result, total
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
