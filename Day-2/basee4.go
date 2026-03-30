package main

import (
	"fmt"
)

func BaseConverter(word string, base int) int64 {
	var result int64 = 0
	var power int64 = 1

	for i := len(word) - 1; i >= 0; i-- {
		digit := int64(word[i] - '0')
		result += digit * power
		power *= int64(base)
	}
	return result
}

func dex(n int64, base int) string {
	if n == 0 {
		return "0"
	}

	var result string = ""

	for n > 0 {
		remainder := n % int64(base)
		result = string(remainder+'0') + result
		n = n / int64(base)
	}

	return result
}

func main() {

	for {
		var choice string
		var word string
		var base int
		var n int64

		fmt.Println("Input Commands:")
		fmt.Println("conv <number> <base>   (any base to decimal)")
		fmt.Println("dec <number> <base>    (decimal to base)")
		fmt.Println("Type 'help' or 'quit'")
		fmt.Print("> ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Error: invalid input")
			continue
		}

		if choice == "quit" {
			fmt.Println("Goodbye")
			return
		}

		if choice == "help" {
			fmt.Println("Examples:")
			fmt.Println("conv 101 2   = 5")
			fmt.Println("dec 5 2      = 101")
			continue
		}

		if choice == "conv" {
			_, err = fmt.Scan(&word, &base)
			if err != nil {
				fmt.Println("Error: enter valid input")
				continue
			}

			fmt.Println("Result:", BaseConverter(word, base))

		} else if choice == "dec" {
			_, err = fmt.Scan(&n, &base)
			if err != nil {
				fmt.Println("Error: enter valid input")
				continue
			}

			fmt.Println("Result:", dex(n, base))

		} else {
			fmt.Println("Unknown command. Type 'help'")
		}
	}
}
