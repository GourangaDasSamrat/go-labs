package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Doing task", id)
}

func main() {
	for i := 1; i <= 10; i++ {
		go task(i)
	}

	time.Sleep(time.Second * 1)
}
