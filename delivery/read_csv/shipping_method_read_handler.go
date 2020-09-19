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

type shippingMethodHandler struct {
	shipMethodUcase domain.ShippingMethodUsecase
}

func NewShippingMethodHandler(u domain.ShippingMethodUsecase) {
	handler := &shippingMethodHandler{
		shipMethodUcase: u,
	}
	handler.read()
}

func (h *shippingMethodHandler) read() {
	csvFile, err := os.Open("./csv_file/ShippingMethods.csv")
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

	var shippingMethods []*domain.ShippingMethod

	for i, item := range records {
		if i > 0 {
			id, err := strconv.Atoi(item[0])
			if err == nil {
				shippingMethods = append(shippingMethods, &domain.ShippingMethod{
					ShippingMethodID: int64(id),
					ShippingMethod:   item[1],
				})
			}
		}
	}

	err = h.shipMethodUcase.Save(context.TODO(), shippingMethods)

	if err != nil {
		log.Print(err.Error())
	} else {
		print("\n")
		fmt.Print("Data Shipping Method has been Saved!")
		print("\n")
	}
}
