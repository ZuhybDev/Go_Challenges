package main

import "fmt"

func checkNumber(n int) {

	sum := 0
	i := 1

	for n > i {

		if n%i == 0 {
			sum += i
		}
		i++

	}

	if sum > n {
		fmt.Println("Abundant")
	}
	if sum == n {
		fmt.Println("Perfect")
	}

	if sum < n {
		fmt.Println("Deficient")
	}

}
