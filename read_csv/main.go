package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
)

type M map[string]interface{}

func main() {
	csvFile, err := os.Open("./csv_file/ShippingMethods.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	r.LazyQuotes = true
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i, line := range records {
		if i > 0 {
			for _, col := range line {
				var re = regexp.MustCompile("[\n|\r|\n\r]")
				s := re.ReplaceAllString(col, "")
				fmt.Println(s)
			}
			print("=========\n")
		}
	}

}
