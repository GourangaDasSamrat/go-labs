package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://gourangadassamrat.github.io/date-wiz-docs" //this is the docs site of my super lightweight date library is js, in case you use js and work with date, must be checkout this library

func main() {
	fmt.Println("Downloading docs and starting server...")

	// 1. Perform the GET request
	res, err := http.Get(url)
	checkNilErr(err)
	defer res.Body.Close()

	// 2. Read the body data
	data, err := io.ReadAll(res.Body)
	checkNilErr(err)

	// 3. Create/Overwrite index.html
	file, err := os.Create("./index.html")
	checkNilErr(err)
	defer file.Close()

	// 4. Write data to file (using _ to ignore the returned length)
	_, err = file.Write(data)
	checkNilErr(err)

	fmt.Println("Successfully saved to index.html")
	fmt.Println("Serving folder at http://localhost:8080")

	// 5. Serve the current directory
	fs := http.FileServer(http.Dir("."))
	err = http.ListenAndServe(":8080", fs)
	checkNilErr(err)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
