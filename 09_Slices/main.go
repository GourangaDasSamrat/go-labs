package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Golang")

	var fideRank = []string{
		"Magnus Carlsen",
		"Hikaru Nakamura",
		"Fabiano Caruana",
		"Vincent Keymer",
		"Anish Giri",
	}

	// add values
	fideRank = append(fideRank, "Gukesh D", "Arjuna Erigaisi")

	fmt.Println(fideRank)
	fmt.Println(len(fideRank))

	fideRank = append(fideRank[3:5])

	fmt.Println(fideRank)

	wtaRank := make([]string, 4)
	wtaRank[0] = "Elena Rybakina"
	wtaRank[1] = "Aryna Sabalenka"
	wtaRank[2] = "Jessica Pegula"
	wtaRank[3] = "Coco Gauff"
	// wtaRank[4] = "Coco Gauff" // it give error

	wtaRank = append(wtaRank, "Elina Svitolina") // but it don't

	fmt.Println(wtaRank)

	scores := []int{93, 10, 12, 34, 2, 89, 11}

	sort.Ints(scores)

	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	// remove value from slices based on index
	var aptRank = []string{"Carlos Alcaraz", "Jannik Sinner", "Novak Djokovic", "Alexander Zverev", "Lorenzo Musetti"}
	fmt.Println(aptRank)

	var index int = 2

	aptRank = append(aptRank[:index],aptRank[index+1:]...)
	fmt.Println(aptRank)
}
