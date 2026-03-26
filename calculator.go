package main

import (
	"fmt"
)

func main() {
	var choice string

	fmt.Println("Choose operation:")
	fmt.Println("add")
	fmt.Println("sub")
	fmt.Println("mul")
	fmt.Println("div")

	fmt.Print("> ")
	fmt.Scan(&choice)

	var a int
	var b int

	fmt.Print("> ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	if choice == "+" {
		fmt.Println(a + b)

	}
	if choice == "-" {
		fmt.Println(a - b)

	}

	if choice == "*" {
		fmt.Println(a * b)

	}

	if choice == "/" {
		fmt.Println(a / b)

	}

	fmt.Println("Choose operation:")

	fmt.Println("You chose:", choice)

}
