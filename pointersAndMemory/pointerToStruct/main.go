package main

import "fmt"

type VideoJob struct {
	Filename   string
	Resolution int
}

// NewVideoJob builds a *VideoJob from a filename and target resolution.
// Write the whole function: signature and body.
func NewVideoJob(filename string, resolution int) *VideoJob {
	v := VideoJob{
		Filename:   filename,
		Resolution: resolution,
	}

	return &v
}

func main() {
	j := NewVideoJob("intro.mp4", 1080)
	fmt.Println(j.Filename, j.Resolution)
}
