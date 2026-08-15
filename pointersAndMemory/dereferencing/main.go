package main

import "fmt"

func main() {
	delay := 200
	p := &delay

	// Read the delay through p (not delay directly), double it, and print the result
	fmt.Printf("%d", (*p * 2))
}
