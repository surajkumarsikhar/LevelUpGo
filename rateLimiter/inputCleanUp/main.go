package main

import (
	"fmt"
	"strings"
)

// normalizeClientID trims whitespace and converts to uppercase.
func normalizeClientID(id string) string {
	// TODO: Implement
	clean := strings.TrimSpace(id)
	upper := strings.ToUpper(clean)
	return upper
}

// isValidClientID returns true if the ID is non-empty and within maxLength.
func isValidClientID(id string, maxLength int) bool {
	// TODO: Implement
	return (len(id) > 0 && len(id) <= maxLength)
}

// --- Completed from prior lessons ---

func getClientList() []string {
	return []string{"SVC-AUTH", "SVC-USERS", "SVC-BILLING"}
}

func assignTier(clients []string, index int) string {
	return strings.ToUpper(clients[index%len(clients)])
}

func main() {
	// Test normalization
	fmt.Printf("Normalize '  svc-auth  ': %q\n", normalizeClientID("  svc-auth  "))
	fmt.Printf("Normalize 'Svc-Users': %q\n", normalizeClientID("Svc-Users"))

	// Test validation
	fmt.Printf("Valid 'SVC-AUTH' (20): %v\n", isValidClientID("SVC-AUTH", 20))
	fmt.Printf("Valid '' (20): %v\n", isValidClientID("", 20))
	fmt.Printf("Valid 'TOOLONGCLIENTIDTHATEXCEEDSLIMIT' (20): %v\n", isValidClientID("TOOLONGCLIENTIDTHATEXCEEDSLIMIT", 20))
}
