package main

import "fmt"

func sumAndDivide(num1, num2 int) float64{

    addingBoth := num1 + num2

	sum := 0

for num1 <= num2 {
   	sum += num1
	num1++
}

result := float64(sum) / float64(addingBoth)

return  result
}

func main() {

	var num1 int
	var num2 int

	fmt.Scanln(&num1)
	fmt.Scanln(&num2)

	sumAndDivide(num1, num2)
}
