package main

import "fmt"

func longestString(str []string) {

	var longest string
	var length int

	for _, v := range str {
		if len(v) > length {
			longest = v
			length = len(v)
		}
	}

	fmt.Printf("%s %d\n", longest, length)

}
