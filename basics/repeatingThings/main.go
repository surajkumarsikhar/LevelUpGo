package main

import "fmt"

func retry(successOn int, maxRetries int) string {
	result := ""

	// Your code here
	for i := 0 ; i < maxRetries ; i++ {
		if i == successOn {
			return "success"
		}
	}
	result = "failed"
	return result
}

func main() {
	fmt.Println(retry(1, 5))
	fmt.Println(retry(6, 5))
	fmt.Println(retry(0, 3))
}