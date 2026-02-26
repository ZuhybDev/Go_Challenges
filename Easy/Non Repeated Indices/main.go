package main

import "slices"

func nonRepeatedIndices(arr []string) []int {

	count := make(map[string]int)
	result := make([]int, 0)

	for _, v := range arr {
		count[v]++
	}

	for i, k := range arr {
		if count[k] == 1 {
			result = append(result, i)
		}
	}
	slices.Sort(result)
	return result
}

func main() {
	arr := []string{
		"a", "fc", "ab", "a", "ab", "b",
	}

	nonRepeatedIndices(arr)
}
