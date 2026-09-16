package racedetection
package main

import (
	"sync"
	"fmt"
)

// Counter is a thread-safe counter
type Counter struct {
	// TODO: Add a sync.Mutex field
	value int
	mu sync.Mutex
}

// Inc increments the counter by 1
func (c *Counter) Inc() {
	// TODO: Lock the mutex before incrementing
	// TODO: Use defer to unlock
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current counter value
func (c *Counter) Value() int {
	// TODO: Lock the mutex before reading
	// TODO: Use defer to unlock
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	for range 100 {
		wg.Go(func() {
			c.Inc()
		})
	}

	wg.Wait()
	fmt.Printf("Counter value: %d\n", c.Value())
}