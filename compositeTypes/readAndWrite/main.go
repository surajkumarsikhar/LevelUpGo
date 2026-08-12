package main

import "fmt"

func RotateStatus(statuses []string, index int, newStatus string) {
	// Your code here
	statuses[index] = newStatus
}

func main() {
	statuses := []string{"healthy", "healthy", "degraded"}
	fmt.Println(statuses)
	RotateStatus(statuses, 2, "healthy")
	fmt.Println(statuses)
}
