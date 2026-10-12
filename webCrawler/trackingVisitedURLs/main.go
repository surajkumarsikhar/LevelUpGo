package main

import (
	"fmt"
	"strings"
	"sync"
)

// normalizeURL converts a URL to a canonical form.
// - Convert to lowercase
// - Remove trailing slashes
func normalizeURL(rawURL string) string {
	// TODO: Implement
	n := strings.ToLower(rawURL)
	n = strings.TrimRight(n, "/")
	return n
}

// VisitedTracker tracks which URLs have been visited in a thread-safe way.
type VisitedTracker struct {
	visited map[string]bool
	mu      sync.RWMutex
}

// NewVisitedTracker creates a new empty tracker.
func NewVisitedTracker() *VisitedTracker {
	// TODO: Implement
	return nil
}

// HasVisited returns true if the URL has been visited before.
// Uses RLock for concurrent read access.
func (v *VisitedTracker) HasVisited(url string) bool {
	// TODO: Implement
	return false
}

// MarkVisited marks a URL as visited.
// Returns true if the URL was newly added (not visited before).
// Returns false if it was already visited.
// This must be atomic — check and set in one locked operation.
func (v *VisitedTracker) MarkVisited(url string) bool {
	// TODO: Implement
	return false
}

// Count returns the number of unique URLs visited.
func (v *VisitedTracker) Count() int {
	// TODO: Implement
	return 0
}

func main() {
	tracker := NewVisitedTracker()

	// Mark some URLs
	urls := []string{
		"http://gadgetshop/shop/",
		"http://gadgetshop/shop/products",
		"http://gadgetshop/shop/about",
		"http://gadgetshop/shop/",          // duplicate
		"http://gadgetshop/shop/products/", // duplicate (trailing slash)
	}

	for _, u := range urls {
		isNew := tracker.MarkVisited(u)
		if isNew {
			fmt.Printf("NEW:  %s\n", u)
		} else {
			fmt.Printf("SKIP: %s\n", u)
		}
	}

	fmt.Printf("\nTotal unique URLs: %d\n", tracker.Count())

	// Check specific URLs
	fmt.Printf("Has visited /shop/products: %v\n", tracker.HasVisited("http://gadgetshop/shop/products"))
	fmt.Printf("Has visited /shop/missing: %v\n", tracker.HasVisited("http://gadgetshop/shop/missing"))

	// Test concurrent safety
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			url := fmt.Sprintf("http://gadgetshop/shop/page-%d", n%10)
			tracker.MarkVisited(url)
		}(i)
	}
	wg.Wait()

	fmt.Printf("After concurrent adds: %d unique URLs\n", tracker.Count())
	fmt.Println("Done!")
}
