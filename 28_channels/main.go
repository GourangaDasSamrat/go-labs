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
/*
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}
*/

// synchronization
/*
func task(done chan bool) {
	defer func() { done <- true }()

	fmt.Println("Processing task...")
}
*/

// email sender worker
/*
func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("Sending mail to", email)
		time.Sleep(5 * time.Nanosecond)
	}
}
*/

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "Hello"
	}()

	for range 2 {
		select {
		case chan1Val := <-chan1:
			fmt.Println(chan1Val)

		case chan2Val := <-chan2:
			fmt.Println(chan2Val)
		}
	}

	/*
		// buffered channel
		emailChan := make(chan string, 100)
		done := make(chan bool) // synchronizer channel

		go emailSender(emailChan, done)

		for range 100 {
			local := make([]byte, 5)
			for j := range local {
				local[j] = byte(rand.Intn(26) + 'a')
			}
			email := string(local) + "@gmail.com"
			emailChan <- email
		}
		close(emailChan)

		<-done // block

	*/

	/*
		done := make(chan bool)

		go task(done)

		<-done
	*/

	/*
		result := make(chan int)

		go sum(result, randNum, randNum)

		res := <-result

		fmt.Println(res)
	*/

	/*
		numChan := make(chan int)

		go processNum(numChan)

		for {
			numChan <- randNum
		}
	*/
}
