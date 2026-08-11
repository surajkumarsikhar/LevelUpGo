package main

import "fmt"

func checkSystem(cpuHigh bool, memoryHigh bool, diskFull bool) string {
	status := "OK"
	// Your code here: reassign `status` to "CRITICAL" or "WARNING" based on
	// the rules in the instructions. Leave it as "OK" otherwise.
	if cpuHigh && memoryHigh{
		status = "CRITICAL"
	} else if (cpuHigh || memoryHigh || diskFull){
		status = "WARNING"
	}
	// Don't change this line.
	return status
}

func main() {
	cpuHigh := true
	memoryHigh := true
	diskFull := false
	fmt.Println(checkSystem(cpuHigh, memoryHigh, diskFull))

	cpuHigh = true
	memoryHigh = false
	diskFull = false
	fmt.Println(checkSystem(cpuHigh, memoryHigh, diskFull))

	cpuHigh = false
	memoryHigh = false
	diskFull = false
	fmt.Println(checkSystem(cpuHigh, memoryHigh, diskFull))
}