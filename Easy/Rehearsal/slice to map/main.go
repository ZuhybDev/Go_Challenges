package main

import "fmt"

func slicetoMap(slice []string) {

	data := make(map[string]int)

	for _, v := range slice {
		if _, ok := data[v]; ok {
			continue
		} else {
			data[v] = len(v)
		}
	}
	fmt.Println(data)

}

func main() {

	data := []string{
		"main", "lamguage", "zuhybdev",
	}

	slicetoMap(data)

}
