package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	delivery "rest_api/delivery/rest_api"
	"rest_api/delivery/rest_api/middleware"
	repoOrder "rest_api/repository/display_order_detail"
	ucaseOrder "rest_api/usecase/display_order_detail"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
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

	e := echo.New()

	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	orderRepo := repoOrder.NewCustomerOrder(dbConn)
	orderUcase := ucaseOrder.NewCustomerOrder(orderRepo, timeout)

	delivery.NewCustomerOrderHandler(e, orderUcase)

	log.Fatal(e.Start(":8080"))
}
