package main

import "fmt"

// Write Limit: takes a *int and clamps it down to max only if it exceeds max.
func Limit(qty *int, max int) {
	*qty = max
}

// Write LimitAll: calls Limit on every quantity in the slice, in place.
func LimitAll(quantities []int, max int) {
	for i, q := range quantities {
		if q > max {
			Limit(&quantities[i], max)
		}
	}
}

func main() {
	quantities := []int{3, 12, 7, 20}
	LimitAll(quantities, 10)
	fmt.Println(quantities) // [3 10 7 10]
}
