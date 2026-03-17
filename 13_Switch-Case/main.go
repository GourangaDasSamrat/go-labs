package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case in golang")

	// Generate a seed using current time (ensures different random values each run)
	seed := time.Now().UnixNano()

	// Create a new random generator using the seed
	r := rand.New(rand.NewSource(seed))

	// Generate a random number between 1 and 6 (like a dice roll)
	diceNumber := r.Intn(6) + 1

	// Switch statement to handle different dice outcomes
	switch diceNumber {

	case 1:
		fmt.Println("Dice value is 1 and you can open")

	case 2:
		fmt.Println("You can move 2 spot")

	case 3:
		fmt.Println("You can move 3 spot")

	case 4:
		fmt.Println("You can move 4 spot")

	case 5:
		fmt.Println("You can move 5 spot")

	case 6:
		fmt.Println("You can move 6 spot and roll the dice again")

	default:
		// This should not happen with current logic
		fmt.Println("What was that!")
	}
}
