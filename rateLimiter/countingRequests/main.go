package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// RequestResult holds the rate limit decision for a single request.
type RequestResult struct {
	ClientID string
	Status   string // "allowed" or "denied"
}

// checkRequests processes a batch of requests against a rate limit.
// priorCounts holds existing counts from earlier in the window.
func checkRequests(requests []string, priorCounts map[string]int, limit int) []RequestResult {
	// TODO: Implement
	count := make(map[string]int, len(priorCounts))
	res := []RequestResult{}
	for k, v := range priorCounts {
		count[k] = v
	}

	for _, req := range requests {
		if count[req] < limit {
			res = append(res, RequestResult{ClientID: req, Status: "allowed"})
			count[req]++
		} else {
			res = append(res, RequestResult{ClientID: req, Status: "denied"})
		}
	}
	return res
}

// --- Completed from prior lessons ---

func getClientList() []string {
	return []string{"SVC-AUTH", "SVC-USERS", "SVC-BILLING"}
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
	return slices.Contains(clients, clientID)
}

func validateRequest(clientID string, maxLength int, clients []string) error {
	if !isValidClientID(clientID, maxLength) {
		return errors.New("invalid client ID format")
	}
	if !isRegisteredClient(clientID, clients) {
		return errors.New("unregistered client")
	}
	return nil
}

func main() {
	priorCounts := map[string]int{"SVC-AUTH": 2}
	requests := []string{"SVC-AUTH", "SVC-USERS", "SVC-AUTH", "SVC-USERS", "SVC-USERS"}
	results := checkRequests(requests, priorCounts, 3)
	for _, r := range results {
		fmt.Printf("[%s:%s] ", r.ClientID, r.Status)
	}
	fmt.Println()
}
