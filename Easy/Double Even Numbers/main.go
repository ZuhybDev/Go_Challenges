package main

// import (
//     "slices"
// )

func doubleEvenNumbers(nums []int) []int {

	odd, even, conc := make([]int, 0), make([]int, 0), make([]int, 0)

	for _, n := range nums {
		if n%2 != 0 {
			odd = append(odd, n)
		} else {
			n += n
			even = append(even, n)
		}
	}

	conc = append(odd, even...)

	// slices.Sort(conc)

	return conc
}
