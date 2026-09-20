package main

import "fmt"

func IndexOf[T comparable](slice []T, target T) int {
	// TODO: implement
	for i, item := range slice {
		if item == target {
			return i
		}
	}
	return -1
}

func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := make([]T, 0, len(slice))

	for _, value := range slice {
		if seen[value] {
			continue
		}

		seen[value] = true
		result = append(result, value)
	}

	return result
}

func main() {
	// Basic demonstration
	names := []string{"alice", "bob", "charlie"}
	fmt.Println("IndexOf bob:", IndexOf(names, "bob"))
	fmt.Println("Unique:", Unique([]int{1, 2, 2, 3}))
}
