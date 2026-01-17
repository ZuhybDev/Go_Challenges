package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go start point")

	for i := 1; i < 100; i++ {
		fmt.Printf("%02d\n", i)
	}
}
