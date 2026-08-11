package main

import "fmt"

func dayType(day string) string {
	// Use a switch on day to return "weekday", "weekend", or "unknown"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday" :
		return "weekday"
	case "Saturday", "Sunday" :
		return "weekend"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println(dayType("Monday"))
	fmt.Println(dayType("Saturday"))
}