package main

import (
	"fmt"
	"strings"
)

type RateLimiter struct {
	MaxRequests   int
	WindowSize    int
	WindowStart   int
	RequestCounts map[string]int
	TotalChecked  int
}

func NewRateLimiter(maxRequests, windowSize int) *RateLimiter {
	// TODO: return a *RateLimiter with the fields set and the map initialized
	return &RateLimiter{
		MaxRequests:   maxRequests,
		WindowSize:    windowSize,
		RequestCounts: make(map[string]int),
	}
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

func main() {
	// After you've defined RateLimiter and NewRateLimiter, try them out:
	// rl := NewRateLimiter(3, 60)
	// fmt.Printf("Max: %d, Window: %d\n", rl.MaxRequests, rl.WindowSize)
	_ = fmt.Sprintf
	_ = strings.ToUpper
}
