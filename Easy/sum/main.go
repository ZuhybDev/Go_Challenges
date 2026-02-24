package sum

func calculateSum(arr []int) int {

	first := arr[0]
	last := arr[len(arr)-1]
	midIndex := len(arr) / 2
	midElement := arr[midIndex]
	var sum int

	if len(arr)%2 == 0 {
		return -1
	}
	sum = first + last + midElement
	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	calculateSum(arr)
}
