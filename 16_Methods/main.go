package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to structs")

	user1 := User{
		"Gouranga Das Samrat",
		"gouranga.samrat@gmail.com",
		true,
		17,
		"his",
	}

	user1.GetStatus()
	user1.NewMail()
}

type Pronoun string

const (
	Male   Pronoun = "his"
	Female Pronoun = "her"
)

type User struct {
	Name       string
	Email      string
	IsLoggedIn bool
	Age        int
	Pronoun    Pronoun
}

func (u User) GetStatus() {
	fmt.Printf("User's name is %v, user's mail is %v, user's age is %v and %v login status is %v.\n",
		u.Name, u.Email, u.Age, u.Pronoun, u.IsLoggedIn)
}

func (u User) NewMail() {
	u.Email = "me@go.dev"
	fmt.Printf("Email of %v is %v", u.Name, u.Email)
}
