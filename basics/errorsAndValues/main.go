package main

import (
	"errors"
	"fmt"
)

func withdraw(balance, amount int) (int, error) {
	// Your code here
	if amount <= 0 {
		return 0, errors.New("amount must be positive")
	} else if amount > balance {
		return 0, errors.New("insufficient funds")
	}
	remaining := balance - amount
	return remaining, nil
}

func main() {
	newBalance, err := withdraw(100, 30)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New balance:", newBalance)
	}
}
