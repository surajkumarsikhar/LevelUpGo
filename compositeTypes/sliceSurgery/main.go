package main

import "fmt"

func RecentDeploys(tags []string, n int) []string {
	// Your code here
	if n > len(tags) {
		return tags
	} else if n <= 0 {
		return []string{}
	}
	i := len(tags) - n
	return tags[i:]
}

func FirstDeploys(tags []string, n int) []string {
	// Your code here
	if n > len(tags) {
		return tags
	} else if n <= 0 {
		return []string{}
	}
	return tags[:n]
}

func main() {
	tags := []string{"v1.0.0", "v1.1.0", "v1.2.0", "v2.0.0", "v2.1.0"}
	fmt.Println(RecentDeploys(tags, 2))
	fmt.Println(FirstDeploys(tags, 3))
}
