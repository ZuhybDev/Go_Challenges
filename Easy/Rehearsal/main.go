package main

import "fmt"

func longestString(arr []string) {

	var longest string
	count := 0

	for _, v := range arr {

		if len(v) > len(longest) {
			longest = v
			count = len(v)
		}
	}

	fmt.Printf("%s %d", longest, count)
}

func main() {
	array := []string{"abc", "go", "zuhybDev"}

	longestString(array)

}
