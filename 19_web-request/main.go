package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// url points to the documentation site you want to download
// This example fetches your lightweight JS date library docs
const url = "https://gourangadassamrat.github.io/date-wiz-docs"

func main() {
	fmt.Println("Downloading docs and starting server...")

	// Step 1: Send an HTTP GET request to the target URL
	res, err := http.Get(url)
	checkNilErr(err)
	defer res.Body.Close() // Ensure the response body is closed after reading

	// Step 2: Read the entire response body into memory
	data, err := io.ReadAll(res.Body)
	checkNilErr(err)

	// Step 3: Create (or overwrite) a local file named index.html
	file, err := os.Create("./index.html")
	checkNilErr(err)
	defer file.Close() // Ensure file is properly closed

	// Step 4: Write the downloaded data into the file
	// We ignore the number of bytes written using _
	_, err = file.Write(data)
	checkNilErr(err)

	fmt.Println("Successfully saved to index.html")
	fmt.Println("Serving folder at http://localhost:8080")

	// Step 5: Serve the current directory over HTTP
	// This allows you to open the downloaded file in a browser
	fs := http.FileServer(http.Dir("."))

	// Start a local web server on port 8080
	err = http.ListenAndServe(":8080", fs)
	checkNilErr(err)
}

// checkNilErr is a simple helper for error handling
// It stops execution immediately if an error occurs
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
