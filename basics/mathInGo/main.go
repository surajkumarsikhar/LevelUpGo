package main

import "fmt"

func main() {
	total := 95.50
	people := 4
	// Create a variable "share" that splits total between people
	// Hint: you can't divide a float by an int directly in Go!
	share := total/float64(people)

	fmt.Printf("Each person pays: $%.2f\n", share)
}