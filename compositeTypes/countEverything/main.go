package main

import "fmt"

func CountVisits(paths []string) map[string]int {
	// Your code here
	m := make(map[string]int)
	for _, path := range paths {
		m[path]++
	}
	return m
}

func main() {
	log := []string{"/api/users", "/healthz", "/api/users", "/api/orders", "/healthz", "/api/users"}
	fmt.Println(CountVisits(log))
}
