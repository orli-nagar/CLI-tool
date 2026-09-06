package main

import (
	"fmt"
	"os"
	"strings"
)

func countWords(s string) int {
	return len(strings.Fields(s))
}

func countCharacters(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <count|reverse> <text>")
		return
	}

	command := os.Args[1]

	if command != "count" && command != "reverse" {
		fmt.Println("Invalid command. Use 'count' or 'reverse'.")
		return
	}

	if len(os.Args) < 3 {
		fmt.Println("Please provide a string to process.")
		return
	}

	if len(os.Args) > 3 {
		fmt.Println("Please provide the text as one argument, using quotes if it contains spaces.")
		return
	}

	text := os.Args[2]

	switch command {
	case "count":
		fmt.Printf("Words: %d | Characters: %d\n",
			countWords(text),
			countCharacters(text),
		)

	case "reverse":
		fmt.Println(reverseString(text))
	}
}
