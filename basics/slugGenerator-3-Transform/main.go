package main

import (
	"fmt"
	"strings"
)

// joinWithHyphens replaces spaces with hyphens, collapsing multiples and avoiding edge hyphens
func joinWithHyphens(s string) string {
	// 1. Replace every literal "-" with a space so hyphens and spaces
	//    become the same kind of separator.
	r := strings.ReplaceAll(s, "-", " ")
	// 2. Use strings.Fields to split on runs of whitespace
	//    (it drops empty pieces, so runs collapse and edges trim).
	fields := strings.Fields(r)
	// 3. Join the fields back together with "-" and return.
	result := strings.Join(fields, "-")
	return result
}

func main() {
	fmt.Printf("%q\n", joinWithHyphens("hello world"))
	fmt.Printf("%q\n", joinWithHyphens("  too   many   spaces  "))
	fmt.Printf("%q\n", joinWithHyphens("already-good"))
	fmt.Printf("%q\n", joinWithHyphens(" edge "))
}
