package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}

func main() {
	var result, error = Sqrt(10.0)
	if error != nil {
		fmt.Println("The error message", error)
	} else {
		fmt.Println("The SQRT :", result)
	}
}
