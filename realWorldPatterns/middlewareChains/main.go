package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func recoveryMiddleware(next http.Handler) http.Handler {
	// Your code here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Errorf("Panic caught: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func chain(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	// Your code here
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}

	return h
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("something went wrong!")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handling: %s\n", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Test recovery middleware with panic
	handler := recoveryMiddleware(http.HandlerFunc(panicHandler))
	req := httptest.NewRequest("GET", "/panic", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	fmt.Printf("Panic handler: %d\n", w.Code)

	// Test chain function
	chained := chain(
		http.HandlerFunc(okHandler),
		recoveryMiddleware,
		loggingMiddleware,
	)
	req = httptest.NewRequest("GET", "/ok", nil)
	w = httptest.NewRecorder()
	chained.ServeHTTP(w, req)
	fmt.Printf("Chained ok: %d\n", w.Code)
}
