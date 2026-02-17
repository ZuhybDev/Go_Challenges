package main

import (
	"fmt"
)

// Very optimized and for speed and memory safety and large datasets
/*
func generatePattern(n int) {
    var line strings.Builder // drifting board

    for i := n; i >= 1; i-- {
        line.Reset()
        for j := 1; j <= i; j++ {
            line.WriteString(strconv.Itoa(j)) // Efficiently add numbers to the line
        }
        fmt.Println(line.String()) // Print the whole row at once
    }
}
*/

func generatePattern(n int) {
	if n > 10 {
		fmt.Println("The number should be less than 10")
		return
	}

	// Outer loop: Start at n, count down to 1
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func main() {
	n := 10 // example

	generatePattern(n)
}
