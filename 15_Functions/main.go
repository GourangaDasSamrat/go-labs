package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to function in golang")

	greater("Gouranga")

	result := adder(2, 3)
	fmt.Println("Result is", result)

	proResult := proAdder(2, 51, 32, 3, 56, 4, 1, 4, 2, 4)
	fmt.Println("Pro result is", proResult)
}

func greater(name string) {
	fmt.Printf("Hello %v, welcome to our app.\n", name)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, v := range values {
		total += v
	}

	return total
}
