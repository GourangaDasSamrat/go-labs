package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Logical CPUs:", runtime.NumCPU())
}
