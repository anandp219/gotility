package gotility

import (
	"reflect"
	"testing"
)

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
