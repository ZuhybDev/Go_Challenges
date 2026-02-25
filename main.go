package main

import "fmt"

func performOperations(nums []int) {

	result := 0
	sumOfEven := 1
	sumOfOdd := 1

	for _, n := range nums {

		if n%2 == 0 {
			sumOfEven *= n
		} else {
			sumOfOdd *= n
		}

	}

	result = sumOfEven + sumOfOdd

	fmt.Println(result)
}

func main() {

	nums := []int{20, 67, 456, 89, 9, 122, 33}
	// nums := []int{1, 2, 3, 4, 5, 6}

	performOperations(nums)
}
