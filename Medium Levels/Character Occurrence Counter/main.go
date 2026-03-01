package main

// import "fmt"

func countOccurrences(char, str string) int {

	var count int

	for _, v := range str {

		if char == string(v) {
			count++
			continue
		}
	}
	return count
}

func main() {

	str := "abdhgsejdfb"
	char := "a"

	countOccurrences(char, str)
}
