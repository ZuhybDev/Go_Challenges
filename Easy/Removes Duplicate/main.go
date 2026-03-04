package main

import (
	"fmt"
	"strings"
)

// removeDuplicateVowels removes consecutive identical vowels from a string.
func removeDuplicateVowels(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result strings.Builder

	for i := 0; i < len(s); i++ {
		char := s[i]
		// Logic: Keep the character if it's the first one,
		// OR if it's not a vowel,
		// OR if it's a vowel but different from the previous character.
		if i == 0 || char != s[i-1] {
			result.WriteByte(char)
		}
	}

	return result.String()
}

func main() {
	// Test cases
	test1 := "goooogle"
	test2 := "vacuum"
	test3 := "aaaaaah"
	test4 := "keep"

	fmt.Printf("Original: %s -> Result: %s\n", test1, removeDuplicateVowels(test1))
	fmt.Printf("Original: %s -> Result: %s\n", test2, removeDuplicateVowels(test2))
	fmt.Printf("Original: %s -> Result: %s\n", test3, removeDuplicateVowels(test3))
	fmt.Printf("Original: %s -> Result: %s\n", test4, removeDuplicateVowels(test4))
}
