package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	createError()
	customError()
}

func createError() {
	result, err := Sqrt(-10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

func customError() {
	result, err := CustomSqrt(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

type dualError struct {
	Num     float64
	problem string
}

func (e dualError) Error() string {  // 实现 error 接口
	return fmt.Sprintf("Wrong!!!,because \"%f\" is a negative number", e.Num)
}

func CustomSqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f}
	}
	return math.Sqrt(f), nil
}