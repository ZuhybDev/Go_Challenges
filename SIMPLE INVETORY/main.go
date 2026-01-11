package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func updateStock(inventory map[string]Product, quantity int, name string) error {
	if _, ok := inventory[name]; !ok {
		return fmt.Errorf("product not found: %s", name)
	} else if qty, _ := inventory[name]; qty.Quantity+quantity < 0 {
		absQty := int(math.Abs(float64(qty.Quantity)))
		return fmt.Errorf("insufficient stock: cannot reduce %s by %d, only %d available", name, absQty, qty.Quantity)
	} else {
		qty := inventory[name]
		inventory[name] = Product{
			Quantity: qty.Quantity + quantity,
			Price:    qty.Price,
			Name:     qty.Name,
		}

	}

	return nil
}

func main() {

	var existingData string = "Laptop:999.99:5"
	var updatingData string = "Laptop:4"

	fmt.Scanln(&existingData)
	fmt.Scanln(&updatingData)

	productEntries := strings.Split(existingData, ",")

	inventory := map[string]Product{}
	for _, product := range productEntries {
		trimmedList := strings.TrimSpace(product)
		entry := strings.Split(trimmedList, ":")

		name := entry[0]
		price, _ := strconv.ParseFloat(entry[1], 64)
		quantity, _ := strconv.Atoi(entry[2])

		inventory[name] = Product{
			Name:     name,
			Price:    price,
			Quantity: quantity,
		}
	}
	processedItems := 0
	successfulItems := 0
	failedItems := 0
	stockUpdates := strings.Split(updatingData, ",")
	fmt.Println("Processing Stock Updates:")

	var failedProducts []string

	for _, product := range stockUpdates {
		trimmedData := strings.TrimSpace(product)

		entry := strings.Split(trimmedData, ":")

		name := entry[0]
		quantity, _ := strconv.Atoi(entry[1])
		err := updateStock(inventory, quantity, name)
		key, _ := inventory[name]
		if err != nil {
			fmt.Printf("%s: Update failed - %v\n", name, err)
			failedProducts = append(failedProducts, name)
			failedItems++
		} else if quantity > 0 {
			fmt.Printf("%s: Added %d units - New stock: %d\n", name, quantity, key.Quantity)
			successfulItems++
		} else if quantity < 0 {
			fmt.Printf("%s: Removed %d units - New stock: %d\n", name, -quantity, key.Quantity)
			successfulItems++
		} else {
			fmt.Printf("%s: No change - Current stock: %d\n", name, key.Quantity)
			successfulItems++
		}

		processedItems++

	}

	fmt.Println("Update Summary:")
	fmt.Printf("Updates processed: %d\n", processedItems)
	fmt.Printf("Updates successful: %d\n", successfulItems)
	fmt.Printf("Updates failed: %d\n", failedItems)

	//"Final Inventory Statistics:"

	totalStock := 0
	totalRevenue := 0.0
	fmt.Println("Updated Inventory:")

	sortedList := make([]Product, 0)
	for _, pro := range inventory {
		totalStock += pro.Quantity
		totalRevenue += pro.Price * float64(pro.Quantity)

		sortedList = append(sortedList, Product{
			Name:     pro.Name,
			Price:    pro.Price,
			Quantity: pro.Quantity,
		})
	}

	sort.Slice(sortedList, func(i, j int) bool {
		return sortedList[i].Name < sortedList[j].Name
	})

	for _, pro := range sortedList {
		fmt.Printf("- %s: $%.2f (Stock: %d)\n", pro.Name, pro.Price, pro.Quantity)
	}

	fmt.Println("Final Inventory Statistics:")
	fmt.Printf("Total Products: %d\n", len(inventory))
	fmt.Printf("Total Items in Stock: %d\n", totalStock)
	fmt.Printf("Total Inventory Value: $%.2f\n", totalRevenue)

	if len(failedProducts) == 0 {

		fmt.Println("All stock updates were processed successfully")
	} else {
		fmt.Printf("Failed updates: %s\n", strings.Join(failedProducts, ","))
	}
}
