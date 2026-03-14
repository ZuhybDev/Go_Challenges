package main

import "fmt"

func printCharacters(str string) {
	n := len(str)

	// Growing
	for i := 1; i <= n; i++ {
		fmt.Println(str[:i])
	}

	// Reverse shrinking (remove from end)
	for i := n - 1; i > 0; i-- {
		fmt.Println(str[:i])
	}
}

func main() {
	printCharacters("hello")
}
