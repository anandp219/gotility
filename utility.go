package gotility

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	// NUMBER_REGEX is the regex for real number
	NUMBER_REGEX = "^[+-]?([0-9]+(\\.[0-9]*)?|\\.[0-9]+)([eE][+-]?[0-9]+)?$"
	// EMAIL_REGEX is the regex for email
	EMAIL_REGEX = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	// URL_REGEX is the regex for url
	URL_REGEX = "^((((H|h)(T|t)|(F|f))(T|t)(P|p)((S|s)?))\\://)?(www.|[a-zA-Z0-9].)[a-zA-Z0-9\\-\\.]+\\.[a-zA-Z]{2,6}(\\:[0-9]{1,5})*(/($|[a-zA-Z0-9\\.\\,\\;\\?\\'\\\\+&amp;%\\$#\\=~_\\-]+))*$"
)

// IsNumber check whether the string is a number or not
func IsNumber(str string) bool {
	var validNumber = regexp.MustCompile(NUMBER_REGEX)
	return validNumber.MatchString(str)
}

// IsEmail check whether the string is an email or not
func IsEmail(str string) bool {
	var validEmail = regexp.MustCompile(EMAIL_REGEX)
	return validEmail.MatchString(str)
}

// IsUrl check whether the string is an url or not
func IsUrl(str string) bool {
	var validEmail = regexp.MustCompile(URL_REGEX)
	return validEmail.MatchString(str)
}

// IsTitleCase check whether the string is `title Case` or not
func IsTitleCase(str string) bool {
	titleCase := strings.Title(str)
	result := strings.Compare(str, titleCase)
	return result == 0
}

// FlattenFloat64 transform 2D matrix of floating point numbers into 1D array
func FlattenFloat64(matrix [][]float64) []float64 {
	row := []float64{}
	for _, value := range matrix {
		row = append(row, value...)
	}
	return row
}

// FlattenInt transform 2D matrix of Integers into 1D array
func FlattenInt(matrix [][]int) []int {
	row := []int{}
	for _, value := range matrix {
		row = append(row, value...)
	}
	return row
}

// FlattenString transform 2D matrix of string into 1D array
func FlattenString(matrix [][]string) []string {
	row := []string{}
	for _, value := range matrix {
		row = append(row, value...)
	}
	return row
}

// ToMatrixInt transform 1D array Integers to 2D Matrix
func ToMatrixInt(row []int, numRows int) ([][]int, error) {
	length := len(row)
	if numRows <= 0 || length%numRows != 0 {
		return nil, fmt.Errorf("division of row to matrix not possible. Invalid numRows : %d ", numRows)
	}
	numCols := length / numRows
	matrix := make([][]int, numRows)
	for i := 0; i < numRows; i += 1 {
		index := i * numCols
		matrix[i] = append(matrix[i], row[index:index+numCols]...)
	}
	return matrix, nil
}

// ToMatrixString transform 1D array string to 2D Matrix
func ToMatrixString(row []string, numRows int) ([][]string, error) {
	length := len(row)
	if numRows <= 0 || length%numRows != 0 {
		return nil, fmt.Errorf("division of row to matrix not possible. Invalid numRows : %d ", numRows)
	}
	numCols := length / numRows
	matrix := make([][]string, numRows)
	for i := 0; i < numRows; i += 1 {
		index := i * numCols
		matrix[i] = append(matrix[i], row[index:index+numCols]...)
	}
	return matrix, nil
}

// ToMatrixFloat64 transform 1D array Floating point numbers to 2D Matrix
func ToMatrixFloat64(row []float64, numRows int) ([][]float64, error) {
	length := len(row)
	if numRows <= 0 || length%numRows != 0 {
		return nil, fmt.Errorf("division of row to matrix not possible. Invalid numRows : %d ", numRows)
	}
	numCols := length / numRows
	matrix := make([][]float64, numRows)
	for i := 0; i < numRows; i += 1 {
		index := i * numCols
		matrix[i] = append(matrix[i], row[index:index+numCols]...)
	}
	return matrix, nil
}

// SumInt returns the sum of an integer array
func SumInt(row []int) int {
	sum := 0
	for _, value := range row {
		sum += value
	}
	return sum
}

// SumFloat64 returns the sum of an floating point array array
func SumFloat64(row []float64) float64 {
	sum := 0.0
	for _, value := range row {
		sum += value
	}
	return sum
}

// MapInt takes an array of integer values and a function to be applied on that value.
// It returns an array with the mapped values
func MapInt(row []int, fn func(int) int) []int {
	mappedValues := []int{}
	for _, value := range row {
		mappedValues = append(mappedValues, fn(value))
	}
	return mappedValues
}

// MapFloat64 takes an array of float64 values and a function to be applied on that value.
// It returns an array with the mapped values
func MapFloat64(row []float64, fn func(float64) float64) []float64 {
	mappedValues := []float64{}
	for _, value := range row {
		mappedValues = append(mappedValues, fn(value))
	}
	return mappedValues
}

// MapString takes an array of string and a function to be applied on that value.
// It returns an array with the mapped values
func MapString(row []string, fn func(string) string) []string {
	mappedValues := []string{}
	for _, value := range row {
		mappedValues = append(mappedValues, fn(value))
	}
	return mappedValues
}

func sumInt64(value reflect.Value) int64 {
	sum := int64(0)
	for i := 0; i < value.Len(); i += 1 {
		sum += int64(value.Index(i).Int())
	}
	return sum
}

func sumFloat64(value reflect.Value) float64 {
	sum := float64(0)
	for i := 0; i < value.Len(); i += 1 {
		sum += float64(value.Index(i).Float())
	}
	return sum
}

// Sum returns the sum of the slice
func Sum(row interface{}) (interface{}, error) {

	slice := reflect.ValueOf(row)
	if slice.Type().Kind() != reflect.Slice {
		return 0, fmt.Errorf("Expected slice, got: " + slice.Type().Kind().String())
	}

	if slice.Len() == 0 {
		return 0, nil
	}

	switch slice.Index(0).Type().Kind() {
	case reflect.Int:
		return sumInt64(slice), nil
	case reflect.Int8:
		return sumInt64(slice), nil
	case reflect.Int32:
		return sumInt64(slice), nil
	case reflect.Int64:
		return sumInt64(slice), nil
	case reflect.Float32:
		return sumFloat64(slice), nil
	case reflect.Float64:
		return sumFloat64(slice), nil
	default:
		return 0, fmt.Errorf("cannot sum the given slice")
	}

}

// FindIndex returns the first index of the element in the given slice and if not found returns -1
func FindIndex(row interface{}, element interface{}) (int, error) {
	slice := reflect.ValueOf(row)
	if slice.Type().Kind() != reflect.Slice {
		return -1, fmt.Errorf("Expected slice, got: " + slice.Type().Kind().String())
	}

	if slice.Len() == 0 {
		return -1, nil
	}

	sliceElementType := slice.Index(0).Type().Kind()
	elementType := reflect.ValueOf(element).Type().Kind()
	if sliceElementType != elementType {
		return -1, fmt.Errorf("Expected element to be " + sliceElementType.String() + ", got: " + elementType.String())
	}

	for i := 0; i < slice.Len(); i += 1 {
		if reflect.DeepEqual(slice.Index(i).Interface(), element) {
			return i, nil
		}
	}

	return -1, nil
}
