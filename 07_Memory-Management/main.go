package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Print the number of logical CPUs available on the machine
	fmt.Println("Logical CPUs:", runtime.NumCPU())

	// Pointer declaration without initialization defaults to nil
	var userName *string
	fmt.Println("Default value of pointer is", userName)

	// Normal variable
	myAge := 17

	// Pointer to the variable myAge
	var ptr = &myAge
	fmt.Println("Memory address of pointer is:", ptr)

	// Dereferencing the pointer to get actual value
	fmt.Println("Actual value of pointer is:", *ptr)

	// Modify the value of myAge via pointer
	*ptr = *ptr * 2
	fmt.Println("My new age is:", myAge) // I'm not this much old 😂️
}