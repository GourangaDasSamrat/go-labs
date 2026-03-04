package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in golang")

	// coffees := []string{"Latte", "Cappuccino", "Americano", "Macchiato", "Cortado"}

	// fmt.Println(coffees)

	// for c := 0; c < len(coffees); c++ {
	// 	fmt.Println(coffees[c])
	// }

	// for i := range coffees {
	// 	fmt.Println(coffees[i])
	// }

	// for i, coffee := range coffees {
	// 	fmt.Printf("Index is %v and value is %v\n", i, coffee)
	// }

	// for _, coffee := range coffees {
	// 	fmt.Println(coffee)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto mailMe
		}

		// if rogueValue == 5 {
		// 	break
		// }

		// if rogueValue == 5 {
		// 	rogueValue++
		// 	continue
		// }

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

mailMe:
	fmt.Println("Go to gouranga.samrat@gmail.com")
}
