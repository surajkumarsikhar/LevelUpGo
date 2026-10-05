package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}
var nextID = 3

// handler: return users
func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	return
}

// handler: process user
func handlePostUsers(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	lastID := users[len(users)-1].ID
	user.ID = lastID + 1
	users = append(users, user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// setupRoutes registers handlers for the users endpoint.
func setupRoutes(mux *http.ServeMux) {
	// Your code here
	mux.HandleFunc("GET /users", handleGetUsers)
	mux.HandleFunc("POST /users", handlePostUsers)
}

func main() {
	mux := http.NewServeMux()
	setupRoutes(mux)

	// Test GET
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	fmt.Printf("GET: %d\n", w.Code)

	// Test POST
	body := strings.NewReader(`{"name":"Charlie"}`)
	req = httptest.NewRequest("POST", "/users", body)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	fmt.Printf("POST: %d\n", w.Code)
}
