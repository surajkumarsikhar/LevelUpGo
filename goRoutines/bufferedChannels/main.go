package main

import (
	"fmt"
	"sync"
)

// TrySend attempts to send value without blocking.
// Returns true if sent, false if channel is full.
func TrySend(ch chan<- int, value int) bool {
	// TODO: Use select with default for non-blocking send
	select {
	case ch <- value:
		return true
	default:
		return false
	}
}

// TryReceive attempts to receive without blocking.
// Returns (value, true) if received, (0, false) if channel empty.
func TryReceive(ch <-chan int) (int, bool) {
	// TODO: Use select with default for non-blocking receive
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
}

// DrainChannel receives all currently buffered values without blocking.
func DrainChannel(ch <-chan int) []int {
	// TODO: Loop TryReceive until channel is empty
	var res []int
	for len(ch) > 0 {
		res = append(res, <-ch)
	}
	return res
}

// WorkerPool launches numWorkers goroutines to process jobs.
// Each worker reads from jobs, squares the value, sends to results.
// Closes results when all workers are done.
func WorkerPool(jobs <-chan int, results chan<- int, numWorkers int) {
	// TODO: Use WaitGroup to track workers
	var wg sync.WaitGroup
	// TODO: For each worker: wg.Add(1), launch goroutine with defer wg.Done()
	// TODO: Each worker: for job := range jobs { results <- job*job }
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				results <- job * job
			}
		}()
	}
	// TODO: Close results after wg.Wait()
	go func() {
		wg.Wait()
		close(results)
	}()
}

func main() {
	// Test TrySend/TryReceive
	ch := make(chan int, 2)
	fmt.Println("TrySend 1:", TrySend(ch, 1)) // true
	fmt.Println("TrySend 2:", TrySend(ch, 2)) // true
	fmt.Println("TrySend 3:", TrySend(ch, 3)) // false (full)

	v, ok := TryReceive(ch)
	fmt.Printf("TryReceive: %d, %v\n", v, ok) // 1, true

	// Test DrainChannel
	fmt.Println("Drain:", DrainChannel(ch)) // [2]

	// Test WorkerPool
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	go WorkerPool(jobs, results, 2)

	sum := 0
	for r := range results {
		sum += r
	}
	fmt.Println("WorkerPool sum:", sum) // 1+4+9+16+25 = 55
}
