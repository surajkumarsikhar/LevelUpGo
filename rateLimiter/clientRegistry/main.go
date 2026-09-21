package main

import "fmt"

func getClientList() []string {
	// TODO: Return a slice with the clients
	return []string{"SVC-AUTH", "SVC-USERS", "SVC-BILLING"}
}

func assignTier(clients []string, index int) string {
	// TODO: Return clients[index % len(clients)]
	return clients[index%len(clients)]
}

func main() {
	clients := getClientList()
	fmt.Printf("Registry: %d clients\n", len(clients))
	for i := 0; i < 3; i++ {
		fmt.Printf("Tier %d: %s\n", i, assignTier(clients, i))
	}
	fmt.Printf("Wrapped: %s\n", assignTier(clients, len(clients)+2))
}
