package main

import (
	"fmt"
)

func main() {

	for {
		var choice string
		var a int
		var b int

		fmt.Println("Input Choices:")
		fmt.Println("add a b")
		fmt.Println("sub a b")
		fmt.Println("mul a b")
		fmt.Println("div a b")
		fmt.Println("Type 'help' or 'quit'")
		fmt.Print("> ")

		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Error: invalid input")
			continue
		}

		if choice == "quit" {
			fmt.Println("Enjoy your day")
			return
		}

		if choice == "help" {
			fmt.Println("Commands:")
			fmt.Println("add a b to add use this= add 5 6 ")
			fmt.Println("sub a b to subtract use = sub 9 2")
			fmt.Println("mul a b to multiply use = mul 7 2")
			fmt.Println("div a b to divide use = div 8 2")
			fmt.Println(" you will see your result deisplayed below the calculation")
			continue
		}

		_, err = fmt.Scan(&a, &b)
		if err != nil {
			fmt.Println("Error: enter valid numbers")
			continue
		}

		if choice == "add" {
			fmt.Println("✦ Result:", a+b)

		} else if choice == "sub" {
			fmt.Println("✦ Result:", a-b)

		} else if choice == "mul" {
			fmt.Println("✦ Result:", a*b)

		} else if choice == "div" {
			if b == 0 {
				fmt.Println("Error: cannot divide by zero")
				continue
			}
			fmt.Println("✦ Result:", a/b)

		} else {
			fmt.Println("Unknown command. Type 'help'")
		}
	}
}
