package main

import (
	"fmt"
	"time"
)

func WithTimeout(ch <-chan int, timeout time.Duration) (int, bool) {
	// TODO: Use select with ch and time.After
	select {
	case v := <-ch:
		return v, true
	case <-time.After(1 * time.Second):
		return 0, false
	}
}

func TryReceive(ch <-chan int) (int, bool) {
	// TODO: Use select with default
	select {
	case v := <-ch:
		return v, true
	default:
		return 0, false
	}
}

func FanIn(ch1, ch2 <-chan int, done <-chan bool) <-chan int {
	// TODO: Multiplex ch1 and ch2, stop on done
	res := make(chan int)
	go func() {
		defer close(res)
		for {
			select {
			case v := <-ch1:
				res <- v
			case v := <-ch2:
				res <- v
			case <-done:
				return
			}
		}
	}()
	return res
}

func main() {
	ch1 := make(chan int, 1)
	ch1 <- 42
	val, ok := WithTimeout(ch1, 1*time.Second)
	fmt.Printf("WithTimeout (success): %d, %v\n", val, ok)

	ch2 := make(chan int)
	val, ok = WithTimeout(ch2, 100*time.Millisecond)
	fmt.Printf("WithTimeout (timeout): %d, %v\n", val, ok)

	ch3 := make(chan int, 1)
	ch3 <- 99
	val, ok = TryReceive(ch3)
	fmt.Printf("TryReceive (available): %d, %v\n", val, ok)

	ch4 := make(chan int)
	val, ok = TryReceive(ch4)
	fmt.Printf("TryReceive (empty): %d, %v\n", val, ok)
}
