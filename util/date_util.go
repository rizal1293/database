package util

import (
	"strings"
)

func ParserToDateYYYMMDD(strDate string) string {
	if len(strDate) == 10 {
		replaced := strings.ReplaceAll(strDate, "/", "-")
		dd := replaced[:2]
		mm := replaced[2:6]
		yyyy := replaced[6:]

		formated := yyyy + mm + dd

		return formated
	}

	return "0000-00-00"
}
