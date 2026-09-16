package contextpackage
package main
import (
	"fmt"
	"io"
	"net/http"
	"context"
)
// baseURL is the API base URL (can be overridden for testing)
var baseURL = "http://bookstore-api"

// TODO: Implement fetchBook
//
// func fetchBook(ctx context.Context, id int) (string, error)
//
// Steps:
// 1. Build URL: fmt.Sprintf("%s/books/%d", baseURL, id)
// 2. Create request: http.NewRequestWithContext(ctx, "GET", url, nil)
// 3. Execute: http.DefaultClient.Do(req)
// 4. Read body: io.ReadAll(resp.Body)
// 5. Return string(body), nil
//
// Don't forget to import: "context", "fmt", "io", "net/http"

func fetchBook(ctx context.Context, id int) (string, error) {
	url := fmt.Sprintf("%s/books/%d",baseURL,id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "",err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "",err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	res := string(body)

	return res, nil

}

func main() {
}