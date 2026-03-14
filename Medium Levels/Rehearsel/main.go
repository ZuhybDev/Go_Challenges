package main

func nonRepeatedIndices(arr []string) []int {

	counter := make(map[string]int)
	idxSlice := make([]int, 0)

	for _, v := range arr {
		if _, ok := counter[v]; ok {
			counter[v]++
		} else {
			counter[v]++
		}
	}

	for i, v := range arr {
		if counter[v] == 1 {
			idxSlice = append(idxSlice, i)
		}
	}

	return idxSlice
}

func main() {
	arr := []string{
		"ag", "fc", "ab", "a", "ab", "b",
	}
	nonRepeatedIndices(arr)
}
