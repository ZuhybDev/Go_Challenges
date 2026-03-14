package main

import "fmt"

func ScheduleMove(department string, itemCount int, truckNumber int) string {
	return fmt.Sprintf("Department %s with %d items is assigned to truck %d.\n", department, itemCount, truckNumber)
}

func main() {
	ScheduleMove("main", 5, 6)
}
