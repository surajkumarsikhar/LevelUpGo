package main

import "fmt"

func NewShortener() map[string]string {
	// Your code here
	lookUpTable := make(map[string]string)
	return lookUpTable
}

func Save(store map[string]string, slug, url string) {
	// Your code here
	store[slug] = url
}

func Resolve(store map[string]string, slug string) string {
	// Your code here
	url, _ := store[slug]
	return url
}

func main() {
	s := NewShortener()
	Save(s, "gh", "https://github.com")
	Save(s, "docs", "https://go.dev/doc")
	fmt.Println(Resolve(s, "gh"))
	fmt.Println(Resolve(s, "docs"))
}
