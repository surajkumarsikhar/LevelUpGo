package main

import (
	"fmt"
	"math"
)

func DescribeNumber(x float64) string {
	if x == 0 {
		return "zero"
	}

	return fmt.Sprintf("%s, %s", sign(x), magnitude(x))
}

func sign(x float64) string {
	if x > 0 {
		return "positive"
	} else if x < 0 {
		return "negative"
	}
	return "zero"
}

func magnitude(x float64) string {
	if math.Abs(x) <= 10 {
		return "small"
	} else if math.Abs(x) <= 100 {
		return "medium"
	}
	return "large"
}

func main() {
	fmt.Println(DescribeNumber(42))
	fmt.Println(DescribeNumber(-150))
	fmt.Println(DescribeNumber(0))
	fmt.Println(DescribeNumber(5))
	fmt.Println(DescribeNumber(-75))
}
