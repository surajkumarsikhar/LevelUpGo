package main

import "fmt"

type Stats struct {
	Count int
	Total int
	Max   int
}

// Returns *Stats, so it escapes to the heap. Refactor to return Stats by value.
func collect(latencies []int) Stats {
	var s Stats
	for _, ms := range latencies {
		accumulate(&s, ms)
	}
	return s
}

// Mutates, so it keeps the pointer. Leave this alone.
func accumulate(s *Stats, ms int) {
	s.Count++
	s.Total += ms
	if ms > s.Max {
		s.Max = ms
	}
}

func main() {
	s := collect([]int{120, 80, 200, 95})
	fmt.Printf("count=%d total=%d max=%d\n", s.Count, s.Total, s.Max)
}
