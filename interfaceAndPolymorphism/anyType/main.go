package main

import "fmt"

// PrintAny prints any value with its type
func PrintAny(v interface{}) {
	// TODO: Use type switch to print type and value
	switch val := v.(type) {
	// case int: fmt.Printf("int: %d\n", val)
	case int:
		fmt.Printf("int: %d\n", val)
	// case string: fmt.Printf("string: %s\n", val)
	case string:
		fmt.Printf("string: %s\n", val)
	// case []int: fmt.Printf("[]int: %v\n", val)
	case []int:
		fmt.Printf("[]int: %v\n", val)
	// case bool: fmt.Printf("bool: %v\n", val)
	case bool:
		fmt.Printf("bool: %v\n", val)
	// default: fmt.Printf("%T: %v\n", val, val)
	default:
		fmt.Printf("%T: %v\n", val, val)
	}
}

// CompareAny compares two values of any type
func CompareAny(a, b interface{}) bool {
	// TODO: Return a == b
	return a == b
}

// Length returns length of strings/slices, -1 for others
func Length(v interface{}) int {
	// TODO: Use type switch
	// case string: return len(val)
	// case []int: return len(val)
	// default: return -1
	switch x := v.(type) {
	case string:
		return len(x)
	case []int:
		return len(x)
	default:
		return -1
	}
}

func main() {
	PrintAny(42)
	PrintAny("hello")
	PrintAny([]int{1, 2, 3})
	PrintAny(true)

	fmt.Println(CompareAny(42, 42))
	fmt.Println(CompareAny("hello", "hello"))
	fmt.Println(CompareAny(42, "42"))

	fmt.Println(Length("hello"))
	fmt.Println(Length([]int{1, 2, 3}))
	fmt.Println(Length(42))
}
