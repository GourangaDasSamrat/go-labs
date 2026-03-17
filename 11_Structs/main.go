package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to structs in Golang")

	// Note: Go does not have inheritance, super, or parent classes like other OOP languages

	// Creating an instance of User struct
	user1 := User{
		"Gouranga Das Samrat",       // Name
		"gouranga.samrat@gmail.com", // Email
		true,                        // IsLoggedIn
		17,                          // Age
	}

	// Print the entire struct (default formatting)
	fmt.Println("User 1:", user1)

	// Print the struct with field names
	fmt.Printf("User 1 details are %+v\n", user1)

	// Access individual fields of the struct
	fmt.Println("User 1 name:", user1.Name)
	fmt.Println("User 1 email:", user1.Email)
	fmt.Println("User 1 age:", user1.Age)
	fmt.Println("User 1 login status:", user1.IsLoggedIn)
}

// User struct defines basic information about a user
type User struct {
	Name       string // User's full name
	Email      string // User's email address
	IsLoggedIn bool   // Login status
	Age        int    // User's age
}
