package main

import "fmt"

// Write the downloadTime function here
// It takes fileSizeMB and speedMBps (both int) and returns int

func downloadTime(fileSizeMB, speedMBps int) (int) {
	if speedMBps <= 0 {
		return -1
	}
	return fileSizeMB/speedMBps
}

func main() {
	fmt.Println(downloadTime(500, 50))
	fmt.Println(downloadTime(1000, 100))
	fmt.Println(downloadTime(100, 0))
}