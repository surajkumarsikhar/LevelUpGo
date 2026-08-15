package main

import "fmt"

func main() {
	maxRetries := 3

	// Store the address of maxRetries in p
	p := &maxRetries

	fmt.Println(p == &maxRetries)
}
