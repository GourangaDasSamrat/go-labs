package main

import (
	"fmt"
	"io"

	"os"
)

func main() {
	fmt.Println("Welcome to file handling in golang")

	content := `Between Rita and my eyes, there is a rifle. And whoever knows Rita, kneels and prays to the divinity in those honey-colored eyes. I remember Rita the way a sparrow remembers its stream. Rita's name was a feast in my mouth, her body was a wedding in my blood, and I was lost in Rita for two years.  -- by Mahmoud Darwish`

	// create file
	file, err := os.Create("./msg.txt")

	// check for error
	checkNilErr(err)

	// write into the file
	len, err := io.WriteString(file, content)

	// check for error
	checkNilErr(err)

	// print length
	fmt.Println("Length of message is", len)

	// close file package
	defer file.Close()

	// invoke read file fuc
	readFile("./msg.txt")
}

func readFile(fileName string) {
	// read the file
	data, err := os.ReadFile(fileName)

	// check for error
	checkNilErr(err)

	// print the data inside the file
	fmt.Println("Message is:\n", string(data))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
