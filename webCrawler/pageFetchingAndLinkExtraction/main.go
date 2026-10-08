package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const baseURL = "http://gadgetshop"

// FetchResult holds the result of fetching a URL
type FetchResult struct {
	URL        string
	Body       string
	StatusCode int
}

// fetchPage makes an HTTP GET request and returns the result.
// Returns a FetchResult and nil error on success.
// Returns a FetchResult and an error if the request or body read fails.
func fetchPage(url string) (FetchResult, error) {
	// TODO: Implement
	resp, err := http.Get(url)
	if err != nil {
		return FetchResult{
			URL:        url,
			Body:       "",
			StatusCode: resp.StatusCode,
		}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return FetchResult{
			URL:        url,
			Body:       "",
			StatusCode: resp.StatusCode,
		}, err
	}
	return FetchResult{
		URL:        url,
		Body:       string(body),
		StatusCode: resp.StatusCode,
	}, nil
}

// extractLinks finds all href="..." values in an HTML body.
// Returns a slice of the raw href values.
func extractLinks(body string) []string {
	// TODO: Implement
	res := []string{}
	for {
		idx := strings.Index(body, `href="`)
		n := len(`href="`)
		if idx == -1 {
			break
		}

		start := idx + n
		var end = start

		for end < len(body) && body[end] != '"' {
			end++
		}
		res = append(res, body[start:end])
		body = body[end+1:]
	}
	return res
}

// resolveURL converts a relative link to an absolute URL using the base URL.
// If the link is already absolute (starts with http:// or https://), return it as-is.
// If the link starts with /, prepend the scheme + host from the base URL.
func resolveURL(base, link string) string {
	// TODO: Implement
	if strings.Contains(link, `https://`) || strings.Contains(link, `http://`) {
		return link
	}
	parts := strings.Split(base, "/")
	newBase := strings.Join(parts[:3], "/")

	return newBase + link
}

// fetchAndExtract fetches a URL and returns the result plus resolved links.
// Returns an error if the fetch fails or if the status code is not 200.
func fetchAndExtract(url string) (FetchResult, []string, error) {
	// TODO: Implement
	res, err := fetchPage(url)
	if err != nil {
		return FetchResult{}, []string{}, err
	} else if res.StatusCode != 200 {
		return FetchResult{}, []string{}, fmt.Errorf("404 not found")
	}
	links := extractLinks(res.Body)

	for i := range links {
		links[i] = resolveURL(res.URL, links[i])
	}
	return res, links, nil
}

func main() {
	// Test fetchPage
	result, err := fetchPage(baseURL + "/shop/")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Fetched %s (status %d, %d bytes)\n", result.URL, result.StatusCode, len(result.Body))
	}

	// Test 404
	result, err = fetchPage(baseURL + "/shop/nonexistent")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Missing page: status=%d\n", result.StatusCode)
	}

	// Test extractLinks
	links := extractLinks(`<a href="/shop/products">Products</a> <a href="/shop/about">About</a>`)
	fmt.Printf("Extracted links: %v\n", links)

	// Test resolveURL
	fmt.Printf("Resolve relative: %s\n", resolveURL("http://gadgetshop/shop/products/1", "/shop/products"))
	fmt.Printf("Resolve absolute: %s\n", resolveURL("http://gadgetshop/shop/", "http://other.com/page"))

	// Test fetchAndExtract
	result, resolved, err := fetchAndExtract(baseURL + "/shop/")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("Home page has %d links:\n", len(resolved))
		for _, link := range resolved {
			fmt.Printf("  -> %s\n", link)
		}
	}

	// Test fetchAndExtract on about page
	result, resolved, err = fetchAndExtract(baseURL + "/shop/about")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("About page has %d links\n", len(resolved))
	}

	// Test fetchAndExtract on non-existent page
	_, _, err = fetchAndExtract(baseURL + "/shop/nonexistent")
	if err != nil {
		fmt.Printf("Fetch error: %v\n", err)
	}
}
