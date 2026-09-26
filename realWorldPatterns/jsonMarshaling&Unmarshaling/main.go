package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
}

// UserToJSON converts a User struct to a JSON string.
func UserToJSON(user User) (string, error) {
	// Your code here
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

// JSONToUser parses a JSON string into a User struct.
func JSONToUser(jsonStr string) (User, error) {
	// Your code here
	user := User{}
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func main() {
	// Test marshaling
	user := User{
		ID:       1,
		Username: "alice",
		Email:    "alice@example.com",
		Active:   true,
	}

	jsonStr, err := UserToJSON(user)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(jsonStr)

	// Test unmarshaling
	parsed, err := JSONToUser(jsonStr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("User: %s (%s)\n", parsed.Username, parsed.Email)
}
