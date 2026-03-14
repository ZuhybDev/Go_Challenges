package main

import "strings"

func countWords(str string) int {

	words := strings.Split(strings.TrimSpace(str), " ")
	return len(words)

}

func main() {

	words := "learn with zohaibDev"
	countWords(words)
}
