package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to my app")

	presentTime := time.Now()

	fmt.Println("Present time is: ", presentTime)

	fmt.Println("Present time is: ", presentTime.Format("01-02-2006 Monday 15:04:05"))
}
