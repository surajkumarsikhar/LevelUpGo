package main

import "fmt"

// TODO: Implement Filter and Map

func Filter[T any](slice []T, predicate func(T) bool) []T {
	res := []T{}
	for _, item := range slice {
		if predicate(item) {
			res = append(res, item)
		}
	}
	return res
}

func Map[T, R any](slice []T, transform func(T) R) []R {
	res := []R{}

	for _, item := range slice {
		res = append(res, transform(item))
	}
	return res
}

func main() {
	evens := Filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n%2 == 0
	})
	fmt.Printf("Evens: %v\n", evens)

	lengths := Map([]string{"Go", "Rust!", "Cat"}, func(s string) int {
		return len(s)
	})
	fmt.Printf("Lengths: %v\n", lengths)

	doubled := Map([]int{1, 2, 3}, func(n int) int {
		return n * 2
	})
	fmt.Printf("Doubled: %v\n", doubled)
}
