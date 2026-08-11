package main

import "fmt"

// Write the bookSeats function here
// It takes available and requested (both int) and returns (int, bool)
func bookSeats(available, requested int) (int, bool){
	booked := available - requested
	if booked < 0 || requested <= 0 {
		return 0, false
	}
	return booked,true
}

func main() {
	remaining, ok := bookSeats(50, 10)
	fmt.Println(remaining, ok)

	remaining, ok = bookSeats(5, 20)
	fmt.Println(remaining, ok)

	remaining, ok = bookSeats(10, 0)
	fmt.Println(remaining, ok)
}