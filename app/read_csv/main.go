package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	delivery "rest_api/delivery/read_csv"

	repCustomer "rest_api/repository/customer"
	repEmployee "rest_api/repository/employee"
	repOrder "rest_api/repository/order"
	repOrdDetail "rest_api/repository/order_detail"
	repProduct "rest_api/repository/product"
	repShipMethod "rest_api/repository/shipping_method"

	ucaseCustomer "rest_api/usecase/customer"
	ucaseEmployee "rest_api/usecase/employee"
	ucaseOrder "rest_api/usecase/order"
	ucaseOrdDetail "rest_api/usecase/order_detail"
	ucaseProduct "rest_api/usecase/product"
	ucaseShipMethod "rest_api/usecase/shipping_method"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "sales"
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	timeout := time.Duration(30) * time.Second

	customerRepo := repCustomer.NewCustomerRepository(dbConn)
	employeeRepo := repEmployee.NewEmployeeRepository(dbConn)
	shipMethodRepo := repShipMethod.NewShippingMethodRepository(dbConn)
	productRepo := repProduct.NewProductRepository(dbConn)
	orderRepo := repOrder.NewOrderRepository(dbConn)
	ordDetailRepo := repOrdDetail.NewOrderDetailRepository(dbConn)

	customerUcase := ucaseCustomer.NewCustomerUsecase(customerRepo, timeout)
	employeeUcase := ucaseEmployee.NewEmployeeUsecase(employeeRepo, timeout)
	shipMethodUcase := ucaseShipMethod.NewShippingMethodRepository(shipMethodRepo, timeout)
	productUcase := ucaseProduct.NewProductUsecase(productRepo, timeout)
	orderUcase := ucaseOrder.NewOrderUsecase(orderRepo, timeout)
	ordDetailUcase := ucaseOrdDetail.NewOrderDetailUsecase(ordDetailRepo, timeout)

	delivery.NewCustomerReadHandler(customerUcase)     // BATCH 1
	delivery.NewEmployeeReadHandler(employeeUcase)     // BATCH 1
	delivery.NewShippingMethodHandler(shipMethodUcase) // BATCH 1
	delivery.NewProductReadHandler(productUcase)       // BATCH 1

	delivery.NewOrderReadHandler(orderUcase) // BATCH 2

	delivery.NewOrderDetailReadHandler(ordDetailUcase) // BATCH 3
}
