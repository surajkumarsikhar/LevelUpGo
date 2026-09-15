package main

import "fmt"

func Produce(ch chan<- int, n int) {
	// TODO: Send 1 to n, close when done
	for i := 1; i <= n; i++ {
		ch <- i
	}
	close(ch)
}

func Consume(ch <-chan int, result chan<- int) {
	// TODO: Sum all values, send sum to result
	sum := 0
	for v := range ch {
		sum += v
	}
	result <- sum
	close(result)
}

func Transform(in <-chan int, out chan<- int) {
	// TODO: Multiply by 10, close when done
	for v := range in {
		out <- v * 10
	}
	close(out)
}

func main() {
	// Test Produce -> Consume
	ch := make(chan int)
	result := make(chan int)

	go Produce(ch, 5)
	go Consume(ch, result)

	fmt.Println("Sum of 1-5:", <-result) // 15

	// Test Produce -> Transform -> Consume
	ch1 := make(chan int)
	ch2 := make(chan int)
	result2 := make(chan int)

	go Produce(ch1, 5)
	go Transform(ch1, ch2)
	go Consume(ch2, result2)

	fmt.Println("Sum of (1-5)*10:", <-result2) // 150
}
