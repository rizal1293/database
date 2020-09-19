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

type employeeReadHandler struct {
	employeeUcase domain.EmployeeUsecase
}

func NewEmployeeReadHandler(u domain.EmployeeUsecase) {
	handler := &employeeReadHandler{
		employeeUcase: u,
	}
	handler.read()
}

func (h *employeeReadHandler) read() {
	csvFile, err := os.Open("./csv_file/Employees.csv")
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

	var employees []*domain.Employee

	for i, item := range records {
		if i > 0 {
			id, err := strconv.Atoi(item[0])
			if err == nil {
				employees = append(employees, &domain.Employee{
					EmployeeID: int64(id),
					FirstName:  item[1],
					LastName:   item[2],
					Title:      item[3],
					WorkPhone:  item[4],
				})
			}
		}
	}

	err = h.employeeUcase.Save(context.TODO(), employees)

	if err != nil {
		log.Print(err.Error())
	} else {
		fmt.Print("Data Employee has been Saved!")
	}
}
