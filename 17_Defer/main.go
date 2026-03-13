package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("I'm defer 1")
	defer fmt.Println("I'm defer 2")
	defer fmt.Println("I'm defer 3")

	fmt.Println("Defer in golang")

	myDefer()
}

func myDefer() {
	for i := range 5 {
		defer fmt.Println(i)
	}
}
