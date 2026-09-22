package main

import (
	"fmt"
	"strconv"
)

// TODO: Implement Zero and Map
func Zero[T any]() T {
	var zero T
	return zero
}

func Map[T, R any](slice []T, fn func(T) R) []R {
	result := []R{}
	for _, item := range slice {
		result = append(result, fn(item))
	}
	return result
}

func main() {
	fmt.Printf("Zero[int]: %d\n", Zero[int]())
	fmt.Printf("Zero[string]: %s(empty)\n", Zero[string]())
	fmt.Printf("Zero[bool]: %v\n", Zero[bool]())

	nums := []int{1, 2, 3}

	doubled := Map(nums, func(n int) int {
		return n * 2
	})
	fmt.Printf("Map to double: %v\n", doubled)

	strings := Map(nums, func(n int) string {
		return strconv.Itoa(n)
	})
	fmt.Printf("Map to string: %v\n", strings)
}
