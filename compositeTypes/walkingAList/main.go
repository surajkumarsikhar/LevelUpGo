package main

import "fmt"

func FormatLogs(logs []string) string {
	// Your code here
	result := ""
	for i, msg := range logs {
		if i == (len(logs) - 1) {
			result += fmt.Sprintf("%d. %s", (i + 1), msg)
		} else {
			result += fmt.Sprintf("%d. %s\n", (i + 1), msg)
		}
	}
	return result
}

func main() {
	result := FormatLogs([]string{"server started", "request received", "response sent"})
	// %q shows the string as a quoted literal: each newline appears
	// as a visible \n so you can spot a stray trailing one
	fmt.Printf("%q\n", result)
	fmt.Println("---")
	// Println shows the same string with real line breaks
	fmt.Println(result)
}
