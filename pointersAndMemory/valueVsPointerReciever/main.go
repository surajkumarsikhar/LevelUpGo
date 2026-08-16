package main

import "fmt"

type Download struct {
	Received int
	Total    int
}

// Advance records that n more bytes arrived.
func (d *Download) Advance(n int) {
	// Your code here
	d.Received += n
}

// Done reports whether the whole file has arrived.
func (d Download) Done() bool {
	// Your code here
	return d.Received == d.Total
}

func main() {
	dl := Download{Total: 500}
	dl.Advance(200)
	dl.Advance(300)

	fmt.Println(dl.Received) // 500
	fmt.Println(dl.Done())   // true
}
