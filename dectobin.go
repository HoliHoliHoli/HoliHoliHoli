package main

import (
	"fmt"
)

func main() {
	fmt.Println("This program coverts decimal numbers to binary form")

	var number int
	fmt.Println("input a number")
	fmt.Scan(&number)

	fmt.Printf("You inputed %d which has a binary representation of %b", number, number)
}
