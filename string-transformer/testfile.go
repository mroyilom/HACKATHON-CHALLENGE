// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: Your Name
// Squad: Your Squad

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	for {

		var command string
		var word string
		var words []string

		fmt.Println("Input Commands:")
		fmt.Println("upper <text>")
		fmt.Println("lower <text>")
		fmt.Println("cap <text>")
		fmt.Println("title <text>")
		fmt.Println("snake <text>")
		fmt.Println("reverse <text>")
		fmt.Println("Type 'help' or 'exit'")
		fmt.Print("> ")

		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println("Error: invalid input")
			continue
		}

		command = strings.ToLower(command)

		if command == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			return
		}

		if command == "help" {

			fmt.Println("Help instructions:")
			fmt.Println("upper text   → converts all letters to UPPERCASE")
			fmt.Println("lower text   → converts all letters to lowercase")
			fmt.Println("cap text     → capitalizes each word")
			fmt.Println("title text   → title case with small words lowercase")
			fmt.Println("snake text   → converts text to snake_case")
			fmt.Println("reverse text → reverses each word")
			fmt.Println("Example: upper hello world")
			continue
		}

		words = []string{}

		for {
			n, _ := fmt.Scan(&word)

			if n == 0 {
				break
			}

			words = append(words, word)

			if len(words) > 50 {
				break
			}
		}

		if len(words) == 0 {
			fmt.Println("Error: no text provided. Usage:", command, "<text>")
			continue
		}

		text := strings.Join(words, " ")

		if command == "upper" {
			fmt.Println("✦ Result:", toUpper(text))

		} else if command == "lower" {
			fmt.Println("✦ Result:", toLower(text))

		} else if command == "cap" {
			fmt.Println("✦ Result:", toCap(text))

		} else if command == "title" {
			fmt.Println("✦ Result:", toTitle(text))

		} else if command == "snake" {
			fmt.Println("✦ Result:", toSnake(text))

		} else if command == "reverse" {
			fmt.Println("✦ Result:", toReverse(text))

		} else {
			fmt.Println("Unknown command. Type 'help'")
		}
	}
}

func toUpper(text string) string {
	return strings.ToUpper(text)
}

func toLower(text string) string {
	return strings.ToLower(text)
}

func toCap(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		word := strings.ToLower(words[i])

		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}
	}

	return strings.Join(words, " ")
}

func toTitle(text string) string {

	small := map[string]bool{
		"a": true, "an": true, "the": true, "and": true,
		"but": true, "or": true, "for": true, "nor": true,
		"on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "as": true,
		"is": true, "it": true,
	}

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		word := strings.ToLower(words[i])

		if i == 0 || !small[word] {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		} else {
			words[i] = word
		}
	}

	return strings.Join(words, " ")
}

func toSnake(text string) string {

	var result []rune

	for _, ch := range text {

		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			result = append(result, unicode.ToLower(ch))

		} else if ch == ' ' {
			result = append(result, '_')
		}
	}

	return string(result)
}

func toReverse(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		words[i] = reverseWord(words[i])
	}

	return strings.Join(words, " ")
}

func reverseWord(word string) string {

	runes := []rune(word)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
