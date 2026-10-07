package main

import (
	"fmt"
)

func SafeDivide(dividend, divisor float64) (float64, error) {
	if dividend == 0.0 && divisor == 0.0 {
		return 0.0, fmt.Errorf("undefined: 0/0")
	} else if divisor == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	} else {
		return dividend / divisor, nil
	}
}

func main() {
	result, err := SafeDivide(10.0, 2.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", int(result))
	}

	result, err = SafeDivide(10.0, 0.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 0 = %f\n", result)
	}

	result, err = SafeDivide(0.0, 0.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("0 / 0 = %f\n", result)
	}

	result, err = SafeDivide(0.0, 5.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("0 / 5 = %d\n", int(result))
	}
}
