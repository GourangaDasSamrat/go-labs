package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to my app")

	// Get the current local time
	presentTime := time.Now()
	fmt.Println("Present time is:", presentTime)

	// Format the current time in a custom layout
	fmt.Println("Present time formatted:", presentTime.Format("01-02-2006 Monday 15:04:05"))

	// Load a specific time zone (Asia/Dhaka)
	loc, err := time.LoadLocation("Asia/Dhaka")
	if err != nil {
		panic(err) // stop execution if timezone loading fails
	}

	// Create a specific date and time in the given location
	createdDate := time.Date(
		2008,         // Year
		time.October, // Month
		5,            // Day
		21,           // Hour (24-hour format)
		0,            // Minute
		0,            // Second
		0,            // Nanosecond
		loc,          // Timezone
	)

	// Print the created date in default format
	fmt.Println("My birth time is:", createdDate)

	// Print the created date in custom formats
	fmt.Println("My birth time (24h format with timezone):", createdDate.Format("2006-01-02 15:04:05 MST"))
	fmt.Println("My birth time (12h format with AM/PM and timezone):", createdDate.Format("2006-01-02 03:04:05 PM MST"))
}
