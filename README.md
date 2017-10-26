# gotility

[![Build Status](https://travis-ci.org/anandp219/gotility.svg?branch=master)](https://travis-ci.org/anandp219/gotility)
[![Go Report Card](https://goreportcard.com/badge/github.com/anandp219/gotility)](https://goreportcard.com/report/github.com/anandp219/gotility)
[![GoDoc](https://godoc.org/github.com/anandp219/gotility?status.svg)](https://godoc.org/github.com/anandp219/gotility)
[![codecov](https://codecov.io/gh/anandp219/gotility/branch/master/graph/badge.svg)](https://codecov.io/gh/anandp219/gotility)

A utility library for go

# Installation

```
go get github.com/anandp219/gotility
```

# Usage 

```
import "github.com/anandp219/gotility"
```
## Functions
List of Functions added till now with examples

### IsNumber

```
IsNumber("12134.0009")  // true
IsNumber("a12134.0009") // false
```

### IsEmail

```
IsEmail("xyz@gmail.com")  // true
IsEmail("mail")                 // false
```

### IsUrl

```
IsUrl("google.com")  // true
IsUrl("123.com1")    // false
```

### IsTitleCase

```
IsTitleCase("Go Is VEry Popular")  // true
IsTitleCase("gO IS NOT POPULAR")   // false
```

### FlattenFloat64

```
FlattenFloat64([][]float64{{1.1, 2.1}, {3.09, 4.5}}) // => [1.1 2.1 3.09 4.5]
```

### FlattenInt

```
FlattenInt([][]int{{1, 2}, {3, 4}}) // => [1 2 3 4]
```

### FlattenString

```
FlattenString([][]string{{"go", "is"}, {"very", "popular"}}) // =>  [go is very popular]
```

### ToMatrixInt

```
ToMatrixInt([]int{1, 2, 3, 4}, 2 /* Number of rows */) // =>  [[1, 2], [3, 4]]
```

### ToMatrixFloat64

```
ToMatrixFloat64([]float64{1.1, 2.1, 3.1, 4.1}, 2 /* Number of rows */) // =>  [[1.1, 2.1], [3.1, 4.1]]
```

### ToMatrixString

```
ToMatrixString([]string{"A", "B", "C", "D"}, 2 /* Number of rows */) // =>  [[A B] [C D]]
```

### SumInt

```
SumInt([]int{1, 2, 3}) // => 6
```

### SumFloat64

```
SumFloat64([]float64{1.1, 2.2, 3.3}) // => 6.6
```

### MapInt

* Function takes 2 arguements. An array of int values and function to be applied on each of the values
* In the following examples we will use square function which is defined as
```
func square(arg int) int {
    return arg*arg
}
```

```
MapInt([]int{1, 2, 3}, square) // => [1, 4, 9]
```

### MapFloat64

* In the following examples we will use square function which is defined as
```
func square(arg float64) float64 {
    return arg*arg
}
```

```
MapFloat64([]float64{1.1, 2.2, 3.3}, square) // => [1.2100000000000002 4.840000000000001 10.889999999999999]
```

### MapString

* In the following examples we will use addSemicolon function which is defined as
```
func addSemicolon(arg string) string {
    return arg+";"
}
```

```
MapString([]string{"a", "b",  "c"}, addSemicolon) // => [a; b; c;]
```

### Sum

* Find the sum of slice for int, int8, int16, int32, int64, float32 and float64 data types
* returns tuple `sum(int64, float64), Error`
```
Sum([]int{1, 2, 3}) // => 6, nil
Sum([]float32{1, 2, 3.2}) // => 6.200000047683716, nil
Sum([]bool{true}) // => 0, "cannot sum the given slice"
```

### FindIndex

* Find the first occurrence of the element in the `slice`.
* returns tuple `int, Error`
```
FindIndex([]int{1, 2, 3}, 1) // => 0, nil
FindIndex([]float64{1, 2, 3}, -1.0) // => -1, nil
FindIndex([]int{1, 2, 3}, 1.0) // => -1, Error("Mismatch type of elements")
```

### FindLastIndex

* Find the last occurrence of the element in the `slice`.
* returns tuple `int, Error`
```
FindLastIndex([]int{1, 2, 3, 1}, 1) // => 3, nil
FindLastIndex([]float64{1, 2, 3}, -1.0) // => -1, nil
FindLastIndex([]int{1, 2, 3}, 1.0) // => -1, Error("Mismatch type of elements")
```

### GetKeys

* Returns the keys as an unordered slice of the given map.
* returns tuple `slice, Error`
```
GetKeys(map[string]string{"key1": "value1", "key2": "value2}) // => ["key1", "key2"], nil
GetKeys(1) // => nil, Error("Expected map")
```

### GetValues

* Returns the values as an unordered slice of the given map.
* returns tuple `slice, Error`
```
GetValues(map[string]string{"key1": "value1", "key2": "value2}) // => ["value1", "value2"], nil
GetValues(1) // => nil, Error("Expected map")
```


## Suggestions

* Suggestions and improvements are welcome.

