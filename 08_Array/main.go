package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome in array in golang")

	wtaRank := [5]string{"Aryna Sabalenka", "Elena Rybakina", "Jessica Pegula", "Coco Gauff", "Elina Svitolina"}

	fmt.Println(wtaRank)
	fmt.Println("My favorite tennis player: ", wtaRank[1])
	fmt.Println(len(wtaRank))

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

	fmt.Println(fideRank)
	fmt.Println("My favorite chess player: ", fideRank[0])
	fmt.Println(len(fideRank))
}
