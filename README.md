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

## Suggestions

* Suggestions and improvements are welcome.

