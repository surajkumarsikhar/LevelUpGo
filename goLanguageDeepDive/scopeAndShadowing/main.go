package scopeandshadowing
package main

import (
	"errors"
	"fmt"
)

// BUG: This should return 10, but returns 0
// Fix the shadowing issue
func BuggyCounter() int {
	count := 0

	for range 10 {
		count = count + 1  // BUG: shadows outer count
	}

	return count
}

// BUG: This should return the error, but returns nil
// Fix the shadowing issue
func BuggyError() error {
	var err error
	var result int
	if true {
		result, err = doWork()  // BUG: shadows outer err
		_ = result
	}

	return err  // Always nil!
}

func doWork() (int, error) {
	return 0, errors.New("something went wrong")
}

// TODO: Implement correctly - sum numbers 1 to 5
// Use proper scoping (no shadowing bugs)
func SumWithScope() int {
	// Sum numbers 1 through 5
	sum := 0
	for i := 1 ; i <= 5 ; i++ {
		sum+=i
	}
	return sum
}

func main() {
	fmt.Printf("BuggyCounter (fixed): %d\n", BuggyCounter())

	err := BuggyError()
	if err != nil {
		fmt.Printf("BuggyError (fixed): %v\n", err)
	} else {
		fmt.Println("BuggyError (fixed): nil - still broken!")
	}

	fmt.Printf("SumWithScope: %d\n", SumWithScope())
}