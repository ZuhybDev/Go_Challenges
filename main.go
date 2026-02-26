package main

import "fmt"

func nonRepeatedIndices(arr []string) []int  {
	
    dataMap := make(map[int]struct)
    result := make([]int, 0)

    for i := range arr {
        if _, ok := dataMap[i]; ok {
            delete(dataMap, i)
        }else {
            result = append(result, i)
            dataMap[i] = struct{}{}
        }
    }

 return result   
}
func main() {

	nums := []int{20, 67, 456, 89, 9, 122, 33}
	// nums := []int{1, 2, 3, 4, 5, 6}

	performOperations(nums)
}
