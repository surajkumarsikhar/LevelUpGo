package main

import (
	"fmt"
	"slices"
)

// TODO: Implement UniqueStrings (concrete, no generics) and Keys (generic)
func UniqueStrings(str []string) []string {
	seen := make(map[string]bool)
	newStr := []string{}

	for _, s := range str {
		if !seen[s] {
			seen[s] = true
			newStr = append(newStr, s)
		}
	}
	return newStr
}

func Keys[K comparable, V any](m map[K]V) []K {
	keyMap := []K{}
	for key := range m {
		keyMap = append(keyMap, key)
	}
	return keyMap
}

func main() {
	// Concrete function for a single type
	fruits := []string{"apple", "banana", "apple", "cherry", "banana"}
	unique := UniqueStrings(fruits)
	slices.Sort(unique)
	fmt.Printf("Unique strings: %v\n", unique)

	// Generic function works with any map
	intMap := map[int]string{1: "one", 2: "two", 3: "three"}
	intKeys := Keys(intMap)
	slices.Sort(intKeys)
	fmt.Printf("Int keys: %v\n", intKeys)

	strMap := map[string]int{"a": 1, "b": 2, "c": 3}
	strKeys := Keys(strMap)
	slices.Sort(strKeys)
	fmt.Printf("String keys: %v\n", strKeys)
}
