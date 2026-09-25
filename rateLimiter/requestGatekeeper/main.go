package main

import (
	"fmt"
	"slices"
	"strings"
)

// isRegisteredClient returns true if clientID is found in the clients slice.
func isRegisteredClient(clientID string, clients []string) bool {
	// TODO: Implement
	return slices.Contains(clients, clientID)
}

// validateRequest checks format first, then registration.
// Returns nil if valid, or an error describing the problem.
func validateRequest(clientID string, maxLength int, clients []string) error {
	// TODO: Implement
	if !isValidClientID(clientID, maxLength) {
		return fmt.Errorf("invalid client ID format")
	}
	normalizedClientID := normalizeClientID(clientID)
	if !isRegisteredClient(normalizedClientID, clients) {
		return fmt.Errorf("unregistered client")
	}
	return nil
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

func main() {
	clients := getClientList()

	fmt.Printf("SVC-AUTH registered: %v\n", isRegisteredClient("SVC-AUTH", clients))
	fmt.Printf("SVC-UNKNOWN registered: %v\n", isRegisteredClient("SVC-UNKNOWN", clients))

	fmt.Printf("SVC-AUTH: %v\n", validateRequest("SVC-AUTH", 20, clients))
	fmt.Printf("empty: %v\n", validateRequest("", 20, clients))
	fmt.Printf("SVC-UNKNOWN: %v\n", validateRequest("SVC-UNKNOWN", 20, clients))
}
