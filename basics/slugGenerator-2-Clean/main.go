package main

import "fmt"

// cleanText returns s with only letters, digits, and spaces kept
func cleanText(s string) string {
	result := ""
	// Your code here
	for _, ch := range s {
		flag := isLetter(ch) || isDigit(ch) || ch == ' '
		if flag {
			result += string(ch)
		}
	}
	return result
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func main() {
	fmt.Printf("%q\n", cleanText("Hello, World!"))
	fmt.Printf("%q\n", cleanText("Go 1.21 is great!!!"))
	fmt.Printf("%q\n", cleanText("@#$ test &*"))
	fmt.Printf("%q\n", cleanText(""))
}
