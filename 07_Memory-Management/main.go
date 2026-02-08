package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Logical CPUs:", runtime.NumCPU())

	var userName *string
	fmt.Println("Default value of pointer is ", userName)

	myAge := 17

	var ptr = &myAge
	fmt.Println("Memory address of pointer is: ", ptr)
	fmt.Println("Actual value of pointer is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("My new age is: ", myAge) // i'm not this much old 😂️
}
