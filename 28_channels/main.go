package main

import (
	"fmt"
	"math/rand"
)

var randNum int = rand.Intn(100)

// send
/*
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("Processing number", num)
		time.Sleep(5 * time.Nanosecond)
	}
}
*/

// receive
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func main() {
	result := make(chan int)

	go sum(result, randNum, randNum)

	res := <-result

	fmt.Println(res)

	/*
		numChan := make(chan int)

		go processNum(numChan)

		for {
			numChan <- randNum
		}
	*/
}
