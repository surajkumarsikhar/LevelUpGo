package main

import (
	"fmt"
	"strings"
)

// formatResults converts a slice of RequestResult into a readable string.
// Each result is formatted as [ClientID:Status], joined by spaces.
func formatResults(results []RequestResult) string {
	// TODO: Implement
	formattedArr := []string{}
	for _, result := range results {
		formattedArr = append(formattedArr, fmt.Sprintf("[%s:%s]", result.ClientID, result.Status))
	}
	return strings.Join(formattedArr, " ")
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

func main() {
	results := checkRequests(
		[]string{"SVC-AUTH", "SVC-USERS", "SVC-AUTH"},
		map[string]int{"SVC-AUTH": 2},
		3,
	)
	fmt.Println(formatResults(results))
}
