package main

import "fmt"

func GenerateNumbers(n int) <-chan int {
	// TODO: Create channel, launch goroutine to send 0 to n-1, close, return channel
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
	// TODO: Create output channel, launch goroutine to square values, close when done
	ch := make(chan int)
	go func() {
		for num := range in {
			ch <- num * num
		}
		close(ch)
	}()
	return ch
}

func main() {
	// Test GenerateNumbers
	for n := range GenerateNumbers(5) {
		fmt.Printf("%d ", n)
	}
	fmt.Println() // 0 1 2 3 4

	// Test pipeline: Generate -> Square
	for sq := range Square(GenerateNumbers(5)) {
		fmt.Printf("%d ", sq)
	}
	fmt.Println() // 0 1 4 9 16
}
