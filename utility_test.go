package gotility

import (
	"fmt"
	"reflect"
	"testing"
)

func squareInt(arg int) int {
	return arg * arg
}

func squareFloat64(arg float64) float64 {
	return arg * arg
}

func squareString(arg string) string {
	return arg + arg
}

func TestIsEmail(t *testing.T) {
	isEmail := IsEmail("xyz@gmail.com")
	if isEmail != true {
		t.Error("Expected true, got ", isEmail)
	}
	isEmail = IsEmail("mail")
	if isEmail != false {
		t.Error("Expected false, got ", isEmail)
	}
}

func TestIsUrl(t *testing.T) {
	isUrl := IsUrl("google.com")
	if isUrl != true {
		t.Error("Expected true, got ", isUrl)
	}
	isUrl = IsUrl("123.com1")
	if isUrl != false {
		t.Error("Expected false, got ", isUrl)
	}
}

func TestIsNumber(t *testing.T) {
	isNumber := IsNumber("12134.0009")
	if isNumber != true {
		t.Error("Expected true, got ", isNumber)
	}
	isNumber = IsNumber("a12134.0009")
	if isNumber != false {
		t.Error("Expected false, got ", isNumber)
	}
}

func TestIsTitleCase(t *testing.T) {
	isTitleCase := IsTitleCase("Go Is VEry Popular")
	if isTitleCase != true {
		t.Error("Expected true, got ", isTitleCase)
	}
	isTitleCase = IsTitleCase("gO IS NOT POPULAR")
	if isTitleCase != false {
		t.Error("Expected false, got ", isTitleCase)
	}
}

func TestFlattenFloat64(t *testing.T) {
	row := FlattenFloat64([][]float64{{1.1, 2.2}, {3.3, 4.4}})
	if !reflect.DeepEqual(row, []float64{1.1, 2.2, 3.3, 4.4}) {
		t.Error("Expected [1.1 2.2 3.3 4.4], got ", row)
	}
}

func TestFlattenInt(t *testing.T) {
	row := FlattenInt([][]int{{1, 2}, {3, 4}})
	if !reflect.DeepEqual(row, []int{1, 2, 3, 4}) {
		t.Error("Expected [1 2 3 4], got ", row)
	}
}

func TestFlattenString(t *testing.T) {
	row := FlattenString([][]string{{"A", "B"}, {"C", "D"}})
	if !reflect.DeepEqual(row, []string{"A", "B", "C", "D"}) {
		t.Error("Expected [A B C D], got ", row)
	}
}

func TestSumInt(t *testing.T) {
	sum := SumInt([]int{1, 2, 3})
	if sum != 6 {
		t.Error("Expected 6, got ", sum)
	}
}

func TestSumFloat64(t *testing.T) {
	sum := SumFloat64([]float64{1.1, 2.2, 3.3})
	if sum != 6.6 {
		t.Error("Expected 6.6 got ", sum)
	}
}

func TestMapInt(t *testing.T) {
	square := MapInt([]int{1, 2, 3}, squareInt)
	if !reflect.DeepEqual(square, []int{1, 4, 9}) {
		t.Error("Expected [1 4 9], got ", square)
	}
}

func TestMapFloat64(t *testing.T) {
	square := MapFloat64([]float64{1.1, 2.2, 3.3}, squareFloat64)
	if !reflect.DeepEqual(square, []float64{1.2100000000000002, 4.840000000000001, 10.889999999999999}) {
		t.Error("Expected 1.2100000000000002 4.840000000000001 10.889999999999999], got ", square)
	}
}

func TestMapString(t *testing.T) {
	square := MapString([]string{"A", "B", "C"}, squareString)
	if !reflect.DeepEqual(square, []string{"AA", "BB", "CC"}) {
		t.Error("Expected [AA BB CC], got ", square)
	}
}

func TestToMatrixInt(t *testing.T) {
	_, err := ToMatrixInt([]int{1, 2, 3, 4}, 0)
	if err == nil {
		t.Error("Expected ", fmt.Errorf("division of row to matrix not possible. Invalid numRows : %d ", 0), ", got ", err)
	}
	_, err = ToMatrixInt([]int{1, 2, 3, 4}, 3)
	if err == nil {
		t.Error("Expected ", fmt.Errorf("division of row to matrix not possible. Invalid numRows : %d ", 3), ", got ", err)
	}
	matrix, _ := ToMatrixInt([]int{1, 2, 3, 4}, 2)
	if !reflect.DeepEqual(matrix, [][]int{{1, 2}, {3, 4}}) {
		t.Error("Expected [[1, 2] [3, 4]], got ", matrix)
	}
}

func TestToMatrixFloat64(t *testing.T) {
	_, err := ToMatrixFloat64([]float64{1.1, 2.2, 3.3, 4.4}, 0)
	if err == nil {
		t.Error("Expected ", fmt.Errorf("division of row to matrix not possible. Invalid numRows : %d ", 0), ", got ", err)
	}
	_, err = ToMatrixFloat64([]float64{1.1, 2.2, 3.3, 4.4}, 3)
	if err == nil {
		t.Error("Expected ", fmt.Errorf("division of row to matrix not possible. Invalid numRows : %d ", 3), ", got ", err)
	}
	matrix, _ := ToMatrixFloat64([]float64{1.1, 2.2, 3.3, 4.4}, 2)
	if !reflect.DeepEqual(matrix, [][]float64{{1.1, 2.2}, {3.3, 4.4}}) {
		t.Error("Expected [[1.1, 2.2] [3.3, 4.4]], got ", matrix)
	}
}

func TestToMatrixString(t *testing.T) {
	_, err := ToMatrixString([]string{"A", "B", "C", "D"}, 0)
	if err == nil {
		t.Error("Expected ", fmt.Errorf("division of row to matrix not possible. Invalid numRows : %d ", 0), ", got ", err)
	}
	_, err = ToMatrixString([]string{"A", "B", "C", "D"}, 3)
	if err == nil {
		t.Error("Expected ", fmt.Errorf("division of row to matrix not possible. Invalid numRows : %d ", 3), ", got ", err)
	}
	matrix, _ := ToMatrixString([]string{"A", "B", "C", "D"}, 2)
	if !reflect.DeepEqual(matrix, [][]string{{"A", "B"}, {"C", "D"}}) {
		t.Error("Expected [[A, B] [C, D]], got ", matrix)
	}
}

func TestSum(t *testing.T) {
	sum, _ := Sum([]int{1, 2, 3})
	if sum != int64(6) {
		t.Error("Expected 6 got ", sum)
	}
	sum, _ = Sum([]float64{1, 2, 3})
	if sum != float64(6) {
		t.Error("Expected 6 got ", sum)
	}
	sum, _ = Sum([]float64{0})
	if sum != float64(0) {
		t.Error("Expected 0 got ", sum)
	}
	_, err := Sum([]bool{true})
	if err == nil {
		t.Error("Expected ", fmt.Errorf("cannot sum the given slice"), " got <nil>")
	}
	_, err = Sum(2)
	if err == nil {
		t.Error("Expected ", "Expected slice, got: int ", " got <nil>")
	}

}
