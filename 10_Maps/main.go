package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps section in golang")

	// Declare a map where both key and value are strings
	languages := make(map[string]string)

	// Add key-value pairs to the map
	languages["Js"] = "JavaScript"
	languages["Ts"] = "TypeScript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"

	// Print the entire map
	fmt.Println("List of all languages:", languages)

	// Access a specific value using its key
	fmt.Println("JS stands for", languages["Js"])

	// Delete a key-value pair from the map
	delete(languages, "Rb")

	// Print the map after deletion
	fmt.Println("List of all languages after deletion:", languages)

	// Loop through the map using range
	// Note: Map iteration order in Go is random
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
