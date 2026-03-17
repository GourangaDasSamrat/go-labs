package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to arrays in golang")

	// Array of top 5 WTA tennis players
	wtaRank := [5]string{
		"Aryna Sabalenka",
		"Elena Rybakina",
		"Jessica Pegula",
		"Coco Gauff",
		"Elina Svitolina",
	}

	// Print the entire array
	fmt.Println("WTA Rankings:", wtaRank)

	// Access a specific element using its index (0-based)
	fmt.Println("My favorite tennis player:", wtaRank[1])

	// Length of the array
	fmt.Println("Total WTA players in array:", len(wtaRank))

	// Array of top 10 FIDE chess players
	var fideRank = [10]string{
		"Magnus Carlsen",
		"Hikaru Nakamura",
		"Fabiano Caruana",
		"Vincent Keymer",
		"Arjuna Erigaisi",
		"Anish Giri",
		"Alireza Firouzja",
		"R Praggnanandhaa",
		"Wei Yi",
		"Gukesh D",
	}

	// Print the chess ranking array
	fmt.Println("FIDE Rankings:", fideRank)

	// Access a specific element
	fmt.Println("My favorite chess player:", fideRank[0])

	// Length of the array
	fmt.Println("Total FIDE players in array:", len(fideRank))
}
