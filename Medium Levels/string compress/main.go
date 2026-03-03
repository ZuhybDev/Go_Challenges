package main

import "fmt"

func compress(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			result += string(s[i-1])
			for j := 0; j < count; j++ {
				result += "#"
			}
			count = 1
		}
	}

	// last group
	result += string(s[len(s)-1])
	for j := 0; j < count; j++ {
		result += "#"
	}

	return result
}

func main() {
	fmt.Println(compress("aabccchbbccaaa"))
}
