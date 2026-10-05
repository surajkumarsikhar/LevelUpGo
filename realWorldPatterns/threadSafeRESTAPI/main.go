package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
)

func (s *UserStore) GetAll() []User {
	// Your code here
	userList := []User{}
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.users {
		userList = append(userList, user)
	}
	return userList
}

func (s *UserStore) Create(user User) User {
	// Your code here
	s.mu.Lock()
	s.mu.Unlock()
	user.ID = s.nextID
	s.nextID++
	s.users[user.ID] = user
	return user
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserStore struct {
	mu     sync.RWMutex
	users  map[int]User
	nextID int
}

func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]User),
		nextID: 1,
	}
}

var store = NewUserStore()

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.GetAll())
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	created := store.Create(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", handleGetUsers)
	mux.HandleFunc("POST /users", handleCreateUser)
}

func main() {
	mux := http.NewServeMux()
	setupRoutes(mux)

	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	fmt.Printf("GET: %d\n", w.Code)

	body := strings.NewReader(`{"name":"Charlie"}`)
	req = httptest.NewRequest("POST", "/users", body)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	fmt.Printf("POST: %d\n", w.Code)
}
