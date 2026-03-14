package main

import (
	"fmt"
)

func printDiamondPattern(n int) {

	for i := 1; i <= n; i++ {
		// Print leading spaces
		for j := 1; j <= n-i; j++ {
			fmt.Printf(" ")
		}
		// Print numbers
		for k := 1; k <= i; k++ {

			fmt.Printf("%d ", k)

		}
		fmt.Println()
	}
}

func main() {
	printDiamondPattern(6)
}
