package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const baseURL = "http://localhost:8000"

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	// Run both requests separately
	if err := handleGetRequest(ctx); err != nil {
		return err
	}

	if err := handlePostRequest(ctx); err != nil {
		return err
	}

	return nil
}

// check is your repetitive error handler (like must, but safer)
func check(err error, msg string) error {
	if err != nil {
		return fmt.Errorf("%s: %w", msg, err)
	}
	return nil
}

func handleGetRequest(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"/get", nil)
	if err := check(err, "create get"); err != nil {
		return err
	}

	res, err := httpClient.Do(req)
	if err := check(err, "do get"); err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err := check(err, "read get"); err != nil {
		return err
	}

	printJSON("GET Success", res.StatusCode, body)
	return nil
}

func handlePostRequest(ctx context.Context) error {
	payload := map[string]any{
		"name":     "Gouranga Das Samrat",
		"age":      17,
		"company":  "Wiz Ecosystem",
		"position": "CEO",
	}

	data, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+"/post", bytes.NewReader(data))
	if err := check(err, "create post"); err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err := check(err, "do post"); err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err := check(err, "read post"); err != nil {
		return err
	}

	printJSON("POST Success", res.StatusCode, body)
	return nil
}

// printJSON outputs a clean JSON line without backslashes
func printJSON(msg string, status int, body []byte) {
	output := map[string]any{
		"time":   time.Now().Format(time.RFC3339),
		"msg":    msg,
		"status": status,
		"body":   json.RawMessage(body), // This removes the backslashes
	}

	// Use encoder to write directly to stdout for speed and cleanliness
	enc := json.NewEncoder(os.Stdout)
	_ = enc.Encode(output)
}
