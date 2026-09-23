package main

import (
	"fmt"
)

// TODO: Implement these two methods on *RateLimiter.

func (rl *RateLimiter) IsWindowExpired(now int) bool {
	// TODO: return true if now is past the end of the current window
	return now >= (rl.WindowStart + rl.WindowSize)
}

func (rl *RateLimiter) RemainingRequests(clientID string) int {
	// TODO: return MaxRequests minus how many the client has used, clamped to 0
	val, _ := rl.RequestCounts[clientID]
	remainingReq := rl.MaxRequests - val
	if remainingReq < 0 {
		return 0
	}
	return remainingReq
}

// --- Completed from prior lessons ---

type RateLimiter struct {
	MaxRequests   int
	WindowSize    int
	WindowStart   int
	RequestCounts map[string]int
	TotalChecked  int
}

func NewRateLimiter(maxRequests, windowSize int) *RateLimiter {
	return &RateLimiter{
		MaxRequests:   maxRequests,
		WindowSize:    windowSize,
		RequestCounts: make(map[string]int),
	}
}

func main() {
	rl := NewRateLimiter(5, 60)
	rl.WindowStart = 100
	fmt.Println(rl.IsWindowExpired(150))
	fmt.Println(rl.RemainingRequests("SVC-AUTH"))
}
