package main

import "fmt"

// TODO: Define Ordered, Min, and Max
type Ordered interface {
	~int | ~int64 | ~float64 | ~string
}

func Min[T Ordered](value, min T) T {
	if value < min {
		return value
	}
	return min
}

func Max[T Ordered](value, max T) T {
	if value > max {
		return value
	}
	return max
}

func main() {
	fmt.Printf("Min(5, 3): %d\n", Min(5, 3))
	fmt.Printf("Max(5, 3): %d\n", Max(5, 3))
	fmt.Printf("Min(hello, world): %s\n", Min("hello", "world"))
	fmt.Printf("Max(1.5, 2.5): %.1f\n", Max(1.5, 2.5))
}
