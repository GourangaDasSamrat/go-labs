package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to my app")
	fmt.Println("Please rate our coffee between 1 to 10:")

	// Create a new reader to read input from standard input (console)
	reader := bufio.NewReader(os.Stdin)

	// Read user input until newline character
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating:", input)

	// Convert input string to float64
	// TrimSpace removes newline and any extra spaces
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		// Handle invalid input (non-numeric)
		fmt.Println("Error parsing input:", err)
	} else {
		// Add 1 to the rating and display
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}
