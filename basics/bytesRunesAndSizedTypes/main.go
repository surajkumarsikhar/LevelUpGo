package main

import "fmt"

func toUpper(r rune) rune {
	// If r is a lowercase letter ('a' to 'z'), subtract 32 to make it uppercase
	// Otherwise return it unchanged
	if r >= 'a' && r <= 'z' {
		r -= 32
	}
	return r
}

func main() {
	fmt.Println(string(toUpper('g')))
	fmt.Println(string(toUpper('A')))
	fmt.Println(string(toUpper('3')))
}
