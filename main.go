package main

import (
	"fmt"
	"os"
)

func createFile() {
	file, err := os.Create("text.txt")

	if err != nil {
		fmt.Println("Error while creating", err)
	}

	fmt.Println(string(file.Name()))
}

func openFile() {
	file, err := os.Open("text.txt")

	if err != nil {
		fmt.Println("Error while opening", err)
	}

	fmt.Println(file.Name())
}

func readFile() {
	file, err := os.ReadFile("text.txt")

	if err != nil {
		fmt.Println("Error while reading", err)
	}

	fmt.Println(string(file))
}

// func writeFile() {
// 	// err := os.WriteFile("text.txt", []byte("Hello i changes from inside function"))

// 	if err != nil {
// 		fmt.Println("Error while writing ", err)
// 	}
// }

func main() {
	fmt.Println("Hello World")
	// createFile()
	// openFile()
	readFile()
}
