package gotility

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	NUMBER_REGEX = "^[+-]?([0-9]+(\\.[0-9]*)?|\\.[0-9]+)([eE][+-]?[0-9]+)?$"
	EMAIL_REGEX  = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	URL_REGEX    = "^((((H|h)(T|t)|(F|f))(T|t)(P|p)((S|s)?))\\://)?(www.|[a-zA-Z0-9].)[a-zA-Z0-9\\-\\.]+\\.[a-zA-Z]{2,6}(\\:[0-9]{1,5})*(/($|[a-zA-Z0-9\\.\\,\\;\\?\\'\\\\+&amp;%\\$#\\=~_\\-]+))*$"
)

func IsNumber(str string) bool {
	var validNumber = regexp.MustCompile(NUMBER_REGEX)
	return validNumber.MatchString(str)
}

func IsEmail(str string) bool {
	var validEmail = regexp.MustCompile(EMAIL_REGEX)
	return validEmail.MatchString(str)
}

func IsUrl(str string) bool {
	var validEmail = regexp.MustCompile(URL_REGEX)
	return validEmail.MatchString(str)
}

func IsTitleCase(str string) bool {
	titleCase := strings.Title(str)
	result := strings.Compare(str, titleCase)
	return result == 0
}

func FlattenFloat64(matrix [][]float64) []float64 {
	row := []float64{}
	for _, value := range matrix {
		row = append(row, value...)
	}
	return row
}

func FlattenInt(matrix [][]int) []int {
	row := []int{}
	for _, value := range matrix {
		row = append(row, value...)
	}
	return row
}

func FlattenString(matrix [][]string) []string {
	row := []string{}
	for _, value := range matrix {
		row = append(row, value...)
	}
	return row
}

func ToMatrixInt(row []int, numRows int) ([][]int, error) {
	length := len(row)
	if numRows <= 0 || length%numRows != 0 {
		return nil, errors.New(fmt.Sprintf("divison of row to matrix not possible. Invalid numRows : %d ", numRows))
	}
	numCols := length / numRows
	matrix := make([][]int, numRows)
	for i := 0; i < numRows; i += 1 {
		index := i * numCols
		matrix[i] = append(matrix[i], row[index:index+numCols]...)
	}
	return matrix, nil
}

func ToMatrixString(row []string, numRows int) ([][]string, error) {
	length := len(row)
	if numRows <= 0 || length%numRows != 0 {
		return nil, errors.New(fmt.Sprintf("divison of row to matrix not possible. Invalid numRows : %d ", numRows))
	}
	numCols := length / numRows
	matrix := make([][]string, numRows)
	for i := 0; i < numRows; i += 1 {
		index := i * numCols
		matrix[i] = append(matrix[i], row[index:index+numCols]...)
	}
	return matrix, nil
}

func ToMatrixFloat64(row []float64, numRows int) ([][]float64, error) {
	length := len(row)
	if numRows <= 0 || length%numRows != 0 {
		return nil, errors.New(fmt.Sprintf("divison of row to matrix not possible. Invalid numRows : %d ", numRows))
	}
	numCols := length / numRows
	matrix := make([][]float64, numRows)
	for i := 0; i < numRows; i += 1 {
		index := i * numCols
		matrix[i] = append(matrix[i], row[index:index+numCols]...)
	}
	return matrix, nil
}
