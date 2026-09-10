package main

import (
	"fmt"
	"strings"
)

// Base interfaces
type Reader interface{ Read() string }
type Writer interface{ Write(data string) error }
type Closer interface{ Close() error }

// TODO: Create ReadWriteCloser that embeds Reader, Writer, Closer
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

type Buffer struct {
	data   []string
	closed bool
}

// TODO: Implement Read() - return strings.Join(b.data, " ")
func (b Buffer) Read() string {
	return strings.Join(b.data, " ")
}

// TODO: Implement Write() - if closed return error, else append
func (b *Buffer) Write(data string) error {
	if b.closed {
		return fmt.Errorf("closed")
	}
	b.data = append(b.data, data)
	return nil
}

// TODO: Implement Close() - if already closed return error, else set closed
func (b *Buffer) Close() error {
	if b.closed {
		return fmt.Errorf("closed")
	}
	b.closed = true
	return nil
}

func main() {}
