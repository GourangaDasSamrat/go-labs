package main

import (
	"fmt"
)

func main() {
	// Deferred statements are executed in LIFO order (Last In, First Out)

	defer fmt.Println("I'm defer 1") // Will execute last
	defer fmt.Println("I'm defer 2")
	defer fmt.Println("I'm defer 3") // Will execute first among these three

	fmt.Println("Defer in golang")

	// Call function to demonstrate defer inside loops
	myDefer()
}

func myDefer() {
	// Loop from 0 to 4 (range over integer literal)
	for i := range 5 {
		// Each iteration registers a deferred call
		// These will execute after the function completes, in reverse order
		defer fmt.Println(i)
	}

	// At this point, nothing has been printed from the loop yet
	// All deferred calls will run after this function returns
}
