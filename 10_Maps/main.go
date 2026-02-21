package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps section in golang")

	// declare map where key and value both are string
	languages := make(map[string]string)

	// add values on map
	languages["Js"] = "JavaScript"
	languages["Ts"] = "TypeScript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"

	// print the map
	fmt.Println("List of all languages: ", languages)

	// print a specific value from map
	fmt.Println("JS stands for ", languages["Js"])

	/// delete a value from map
	delete(languages, "Rb")

	// print the map
	fmt.Println("List of all languages: ", languages)

	// loop on map
	for key, value := range languages {
		fmt.Printf("For key %v value is %v\n", key, value)
	}
}
