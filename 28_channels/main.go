package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {

	for num := range numChan {
		fmt.Println("Processing number", num)
		time.Sleep(5 * time.Nanosecond)
	}
}

func main() {
	numChan := make(chan int)

	go processNum(numChan)

	for {
		numChan <- rand.Intn(100)
	}
}
