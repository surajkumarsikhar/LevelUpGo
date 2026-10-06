package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Implement ParseCoordinate and MinMax here
func ParseCoordinate(s string) (x, y int, err error) {
	nums := strings.Split(s, ",")
	if len(nums) != 2 {
		err = fmt.Errorf("invalid format")
		return
	}
	x, err = strconv.Atoi(nums[0])
	if err != nil {
		err = fmt.Errorf("invalid format")
		return
	}
	y, err = strconv.Atoi(nums[1])
	if err != nil {
		err = fmt.Errorf("invalid format")
		return
	}
	return
}

func MinMax(nums []int) (min, max int, ok bool) {
	if len(nums) < 1 {
		return
	}
	min, max = nums[0], nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
		if max < num {
			max = num
		}
	}
	ok = true
	return
}

func main() {
	x, y, err := ParseCoordinate("3,4")
	fmt.Printf("ParseCoordinate(\"3,4\"): x=%d, y=%d, err=%v\n", x, y, err)

	x, y, err = ParseCoordinate("invalid")
	fmt.Printf("ParseCoordinate(\"invalid\"): x=%d, y=%d, err=%v\n", x, y, err)

	min, max, ok := MinMax([]int{5, 2, 8, 1, 9})
	fmt.Printf("MinMax([5, 2, 8, 1, 9]): min=%d, max=%d, ok=%v\n", min, max, ok)

	min, max, ok = MinMax([]int{})
	fmt.Printf("MinMax([]): min=%d, max=%d, ok=%v\n", min, max, ok)
}
