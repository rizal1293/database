package util

import 	"regexp"

func RemoveLineBreaks(str string) string {
	var re = regexp.MustCompile("[\n|\r|\n\r]")
	s := re.ReplaceAllString(str, "")
	return s
}