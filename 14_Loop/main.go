package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in golang")

	coffees := []string{"Latte", "Cappuccino", "Americano", "Macchiato", "Cortado"}

	fmt.Println(coffees)

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
}
