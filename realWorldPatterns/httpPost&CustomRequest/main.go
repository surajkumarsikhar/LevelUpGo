package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser sends a POST request to create a new user.
func CreateUser(baseURL string, user *User) error {
	// Your code here
	url := fmt.Sprintf("%s/users", baseURL)
	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return fmt.Errorf("Failed to create user")
	}
	return nil
}

func main() {
	// Create test server that accepts POST and returns 201
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Missing Content-Type header")
			return
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer server.Close()

	user := &User{Name: "bob", Email: "bob@example.com"}
	err := CreateUser(server.URL, user)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("User created successfully")
}
