package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to conditions section")

	// Number of times a user has logged in
	loginCount := 23

	// Variable to store user category
	var result string

	// Conditional logic to determine user type
	if loginCount > 10 {
		result = "Regular user"
	} else if loginCount > 20 {
		result = "Pro user"
	} else {
		result = "Irregular user"
	}

	// Print the result
	fmt.Println("User status: ", result)

	// Short variable declaration inside if statement
	// 'num' exists only within this if-else block
	if num := 3; num%2 == 0 {
		fmt.Printf("%v is even", num)
	} else {
		fmt.Printf("%v is odd", num)
	}
}
