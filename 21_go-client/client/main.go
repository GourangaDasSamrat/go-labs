package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// baseURL defines the root server address used across requests
const baseURL = "http://localhost:8000"

// httpClient is a reusable client with timeout for all HTTP calls
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

// main is the entry point and delegates execution to run
func main() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
	}
}

// run controls the application flow and stops on first failure
func run() error {
	if err := handleGetRequest(); err != nil {
		return err
	}

	return nil
}

// handleGetRequest sends a GET request and prints response details
func handleGetRequest() error {
	url := baseURL + "/get"

	// create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	// execute the request using shared client
	res, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer res.Body.Close()

	// read response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	// print response information
	fmt.Println("Status code:", res.StatusCode)
	fmt.Println("Content length:", res.ContentLength)
	fmt.Println(string(body))

	return nil
}
