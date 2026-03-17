package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to structs")

	// Creating an instance of User struct
	user1 := User{
		"Gouranga Das Samrat",       // Name
		"gouranga.samrat@gmail.com", // Email
		true,                        // IsLoggedIn status
		17,                          // Age
		"his",                       // Pronoun (custom type)
	}

	// Calling methods on the struct
	user1.GetStatus()
	user1.NewMail()
}

// Pronoun is a custom type based on string
type Pronoun string

// Constants for Pronoun type (helps avoid hardcoding strings)
const (
	Male   Pronoun = "his"
	Female Pronoun = "her"
)

// User struct defines a user with basic fields
type User struct {
	Name       string  // User's full name
	Email      string  // User's email address
	IsLoggedIn bool    // Login status
	Age        int     // User's age
	Pronoun    Pronoun // Custom type for pronoun
}

// GetStatus prints the current state of the user
func (u User) GetStatus() {
	fmt.Printf(
		"User's name is %v, user's mail is %v, user's age is %v and %v login status is %v.\n",
		u.Name,
		u.Email,
		u.Age,
		u.Pronoun,
		u.IsLoggedIn,
	)
}

// NewMail tries to update the user's email
// Note: This uses a value receiver, so it does NOT change the original struct
func (u User) NewMail() {
	u.Email = "me@go.dev"
	fmt.Printf("Email of %v is %v", u.Name, u.Email)
}
