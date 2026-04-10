package main

import (
	"fmt"
)

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
	fmt.Println()
}

type Stack[T int | string] struct {
	items []T
}

func main() {
	names := []string{"Iga Swiatek", "Aryna Sabalenka", "Ons Jabeur", "Jessica Pegula", "Coco Gauff"}
	nums := []int{1, 2, 3, 4, 5}

	printSlice(names)
	printSlice(nums)

	stack := Stack[string]{items: names}
	stack2 := Stack[int]{items: nums}

	fmt.Println("Stack items:", stack.items)
	fmt.Println("2nd Stack items:", stack2.items)
}
