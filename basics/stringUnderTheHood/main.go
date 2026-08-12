package main

import (
	"fmt"
	"unicode/utf8"
)

func firstChar(s string) string {
	// 1. If s is empty, return "" right away.
	if s == "" {
		return ""
	}
	// 2. Decode the first rune (import "unicode/utf8").
	r, _ := utf8.DecodeRuneInString(s)
	// 3. Convert that rune to a string and return it.
	z := string(r)
	return z
}

func main() {
	fmt.Println(firstChar("Hello"))
	fmt.Println(firstChar("世界"))
	fmt.Println(firstChar("🚀 launch"))
	fmt.Println(firstChar(""))
}
