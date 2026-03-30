package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BaseConverter(word string, base int) (int64, error) {
	return strconv.ParseInt(word, base, 64)
}

func dex(n int64, base int) string {
	result := strconv.FormatInt(n, int(base))
	return strings.ToUpper(result)
}

func main() {

	for {
		var choice string
		var word string
		var base string
		var n int64

		fmt.Println("Input Commands:")
		fmt.Println("convert <number> <base>")
		fmt.Println("base: hex | bin | dec")
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
			fmt.Println("convert 1E hex   = 30")
			fmt.Println("convert 10 bin   = 2")
			fmt.Println("convert 255 dec  = 11111111 and FF")
			continue
		}

		if choice == "convert" {

			_, err = fmt.Scan(&word, &base)
			if err != nil {
				fmt.Println("Error: enter valid input")
				continue
			}

			base = strings.ToLower(base)

			if base == "hex" {
				result, err := BaseConverter(word, 16)
				if err != nil {
					fmt.Println("Error: invalid hex")
					continue
				}
				fmt.Println("✦ Decimal:", result)

			} else if base == "bin" {
				result, err := BaseConverter(word, 2)
				if err != nil {
					fmt.Println("Error: invalid binary")
					continue
				}
				fmt.Println("✦ Decimal:", result)

			} else if base == "dec" {
				// handle negative numbers too
				n, err = strconv.ParseInt(word, 10, 64)
				if err != nil {
					fmt.Println("Error: invalid decimal")
					continue
				}

				fmt.Println("✦ Binary:", dex(n, 2))
				fmt.Println("✦ Hex:", dex(n, 16))

			} else {
				fmt.Println("Error: base must be hex, bin, or dec")
			}

		} else {
			fmt.Println("Unknown command. Type 'help'")
		}
	}
}
