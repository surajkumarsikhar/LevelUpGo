package main

import "fmt"

// isLetter returns true if c is a letter (a-z or A-Z)
func isLetter(c rune) bool {
	// Your code here
	if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
		return true
	}
	return false
}

// isDigit returns true if c is a digit (0-9)
func isDigit(c rune) bool {
	// Your code here
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func main() {
	fmt.Println(isLetter('G')) // true
	fmt.Println(isLetter('5')) // false
	fmt.Println(isDigit('9'))  // true
	fmt.Println(isDigit('!'))  // false
}
