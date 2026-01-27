package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to my app"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// comma ok || error

	input, _ := reader.ReadString('\n')
	fmt.Println("Welcome, ", input)
}
