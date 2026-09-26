package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// FetchUser fetches a user from the API by ID.
func FetchUser(baseURL string, id int) (*User, error) {
	// Your code here
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("%s/users/%d", baseURL, id)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned %d", resp.StatusCode)
	}

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}
	return &user, nil
}

func main() {
	// Create a test server that returns mock user data
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := User{ID: 1, Name: "alice", Email: "alice@example.com"}
		json.NewEncoder(w).Encode(user)
	}))
	defer server.Close()

	user, err := FetchUser(server.URL, 1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User: %s (%s)\n", user.Name, user.Email)
}
