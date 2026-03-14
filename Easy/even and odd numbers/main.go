package main

func performOperations(nums []int) int {

	var (
		sumOfEven int
		sumOfOdd  int = 1
		result    int
	)

	for _, n := range nums {

		if n%2 == 0 {
			sumOfEven += n
		} else {
			sumOfOdd *= n
		}
	}

	result = sumOfEven + sumOfOdd

	return result
}

func main() {
	num := []int{1, 2, 4, 5, 6, 7, 8}

	performOperations(num)
}
