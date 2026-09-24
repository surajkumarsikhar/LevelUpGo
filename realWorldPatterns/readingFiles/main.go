package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CountLogLevels reads a log file and counts entries by level.
// Returns a map with keys: ERROR, WARN, INFO, DEBUG
func CountLogLevels(filename string) (map[string]int, error) {
	// Your code here
	file, err := os.Open(filename)
	res := make(map[string]int)
	if err != nil {
		return res, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		key := strings.Split(line, " ")[0]
		normalizedKey := key[1 : len(key)-1]
		res[normalizedKey]++
	}
	return res, nil
}

func main() {
	// Create sample log file for testing
	createSampleLog()

	counts, err := CountLogLevels("app.log")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("ERROR: %d\n", counts["ERROR"])
	fmt.Printf("WARN: %d\n", counts["WARN"])
	fmt.Printf("INFO: %d\n", counts["INFO"])
}

func createSampleLog() {
	content := `[INFO] 2024-01-15 10:00:00 - Server started
[ERROR] 2024-01-15 10:05:00 - Database connection failed
[INFO] 2024-01-15 10:06:00 - Retrying connection
[WARN] 2024-01-15 10:07:00 - High memory usage
[INFO] 2024-01-15 10:08:00 - Connection restored
[ERROR] 2024-01-15 10:10:00 - Request timeout`
	os.WriteFile("app.log", []byte(content), 0644)
}
