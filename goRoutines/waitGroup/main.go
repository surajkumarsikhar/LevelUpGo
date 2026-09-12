package main

import (
	"fmt"
	"sync"
)

// ProcessTasks processes each task by calling the processor function concurrently.
// It must use goroutines with WaitGroup to run processors in parallel.
// The processor writes its result to the provided channel.
func ProcessTasks(tasks []string, results chan<- string, processor func(task string, out chan<- string)) {
	// TODO: Create WaitGroup
	var wg sync.WaitGroup
	// TODO: For each task, use wg.Go(func() { processor(task, results) })
	for _, task := range tasks {
		wg.Go(func() {
			processor(task, results)
		})
	}
	// TODO: wg.Wait() at the end
	wg.Wait()
}

func main() {
	tasks := []string{"download-file", "send-email", "generate-report"}
	results := make(chan string, len(tasks))

	processor := func(task string, out chan<- string) {
		out <- "processed: " + task
	}

	ProcessTasks(tasks, results, processor)
	close(results)

	fmt.Println("Results:")
	for r := range results {
		fmt.Printf("  %s\n", r)
	}
}
