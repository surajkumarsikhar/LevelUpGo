package main

import (
	"errors"
	"fmt"
)

// AsInt safely converts interface{} to int
func AsInt(v interface{}) (int, bool) {
	// TODO: Use type assertion with comma-ok idiom
	// val, ok := v.(int)
	val, ok := v.(int)
	// return val, ok
	return val, ok
}

// AsString safely converts interface{} to string
func AsString(v interface{}) (string, bool) {
	// TODO: Use type assertion with comma-ok idiom
	val, ok := v.(string)
	return val, ok
}

// Add adds two interface{} values if they're both numbers
func Add(a, b interface{}) (float64, error) {
	// TODO: Use type switch to convert a to float64
	// TODO: Use type switch to convert b to float64
	// TODO: Return error if either is not int or float64
	var newa, newb float64
	switch x := a.(type) {
	case int:
		newa = float64(x)
	case float64:
		newa = x
	default:
		return 0, errors.New("unsupported types")
	}
	switch x := b.(type) {
	case int:
		newb = float64(x)
	case float64:
		newb = x
	default:
		return 0, errors.New("unsupported types")
	}

	return (newa + newb), nil

}

func main() {
	val, ok := AsInt(42)
	fmt.Printf("AsInt(42): %d, %v\n", val, ok)

	val, ok = AsInt("hello")
	fmt.Printf("AsInt(\"hello\"): %d, %v\n", val, ok)

	str, ok := AsString("world")
	fmt.Printf("AsString(\"world\"): %s, %v\n", str, ok)

	sum, err := Add(10, 20)
	fmt.Printf("Add(10, 20): %.1f, %v\n", sum, err)

	sum, err = Add(10.5, 20.5)
	fmt.Printf("Add(10.5, 20.5): %.1f, %v\n", sum, err)

	sum, err = Add(10, "x")
	fmt.Printf("Add(10, \"x\"): %.1f, %v\n", sum, err)
}
