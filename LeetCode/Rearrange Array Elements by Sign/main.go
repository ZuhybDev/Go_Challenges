package main

import "fmt"

func rearrangeArray(nums []int) []int {

	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	result := make([]int, 0)

	for _, n := range nums {
		if n > 0 {
			arr1 = append(arr1, n)
		} else {
			arr2 = append(arr2, n)
		}
	}

	for i := 0; i < len(arr1); i++ {
		result = append(result, arr1[i])
		result = append(result, arr2[i])
	}

	return result
}

func main() {

	nums := []int{3, 1, -2, -5, 2, -4}

	result := rearrangeArray(nums)

	fmt.Println("Original:", nums)
	fmt.Println("Rearranged:", result)

}
