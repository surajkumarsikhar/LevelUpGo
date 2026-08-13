package main

import "fmt"

func AddLogs(logBuffer []string, newLines []string) (oldCap, newCap int, resized bool) {
	// Your code here
	oldCap = cap(logBuffer)
	logBuffer = append(logBuffer, newLines...)
	newCap = cap(logBuffer)
	resized = oldCap != newCap
	return
}

func main() {
	spacious := make([]string, 0, 10)
	fmt.Println(AddLogs(spacious, []string{"boot", "ready", "listening"})) // 10 10 false

	tight := make([]string, 0, 2)
	fmt.Println(AddLogs(tight, []string{"boot", "ready", "listening"})) // 2 <bigger> true
}
