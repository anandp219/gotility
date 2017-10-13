package gotility

import (
	"regexp"
	"strings"
)

const (
	ASCII_VALUE_0 = 48
	ASCII_VALUE_9 = 57
	EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	URL_REGEX = "^((((H|h)(T|t)|(F|f))(T|t)(P|p)((S|s)?))\\://)?(www.|[a-zA-Z0-9].)[a-zA-Z0-9\\-\\.]+\\.[a-zA-Z]{2,6}(\\:[0-9]{1,5})*(/($|[a-zA-Z0-9\\.\\,\\;\\?\\'\\\\+&amp;%\\$#\\=~_\\-]+))*$"
)

func IsNumber (str string) bool {
	if len(str) == 0 {
		return false
	}
	for index := range str {
		if uint8(str[index]) < ASCII_VALUE_0 || uint8(str[index]) > ASCII_VALUE_9 {
			return false
		}
	}
	return true
}

func IsEmail (str string) bool {
	var validEmail = regexp.MustCompile(EMAIL_REGEX)
	return validEmail.MatchString(str)
}

func IsUrl (str string) bool {
	var validEmail = regexp.MustCompile(URL_REGEX)
	return validEmail.MatchString(str)
}

func IsTitleCase (str string) bool {
	titleCase := strings.Title(str)
	result := strings.Compare(str, titleCase)
	return result == 0
}

func FlattenFloat64 (matrix [][]float64) []float64 {
	row := []float64{}
	for _,value := range matrix {
		row = append(row, value...)
	}
	return row
}

func FlattenInt (matrix [][]int) []int {
	row := []int{}
	for _,value := range matrix {
		row = append(row, value...)
	}
	return row
}

func FlattenString (matrix [][]string) []string {
	row := []string{}
	for _,value := range matrix {
		row = append(row, value...)
	}
	return row
}
