package main

import "fmt"

type User struct {
	Email string
}

type Request struct {
	Path string
	User *User // nil when no user is attached
}

// UserLabel returns a display label for the request's user. Right now it
// dereferences r.User without checking, so a request with no user panics.
// Guard the deref so a nil User returns "anonymous" instead.
func (r *Request) UserLabel() string {
	if r.User == nil {
		return "anonymous"
	}
	return r.User.Email
}

func main() {
	// A request that has a user. The submit-time tests also check a
	// request with no user (r.User == nil), which is the case that crashes.
	req := &Request{Path: "/dashboard", User: &User{Email: "ada@example.com"}}
	fmt.Println(req.UserLabel())
}
