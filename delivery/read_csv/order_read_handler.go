package delivery

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"rest_api/domain"
	util "rest_api/util"
	"strconv"
)

type orderReadHandler struct {
	orderUcase domain.OrderUsecase
}

func NewOrderReadHandler(u domain.OrderUsecase) {
	handler := &orderReadHandler{
		orderUcase: u,
	}
	handler.read()
}

func (h *orderReadHandler) read() {
	csvFile, err := os.Open("./csv_file/Orders.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	r.LazyQuotes = true
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	var orders []*domain.Order

	for i, item := range records {
		if i > 0 {
			id, err := strconv.Atoi(item[0])
			if err == nil {

				cusId, _ := strconv.Atoi(item[1])
				empId, _ := strconv.Atoi(item[2])
				shipMethodId, _ := strconv.Atoi(item[6])

				freightCharge, _ := strconv.Atoi(item[7])

				taxes, _ := strconv.Atoi(item[8])
				paymentReceived, _ := strconv.Atoi(item[9])

				orders = append(orders, &domain.Order{
					OrderID: int64(id),
					Customer: domain.Customer{
						CustomerID: int64(cusId),
					},
					Employee: domain.Employee{
						EmployeeID: int64(empId),
					},
					OrderDate:           util.ParserToDateYYYMMDD(item[3]),
					PurchaseOrderNumber: item[4],
					ShipDate:            util.ParserToDateYYYMMDD(item[5]),
					ShippingMethod: domain.ShippingMethod{
						ShippingMethodID: int64(shipMethodId),
					},
					FreightCharge:   freightCharge,
					Taxes:           taxes,
					PaymentReceived: paymentReceived,
					Comment:         item[10],
				})
			}
		}
	}

	err = h.orderUcase.Save(context.TODO(), orders)

	if err != nil {
		log.Print(err)
	} else {
		print("\n")
		fmt.Print("Data Orders has been Saved!")
		print("\n")
	}
}
