package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to CLI base converter Tool")
	var (
		n       int64
		base    int
		word    string
		command int
	)

	for {

		fmt.Println("Option 1 convert to any Base | option 2 dec to base |  option 3 Exit Programme ")
		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if command < 1 || command > 3 {
			fmt.Println("error! invlid command token")
			continue
		}
		if command == 1 {
			fmt.Println("Please enter word to convert")
			fmt.Scan(&word)

			fmt.Println("Please enter base")
			fmt.Scan(&base)

			fmt.Println(BaseConverter(word, base))
		}
		if command == 2 {
			fmt.Println("please enter a digit.")
			fmt.Scan(&n)

			fmt.Println("Enter base:")
			fmt.Scan(&base)

			fmt.Println(dex(n, base))
		}
		if command == 3 {
			fmt.Println("goodbye buddy! sse ya later")
			break
		}
	}
}
