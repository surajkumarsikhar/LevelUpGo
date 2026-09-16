package main

import "fmt"

// TODO: Define Number, Sum, and Average
type Number interface {
	int | int64 | float64
}

func Sum[T Number](nums []T) T {
	var acc T
	for _, num := range nums {
		acc += num
	}
	return acc
}

func Average[T Number](nums []T) float64 {
	if len(nums) == 0 {
		return 0.0
	}
	var acc T
	for _, num := range nums {
		acc += num
	}
	n := float64(len(nums))
	sum := float64(acc)
	return sum / n
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum: %d, Average: %.1f\n", Sum(ints), Average(ints))
}
