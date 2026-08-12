package main

import "fmt"

// formatDownload returns a string like "backup.zip (3.50 MB) downloaded in 12s"
func formatDownload(filename string, sizeMB float64, seconds int) string {
	// Use fmt.Sprintf to build the formatted string
	msg := fmt.Sprintf("%s (%.2f MB) downloaded in %ds", filename, sizeMB, seconds)
	return msg
}

func main() {
	fmt.Println(formatDownload("backup.zip", 3.5, 12))
	fmt.Println(formatDownload("photo.png", 0.8, 1))
	fmt.Println(formatDownload("video.mp4", 125.0, 45))
}
