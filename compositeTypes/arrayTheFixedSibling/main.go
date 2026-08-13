package main

import "fmt"

func StatusCategory(code int) string {
	// Your code here
	if code > 599 || code < 100 {
		return "Unknown"
	}
	stauses := [5]string{"Informational", "Success", "Redirection", "Client Error", "Server Error"}
	index := (code / 100) - 1
	return stauses[index]
}

func main() {
	// Define your categories array here

	fmt.Println(StatusCategory(200))
	fmt.Println(StatusCategory(404))
	fmt.Println(StatusCategory(600))
}
