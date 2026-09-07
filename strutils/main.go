package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
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

var countCmd = &cobra.Command{
	Use:   "count <text>",
	Short: "Count the words and characters in a string",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		text := args[0]
		fmt.Printf("Words: %d | Characters: %d\n", countWords(text), countCharacters(text))
	},
}
var reverseCmd = &cobra.Command{
	Use:   "reverse <text>",
	Short: "Reverse a string",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(reverseString(args[0]))
	},
}
var rootCmd = &cobra.Command{
	Use:   "strutils",
	Short: "A CLI tool for string utilities",
}

func main() {
	rootCmd.AddCommand(countCmd, reverseCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
