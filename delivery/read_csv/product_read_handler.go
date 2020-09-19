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

type productReadHandler struct {
	productUcase domain.ProductUsecase
}

func NewProductReadHandler(u domain.ProductUsecase) {
	handler := &productReadHandler{
		productUcase: u,
	}
	handler.read()
}

func (h *productReadHandler) read() {
	csvFile, err := os.Open("./csv_file/Products.csv")
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

	var products []*domain.Product

	for i, item := range records {
		if i > 0 {
			id, err := strconv.Atoi(item[0])
			up, _ := strconv.Atoi(item[2])

			if err == nil {
				products = append(products, &domain.Product{
					ProductID:   int64(id),
					ProductName: item[1],
					UntiPrice:   up,
					InStock:     item[3],
				})
			}
		}
	}

	err = h.productUcase.Save(context.TODO(), products)

	if err != nil {
		log.Print(err.Error())
	} else {
		print("\n")
		fmt.Print("Data Product has been Saved!")
		print("\n")
	}
}
