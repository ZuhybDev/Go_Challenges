package main

import "fmt"

func calculate(nums []float64) {

	sum, average := 0.0, 0.0
	min := nums[0]
	max := nums[1]

	for i, n := range nums {

		sum += n

		if min > nums[i] {
			min = nums[i]
		}

		if max < nums[i] {
			max = nums[i]
		}
	}

	average = sum / float64(len(nums))

	fmt.Printf("%.1f ", sum)
	fmt.Printf("%v ", average)
	fmt.Printf("%.1f ", min)
	fmt.Printf("%g ", max)

}
