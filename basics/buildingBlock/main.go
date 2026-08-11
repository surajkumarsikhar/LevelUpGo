package main

import "fmt"

func main() {
	statusCode := 200
	latency := 3.0/2.0
	/* use + to join with a string */
	fmt.Println("Status: " + "OK")
	/* use == with statusCode and 200 */
	fmt.Println("Healthy:", statusCode==200)
	/* divide two float64 values to get 1.5 */
	fmt.Println("Latency:",latency)
}