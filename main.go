package main

import "fmt"

func main() {
	// Example with an odd number of elements
	oddArray := []int{100, 4, 6, 45, 8, 8, 67}
	middleIndexOdd := len(oddArray) / 2
	middleElementOdd := oddArray[middleIndexOdd]
	fmt.Printf("Odd length array: %v\n", oddArray)
	fmt.Printf("Middle element (index %d): %d\n\n", middleIndexOdd, middleElementOdd)

	// Example with an even number of elements
	evenSlice := []string{"apple", "banana", "cherry", "date"}
	middleIndexEven := len(evenSlice) / 2 // 4 / 2 = 2

	fmt.Printf("Even length slice: %v\n", evenSlice)

	// For an even length, you may consider one or both of the two middle elements
	// The index calculation gives the second of the two middle elements
	middleElementEven1 := evenSlice[middleIndexEven-1]
	middleElementEven2 := evenSlice[middleIndexEven]
	fmt.Printf("Middle elements (indices %d and %d): %s and %s\n", middleIndexEven-1, middleIndexEven, middleElementEven1, middleElementEven2)
}
