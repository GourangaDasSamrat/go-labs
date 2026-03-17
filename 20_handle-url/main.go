package main

import (
	"fmt"
	"net/url"
)

// myUrl holds the URL string we want to parse and analyze
const myUrl = "https://www.google.com/search?client=firefox-b-d&q=gouranga+das+samrat"

func main() {
	// Parse the raw URL string into a structured URL object
	result, err := url.Parse(myUrl)
	checkNilErr(err)

	// You can inspect different parts of the parsed URL by uncommenting below

	// fmt.Println("Scheme =>", result.Scheme)        // Protocol (https)
	// fmt.Println("Host =>", result.Host)            // Domain (www.google.com)
	// fmt.Println("Query =>", result.Query())        // Query parameters as map
	// fmt.Println("Raw query =>", result.RawQuery)   // Full query string
	// fmt.Println("Path =>", result.Path)            // Endpoint (/search)
	// fmt.Println("Request URL =>", result.RequestURI())

	// Extract query parameters into a map[string][]string
	qParams := result.Query()

	// Accessing specific query parameters (if needed)
	// fmt.Println(qParams["q"])       // Search keyword
	// fmt.Println(qParams["client"])  // Client identifier

	// Loop through all query parameter values and print them
	for _, v := range qParams {
		fmt.Println("Param is =>", v)
	}

	// Manually constructing a URL using the url.URL struct
	partsOfUrl := &url.URL{
		Scheme:   "https",                                    // Protocol
		Host:     "www.google.com",                           // Domain
		Path:     "/search",                                  // Path
		RawQuery: "client=firefox-b-d&q=gouranga+das+samrat", // Query string
	}

	// Convert the structured URL back into a string
	anotherUrl := partsOfUrl.String()

	// Print the reconstructed URL
	fmt.Println(anotherUrl)
}

// checkNilErr is a helper function for quick error handling
// It panics if an error is encountered
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
