package main

import (
	"fmt"
	"strings"
	"sync"
)

type RateLimiter struct {
	MaxRequests   int
	WindowSize    int
	WindowStart   int
	RequestCounts map[string]int
	TotalChecked  int
	mu            sync.Mutex
	// TODO: add a sync.Mutex field named mu
}

func NewRateLimiter(maxRequests, windowSize int) *RateLimiter {
	return &RateLimiter{
		MaxRequests:   maxRequests,
		WindowSize:    windowSize,
		RequestCounts: make(map[string]int),
	}
}

func (rl *RateLimiter) IsWindowExpired(now int) bool {
	return now >= rl.WindowStart+rl.WindowSize
}

func (rl *RateLimiter) AllowRequest(clientID string, now int) (RequestResult, error) {
	clientID = normalizeClientID(clientID)

	msg, ok := validateRequest(clientID, 20, getClientList())
	if !ok {
		return RequestResult{}, fmt.Errorf("rejected: %s", msg)
	}

	// TODO: lock rl.mu and defer unlock before touching shared state
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.IsWindowExpired(now) {
		rl.WindowStart = now
		rl.RequestCounts = make(map[string]int)
	}

	if rl.RequestCounts[clientID] >= rl.MaxRequests {
		rl.TotalChecked++
		return RequestResult{ClientID: clientID, Status: "denied"}, nil
	}

	rl.RequestCounts[clientID]++
	rl.TotalChecked++
	return RequestResult{ClientID: clientID, Status: "allowed"}, nil
}

// --- Completed from prior lessons ---

func getClientList() []string {
	return []string{"SVC-AUTH", "SVC-USERS", "SVC-BILLING"}
}

func normalizeClientID(id string) string {
	return strings.ToUpper(strings.TrimSpace(id))
}

func isValidClientID(id string, maxLength int) bool {
	return len(id) > 0 && len(id) <= maxLength
}

func isRegisteredClient(clientID string, clients []string) bool {
	for _, c := range clients {
		if c == clientID {
			return true
		}
	}
	return false
}

func validateRequest(clientID string, maxLength int, clients []string) (string, bool) {
	if !isValidClientID(clientID, maxLength) {
		return "invalid client ID format", false
	}
	if !isRegisteredClient(clientID, clients) {
		return "unregistered client", false
	}
	return "", true
}

type RequestResult struct {
	ClientID string
	Status   string
}

func main() {
	rl := NewRateLimiter(3, 60)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := rl.AllowRequest("SVC-AUTH", 100)
			if err == nil {
				fmt.Printf("%s: %s\n", result.ClientID, result.Status)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("TotalChecked: %d\n", rl.TotalChecked)
}
