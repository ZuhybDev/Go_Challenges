package main

import "fmt"

func main() {

	res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0

	min := res[0]
	max := res[1]
	for i, v := range res {

		sum += v

		if max < res[i] {
			max = res[i]
		}

		if min > res[i] {
			min = res[i]
		}
	}

	fmt.Println("Sum: ", sum)

	fmt.Println("Max: ", max)
	fmt.Println("Min: ", min)
}
