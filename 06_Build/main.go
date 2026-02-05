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

	loc, err := time.LoadLocation("Asia/Dhaka")
	if err != nil {
		panic(err)
	}

	createdDate := time.Date(
		2008,
		time.October,
		5,
		21,
		00,
		0,
		0,
		loc,
	)

	fmt.Println("My birth time is: ", createdDate)
	fmt.Println(
		"My birth time is:",
		createdDate.Format("2006-01-02 15:04:05 MST"),
	)
	fmt.Println(
		"My birth time is:",
		createdDate.Format("2006-01-02 03:04:05 PM MST"),
	)
}
