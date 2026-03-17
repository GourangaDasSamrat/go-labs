package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to file handling in golang")

	// Multi-line string content to be written into the file
	content := `Between Rita and my eyes, there is a rifle. And whoever knows Rita, kneels and prays to the divinity in those honey-colored eyes. I remember Rita the way a sparrow remembers its stream. Rita's name was a feast in my mouth, her body was a wedding in my blood, and I was lost in Rita for two years.  -- by Mahmoud Darwish`

	// Step 1: Create a new file (or overwrite if it already exists)
	file, err := os.Create("./msg.txt")

	// Check for any error during file creation
	checkNilErr(err)

	// Step 2: Write the content string into the file
	// 'len' stores the number of bytes written
	len, err := io.WriteString(file, content)

	// Check for any error during writing
	checkNilErr(err)

	// Print the length of the written content
	fmt.Println("Length of message is", len)

	// Ensure the file is closed after all operations are done
	defer file.Close()

	// Step 3: Read and display the file content
	readFile("./msg.txt")
}

// readFile reads the content of a file and prints it to the console
func readFile(fileName string) {
	// Read the entire file content into memory
	data, err := os.ReadFile(fileName)

	// Check for any error during reading
	checkNilErr(err)

	// Convert byte data to string and print it
	fmt.Println("Message is:\n", string(data))
}

// checkNilErr is a helper function for error handling
// It stops execution if an error is encountered
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
