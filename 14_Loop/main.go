package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in golang")

	// Example slice (currently commented out)
	// coffees := []string{"Latte", "Cappuccino", "Americano", "Macchiato", "Cortado"}

	// Different ways to loop over a slice:

	// Classic for loop using index
	// for c := 0; c < len(coffees); c++ {
	// 	fmt.Println(coffees[c])
	// }

	// Using range to get only index
	// for i := range coffees {
	// 	fmt.Println(coffees[i])
	// }

	// Using range to get both index and value
	// for i, coffee := range coffees {
	// 	fmt.Printf("Index is %v and value is %v\n", i, coffee)
	// }

	// Ignoring index using underscore (_)
	// for _, coffee := range coffees {
	// 	fmt.Println(coffee)
	// }

	// Demonstrating loop with condition (like a while loop)
	rogueValue := 1

	for rogueValue < 10 {

		// Using goto to jump to a label
		if rogueValue == 2 {
			goto mailMe
		}

		// Example of breaking the loop (currently commented)
		// if rogueValue == 5 {
		// 	break
		// }

		// Example of skipping an iteration (currently commented)
		// if rogueValue == 5 {
		// 	rogueValue++
		// 	continue
		// }

		fmt.Println("Value is: ", rogueValue)

		// Increment the loop variable
		rogueValue++
	}

	// Label used by goto statement
mailMe:
	fmt.Println("Go to gouranga.samrat@gmail.com")
}
