package main

import (
	"fmt"
	"sync"
)

// These are provided for you
func GenerateNumbers(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := range n {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func Square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

// Implement these
func SumChannel(in <-chan int) <-chan int {
	// TODO: Sum all input values, send the total, close output
	ch := make(chan int)
	go func() {
		sum := 0
		for v := range in {
			sum += v
		}
		ch <- sum
		close(ch)
	}()
	return ch
}

func Merge(channels ...<-chan int) <-chan int {
	// TODO: Forward all values from all input channels to output
	out := make(chan int)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Go(func() {
			for v := range ch {
				out <- v
			}
		})
	}
	// TODO: Close output when all inputs are exhausted
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// Pipeline: Generate -> Square -> Sum
	result := <-SumChannel(Square(GenerateNumbers(5)))
	fmt.Printf("Sum of squares 0-4: %d\n", result) // 0+1+4+9+16 = 30

	// Merge multiple generators
	merged := Merge(GenerateNumbers(3), GenerateNumbers(3))
	sum := 0
	for n := range merged {
		sum += n
	}
	fmt.Printf("Merged sum: %d\n", sum) // (0+1+2)*2 = 6
}
