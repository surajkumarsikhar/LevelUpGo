package main

import "fmt"

func main() {
	// Define const AppName with value "LevelUpGo"
	const AppName = "LevelUpGo"
	// Define const MaxUsers with value 100
	const MaxUsers = 100
	// Declare var currentUsers as int (no value - use zero value)
	var currentUsers int

	fmt.Printf("%s: %d/%d users\n", AppName, currentUsers, MaxUsers)
}