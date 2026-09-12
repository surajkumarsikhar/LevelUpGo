package main

import (
	"fmt"
	"sync"
)

// Counter is a thread-safe counter
type Counter struct {
	// TODO: Add mu sync.Mutex
	// TODO: Add value int
	mu    sync.Mutex
	value int
}

// Inc increments the counter by 1
func (c *Counter) Inc() {
	// TODO: Lock, increment, unlock (use defer)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current counter value
func (c *Counter) Value() int {
	// TODO: Lock, read, unlock (use defer)
	c.mu.Lock()
	defer c.mu.Unlock()
	res := c.value
	return res
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
