package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to conditions section")

	loginCount := 23
	var result string

	if loginCount > 10 {
		result = "Regular user"
	} else if loginCount > 20 {
		result = "Pro user"
	} else {
		result = "Irregular user"
	}

	fmt.Println("User status: ", result)

	// another syntax
	if num := 3; num%2 == 0 {
		fmt.Printf("%v is even", num)
	} else {
		fmt.Printf("%v is odd", num)
	}
}
