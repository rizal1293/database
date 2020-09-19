package delivery

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"rest_api/domain"
	"strconv"
)

type orderDetailReadHandler struct {
	orderDetailUcase domain.OrderDetailUsecase
}

func NewOrderDetailReadHandler(u domain.OrderDetailUsecase) {
	handler := &orderDetailReadHandler{
		orderDetailUcase: u,
	}
	handler.read()
}

func (h *orderDetailReadHandler) read() {
	csvFile, err := os.Open("./csv_file/OrderDetails.csv")
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

	var orderDetails []*domain.OrderDetail

	for i, item := range records {
		if i > 0 {
			id, err1 := strconv.Atoi(item[0])
			idOrder, err2 := strconv.Atoi(item[1])
			idProduct, err3 := strconv.Atoi(item[2])

			if err1 == nil && err2 == nil && err3 == nil {

				qty, _ := strconv.Atoi(item[3])
				up, _ := strconv.Atoi(item[4])
				dsc, _ := strconv.Atoi(item[5])

				orderDetails = append(orderDetails, &domain.OrderDetail{
					OrderDetailID: int64(id),
					Order: domain.Order{
						OrderID: int64(idOrder),
					},
					Product: domain.Product{
						ProductID: int64(idProduct),
					},
					Quantity:  qty,
					UnitPrice: up,
					Discount:  dsc,
				})
			}
		}
	}

	err = h.orderDetailUcase.Save(context.TODO(), orderDetails)

	if err != nil {
		log.Println(err)
	} else {
		print("\n")
		fmt.Print("Data Order Detail has been Saved!")
		print("\n")
	}
}
