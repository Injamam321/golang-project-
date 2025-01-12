package main

import (
	"fmt"
	"strings"
)

// Function to convert the string to uppercase
func toUpperCase(input string) string {
	return strings.ToUpper(input)
}

// Function to count the number of words in the string
func wordCount(input string) int {
	words := strings.Fields(input)
	return len(words)
}

// Function to reverse the string
func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Function to check if the string is a palindrome
func isPalindrome(input string) bool {
	reversed := reverseString(input)
	return strings.ToLower(input) == strings.ToLower(reversed)
}

func main() {
	// Input string
	var input string
	fmt.Println("Enter a string:")
	fmt.Scanln(&input)

	// Applying functions
	upper := toUpperCase(input)
	wordCountResult := wordCount(input)
	reversed := reverseString(input)
	palindrome := isPalindrome(input)

	// Printing results
	fmt.Println("Original String:", input)
	fmt.Println("Uppercase:", upper)
	fmt.Println("Word Count:", wordCountResult)
	fmt.Println("Reversed String:", reversed)
	if palindrome {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}
