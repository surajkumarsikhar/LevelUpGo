package main

import "fmt"

type User struct {
	ID   string
	Name string
}

type UserStore struct {
	users map[string]User
}

// Write Find as a method on *UserStore. It takes an id and returns a
// pointer to the matching user, or nil when the store has no such user.
func (u *UserStore) Find(id string) *User {
	user, ok := u.users[id]
	if !ok {
		return nil
	}
	return &user
}

func main() {
	s := &UserStore{users: map[string]User{
		"u1": {ID: "u1", Name: "Ada"},
		"u2": {ID: "u2", Name: "Linus"},
	}}

	if u := s.Find("u1"); u != nil {
		fmt.Println("found", u.Name)
	}
	if u := s.Find("u9"); u == nil {
		fmt.Println("u9 not found")
	}
}
