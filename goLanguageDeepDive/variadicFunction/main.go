package main

import (
	"fmt"
)

// Implement Sum and Max here
func Sum(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max(nums ...int) (int, error) {
	if len(nums) < 1 {
		return 0, fmt.Errorf("no values provided")
	}
	var maxNum int = nums[0]
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}
	return maxNum, nil
}

func main() {
	fmt.Printf("Sum(): %d\n", Sum())
	fmt.Printf("Sum(1, 2, 3): %d\n", Sum(1, 2, 3))
	fmt.Printf("Sum(10, 20, 30, 40): %d\n", Sum(10, 20, 30, 40))

	max, err := Max(3, 1, 4, 1, 5)
	if err != nil {
		fmt.Printf("Max(3, 1, 4, 1, 5): error - %v\n", err)
	} else {
		fmt.Printf("Max(3, 1, 4, 1, 5): %d\n", max)
	}

	_, err = Max()
	if err != nil {
		fmt.Printf("Max(): error - %v\n", err)
	}
}
