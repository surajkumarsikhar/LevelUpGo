package main

import "fmt"

func addLatency(total *int, ms int) {
	// Add ms to the value total points to
	*total += ms
}

func main() {
	totalMillis := 0
	addLatency(&totalMillis, 12)
	addLatency(&totalMillis, 30)
	addLatency(&totalMillis, 8)
	fmt.Println(totalMillis)
}
