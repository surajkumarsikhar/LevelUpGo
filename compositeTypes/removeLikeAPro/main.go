package main

import (
	"errors"
	"fmt"
	"slices"
)

func RemoveAt(tables []string, index int) ([]string, error) {
	// Your code here
	n := len(tables)
	if index < 0 || index >= n || n == 0 {
		return nil, errors.New("index out of bounds")
	}
	tables = slices.Delete(tables, index, (index + 1))
	return tables, nil
}

func main() {
	tables := []string{"users", "orders", "sessions", "orders", "products"}

	result, err := RemoveAt(tables, 2)
	fmt.Println(result, err)
}
