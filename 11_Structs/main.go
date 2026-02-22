package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to structs")
	// no inheritance in golang
	// no super or parent in golang

	user1 := User{
		"Gouranga Das Samrat",
		"gouranga.samrat@gmail.com",
		true,
		17,
	}

	fmt.Println("User 1: ",user1)
	fmt.Printf("User 1 details are %+v\n",user1)
	fmt.Println("User 1 name: ",user1.Name)
	fmt.Println("User 1 email: ",user1.Email)
	fmt.Println("User 1 age: ",user1.Age)
	fmt.Println("User 1 login status: ",user1.IsLoggedIn)
}

type User struct {
	Name       string
	Email      string
	IsLoggedIn bool
	Age        int
}
