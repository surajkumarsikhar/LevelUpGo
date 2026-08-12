package main

import "fmt"

// Write your function here
func AddPermission(permissions []string, permission string) []string {
	permissions = append(permissions, permission)
	return permissions
}

func main() {
	result := AddPermission([]string{"read"}, "write")
	fmt.Println(result)
}
