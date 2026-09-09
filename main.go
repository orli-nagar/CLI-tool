package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var text string
var oldString string
var newString string

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

func replaceString() string {
	return strings.ReplaceAll(text, oldString, newString)
}

var countCmd = &cobra.Command{
	Use:   "count <text>",
	Short: "Count the words and characters in a string",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		text := args[0]
		fmt.Printf("Words: %d | Characters: %d\n", countWords(text), countCharacters(text))
	},
}
var reverseCmd = &cobra.Command{
	Use:   "reverse <text>",
	Short: "Reverse a string",
	Args:  cobra.ExactArgs(1),
	Run: func(_ *cobra.Command, args []string) {
		fmt.Println(reverseString(args[0]))
	},
}
var replaceCmd = &cobra.Command{
	Use:   "replace <input>",
	Short: "Replace a string with a new string",
	Run: func(_ *cobra.Command, args []string) {
		text = args[0]
		fmt.Println(replaceString())
	},
}
var rootCmd = &cobra.Command{
	Use:   "strutils",
	Short: "A CLI tool for string utilities",
}

func main() {
	replaceCmd.Flags().StringVarP(
		&oldString,
		"old",
		"o",
		"",
		"The old string to replace",
	)

	replaceCmd.Flags().StringVarP(
		&newString,
		"new",
		"n",
		"",
		"The new string to replace with",
	)

	replaceCmd.MarkFlagRequired("old")

	rootCmd.AddCommand(countCmd, reverseCmd, replaceCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
