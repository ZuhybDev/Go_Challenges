package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Orders struct {
	Vendor   string
	Quantity int
}

func processMarketInventory(inventory []string, sales []string) []string {

	inventoryMap := make(map[string]Orders)

	for _, v := range inventory {
		entries := strings.Split(v, ":")

		vendor := entries[0]
		name := entries[1]
		quantity, _ := strconv.Atoi(entries[2])

		inventoryMap[name] = Orders{
			Vendor:   vendor,
			Quantity: quantity,
		}

	}

	soldOutItems := make([]string, 0)

	for _, sale := range sales {
		entries := strings.Split(strings.TrimSpace(sale), ":")

		name := entries[0]
		amount, _ := strconv.Atoi(entries[1])

		if value, ok := inventoryMap[name]; ok {

			value.Quantity -= amount
			inventoryMap[name] = value
			if value.Quantity < 0 || value.Quantity == 0 {
				soldOutItems = append(soldOutItems, name)
			}

		} else {
			fmt.Println("Items does'nt exist")
		}

	}

	slices.Reverse(soldOutItems)

	return soldOutItems
}

func main() {

	inventory := []string{
		"I:fig:2", "J:grape:2", "K:fig:3",
	}
	sales := []string{
		"fig:5", "grape:2",
	}

	processMarketInventory(inventory, sales)

}
