package main

import (
	"fmt"
	"strings"
)

// AllowRequest processes a single request through the rate limiter.
// Returns a RequestResult and an error if the client is invalid.
func (rl *RateLimiter) AllowRequest(clientID string, now int) (RequestResult, error) {
	// TODO: Implement the 5 steps
	normalizedID := normalizeClientID(clientID)
	if _, ok := validateRequest(normalizedID, 20, getClientList()); !ok {
		return RequestResult{}, fmt.Errorf("rejected")
	}

	if rl.IsWindowExpired(now) {
		rl.WindowStart = now
		rl.RequestCounts = make(map[string]int)
	}

	if rl.RemainingRequests(normalizedID) == 0 {
		rl.TotalChecked++
		return RequestResult{
			ClientID: normalizedID,
			Status:   "denied",
		}, nil
	}
	rl.RequestCounts[normalizedID]++
	rl.TotalChecked++
	return RequestResult{ClientID: normalizedID, Status: "allowed"}, nil
}

// --- Completed from prior lessons ---

func getClientList() []string {
	return []string{
		"SVC-AUTH", "SVC-USERS", "SVC-BILLING",
	}
}

func assignTier(clients []string, index int) string {
	return strings.ToUpper(clients[index%len(clients)])
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

func checkRequests(requests []string, windowCounts map[string]int, limit int) []RequestResult {
	results := make([]RequestResult, len(requests))
	batchCounts := map[string]int{}
	for i, clientID := range requests {
		total := windowCounts[clientID] + batchCounts[clientID]
		if total < limit {
			results[i] = RequestResult{ClientID: clientID, Status: "allowed"}
			batchCounts[clientID]++
		} else {
			results[i] = RequestResult{ClientID: clientID, Status: "denied"}
		}
	}
	return results
}

func formatResults(results []RequestResult) string {
	parts := make([]string, len(results))
	for i, r := range results {
		parts[i] = fmt.Sprintf("[%s:%s]", r.ClientID, r.Status)
	}
	return strings.Join(parts, " ")
}

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
		WindowStart:   0,
		RequestCounts: make(map[string]int),
		TotalChecked:  0,
	}
}

func (rl *RateLimiter) IsWindowExpired(now int) bool {
	return now >= rl.WindowStart+rl.WindowSize
}

func (rl *RateLimiter) RemainingRequests(clientID string) int {
	remaining := rl.MaxRequests - rl.RequestCounts[clientID]
	if remaining < 0 {
		return 0
	}
	return remaining
}

func main() {
	rl := NewRateLimiter(3, 60)

	ids := []string{"SVC-AUTH", "SVC-AUTH", "SVC-AUTH", "SVC-AUTH"}
	for _, id := range ids {
		result, err := rl.AllowRequest(id, 100)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("%s: %s (remaining: %d)\n", result.ClientID, result.Status, rl.RemainingRequests(id))
	}

	fmt.Printf("\nTotal checked: %d\n", rl.TotalChecked)

	fmt.Println("\n--- Window expires ---")
	result, _ := rl.AllowRequest("SVC-AUTH", 200)
	fmt.Printf("%s: %s\n", result.ClientID, result.Status)
}
