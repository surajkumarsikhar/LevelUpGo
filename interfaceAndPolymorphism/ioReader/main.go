package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// TODO: Define ROT13Reader struct with an io.Reader field named 'r'
type ROT13Reader struct {
	r io.Reader
}

// TODO: Implement Read method on *ROT13Reader that:
//   - Reads from the wrapped reader into p
//   - Applies rot13 to each byte read
//   - Returns n and err from the wrapped reader
func (rot *ROT13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}
	return
}

// TODO: Implement rot13(b byte) byte function that:
//   - Transforms A-M to N-Z (add 13)
//   - Transforms N-Z to A-M (subtract 13)
//   - Transforms a-m to n-z (add 13)
//   - Transforms n-z to a-m (subtract 13)
//   - Returns unchanged for non-letters
func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return byte((int(b-'A')+13)%26) + 'A'
	} else if b >= 'a' && b <= 'z' {
		return byte((int(b-'a')+13)%26) + 'a'
	}
	return b
}

func main() {
	s := strings.NewReader("Hello, World!")
	r := ROT13Reader{r: s}

	io.Copy(os.Stdout, &r)
	fmt.Println()

	// Demonstrate ROT13 is reversible
	s2 := strings.NewReader("Uryyb, Jbeyq!")
	r2 := ROT13Reader{r: s2}
	io.Copy(os.Stdout, &r2)
	fmt.Println()
}
