package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Golang")

	// Slice of top FIDE chess players
	fideRank := []string{
		"Magnus Carlsen",
		"Hikaru Nakamura",
		"Fabiano Caruana",
		"Vincent Keymer",
		"Anish Giri",
	}

	// Append more values to the slice
	fideRank = append(fideRank, "Gukesh D", "Arjuna Erigaisi")

	fmt.Println("Updated FIDE Rank:", fideRank)
	fmt.Println("Length of slice:", len(fideRank))

	// Slice from index 3 to 5 (5 is excluded)
	fideRank = fideRank[3:5]
	fmt.Println("Sliced FIDE Rank (3:5):", fideRank)

	// Creating a slice with predefined length
	wtaRank := make([]string, 4)
	wtaRank[0] = "Elena Rybakina"
	wtaRank[1] = "Aryna Sabalenka"
	wtaRank[2] = "Jessica Pegula"
	wtaRank[3] = "Coco Gauff"
	// wtaRank[4] = "Elina Svitolina" // ❌ Error: index out of range

	// Use append to add elements beyond initial length
	wtaRank = append(wtaRank, "Elina Svitolina")
	fmt.Println("WTA Rank:", wtaRank)

	// Sorting integer slice
	scores := []int{93, 10, 12, 34, 2, 89, 11}
	sort.Ints(scores)
	fmt.Println("Sorted scores:", scores)
	fmt.Println("Is sorted?", sort.IntsAreSorted(scores))

	// Removing an element from slice by index
	aptRank := []string{"Carlos Alcaraz", "Jannik Sinner", "Novak Djokovic", "Alexander Zverev", "Lorenzo Musetti"}
	fmt.Println("Original ATP Rank:", aptRank)

	index := 2 // remove "Novak Djokovic"
	aptRank = append(aptRank[:index], aptRank[index+1:]...)
	fmt.Println("Updated ATP Rank:", aptRank)
}
