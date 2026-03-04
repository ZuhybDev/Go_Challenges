// Abundant Number
// Purpose: Determine whether a positive integer is an abundant number
// Example: call `isAbundantNumber(96)` — the program prints the result to stdout
// Note: This is documentation only; no logic changes were made.
package main

import "fmt"

func isAbundantNumber(n int) {

	i := 1
	sum := 0
	isAbundant := false
	for n > i {

		if sum > n {
			isAbundant = true
		} else {
			isAbundant = false
		}
		if n%i == 0 {
			sum += i
		}
		i++

	}

	fmt.Println(isAbundant)

}

func main() {
	isAbundantNumber(96) // example

}
