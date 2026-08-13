package main

import "fmt"

// Define your Config struct here
type Config struct {
	Host     string
	Port     int
	MaxConns int
}

// Write your NewConfig function here
func NewConfig(host string, port int) (config Config) {
	config.Host = host
	config.Port = port
	config.MaxConns = 100
	return
}

func main() {
	c := NewConfig("localhost", 8080)
	fmt.Printf("%s:%d max=%d\n", c.Host, c.Port, c.MaxConns)
}
