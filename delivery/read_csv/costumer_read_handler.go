package delivery

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"rest_api/domain"
	"rest_api/util"
	"strconv"
)

type customerReadHandler struct {
	customerUcase domain.CustomerUsecase
}

func NewCustomerReadHandler(u domain.CustomerUsecase) {
	handler := &customerReadHandler{
		customerUcase: u,
	}
	handler.read()
}

func (h *customerReadHandler) read() {
	csvFile, err := os.Open("./csv_file/Customers.csv")
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

	var costumers []*domain.Customer

	for i, item := range records {
		if i > 0 {
			id, err := strconv.Atoi(item[0])
			if err == nil {
				costumers = append(costumers, &domain.Customer{
					CustomerID:          int64(id),
					CompanyName:         util.RemoveLineBreaks(item[1]),
					FisrtName:           util.RemoveLineBreaks(item[2]),
					LastName:            util.RemoveLineBreaks(item[3]),
					BillingAddress:      util.RemoveLineBreaks(item[4]),
					City:                util.RemoveLineBreaks(item[5]),
					StateOrProvince:     util.RemoveLineBreaks(item[6]),
					ZIPCode:             util.RemoveLineBreaks(item[7]),
					Email:               util.RemoveLineBreaks(item[8]),
					CompanyWebsite:      util.RemoveLineBreaks(item[9]),
					PhoneNumber:         util.RemoveLineBreaks(item[10]),
					FaxNumber:           util.RemoveLineBreaks(item[11]),
					ShipAddress:         util.RemoveLineBreaks(item[12]),
					ShipCity:            util.RemoveLineBreaks(item[13]),
					ShipStateOrProvince: util.RemoveLineBreaks(item[14]),
					ShipZIPCode:         util.RemoveLineBreaks(item[15]),
					ShipPhoneNumber:     util.RemoveLineBreaks(item[16]),
				})
			}

		}
	}

	err = h.customerUcase.Save(context.TODO(), costumers)
	if err != nil {
		log.Print(err.Error())
	} else {
		print("\n")
		fmt.Print("Data Customer has been Saved!")
		print("\n")
	}
}
