package main

import "fmt"

func main() {
	maxConns := 10
	p := &maxConns

	// Use p to set maxConns to 25
	*p = 25

	fmt.Println(maxConns)
}
