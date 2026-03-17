package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to function in golang")

	// Calling a simple function with one parameter
	greater("Gouranga")

	// Calling a function that adds two numbers
	result := adder(2, 3)
	fmt.Println("Result is", result)

	// Calling a variadic function (accepts multiple integers)
	proResult := proAdder(2, 51, 32, 3, 56, 4, 1, 4, 2, 4)
	fmt.Println("Pro result is", proResult)
}

// greater prints a greeting message with the given name
func greater(name string) {
	fmt.Printf("Hello %v, welcome to our app.\n", name)
}

// adder takes two integers and returns their sum
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// proAdder is a variadic function
// It can accept any number of integer arguments
func proAdder(values ...int) int {
	total := 0

	// Loop through all provided values and accumulate the sum
	for _, v := range values {
		total += v
	}

	return total
}
