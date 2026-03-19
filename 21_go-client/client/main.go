package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "http://localhost:8000"

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
	}
}

// run coordinates the application logic
func run() error {
	must(handleGetRequest(), "handle GET request")
	return nil
}

// handleGetRequest sends a GET request and prints the response
func handleGetRequest() error {
	url := baseURL + "/get"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	must(err, "create request")

	res, err := httpClient.Do(req)
	must(err, "send request")
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	must(err, "read response")

	fmt.Println("Status code:", res.StatusCode)
	fmt.Println("Content length:", res.ContentLength)
	fmt.Println(string(body))

	return nil
}


// must checks if err is not nil and wraps it with a message
// It panics immediately if there is an error
func must(err error, message string) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", message, err))
	}
}
