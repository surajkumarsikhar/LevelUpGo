package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func loggingMiddleware(next http.Handler) http.Handler {
	// Your code here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		fmt.Println("Response complete")
	})
}

func authMiddleware(next http.Handler) http.Handler {
	// Your code here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Protected content")
}

func main() {
	// Test logging middleware
	handler := loggingMiddleware(http.HandlerFunc(protectedHandler))
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	fmt.Printf("Logging test: %d\n\n", w.Code)

	// Test auth middleware - no token (should fail)
	handler = authMiddleware(http.HandlerFunc(protectedHandler))
	req = httptest.NewRequest("GET", "/protected", nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	fmt.Printf("Auth (no token): %d\n", w.Code)

	// Test auth middleware - with token (should succeed)
	req = httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer token123")
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	fmt.Printf("Auth (with token): %d\n", w.Code)
}
