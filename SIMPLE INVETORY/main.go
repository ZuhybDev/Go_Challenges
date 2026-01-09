package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// TODO: Define the Product struct here
type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	// Read input using bufio.Scanner to handle spaces properly
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	storeInfo := scanner.Text()

	scanner.Scan()
	productData := scanner.Text()

	// TODO: Write your code below
	// 1. Parse store information (split by comma, check length)
	storeParts := strings.Split(storeInfo, ",")
	storeName := storeParts[0]
	storeLocation := storeParts[1]
	// 2. Parse product data (split by comma, then by colon for each product)
	productEntries := strings.Split(productData, ",")

	// 3. Create inventory map
	// 4. Convert strings to appropriate types and populate inventory

	inventory := make(map[string]Product)
	for _, entry := range productEntries {
		productParts := strings.Split(entry, ":")
		productName := productParts[0]
		productPrice, _ := strconv.ParseFloat(productParts[1], 64)
		productQuantity, _ := strconv.Atoi(productParts[2])
		inventory[productName] = Product{
			Name:     productName,
			Price:    productPrice,
			Quantity: productQuantity,
		}
	}
	// 5. Display store information
	fmt.Printf("=== %s Inventory System ===\n", storeName)
	fmt.Printf("Location: %s\n", storeLocation)
	fmt.Printf("Inventory initialized with %d products\n", len(inventory))
	// 6. Display current inventory (sorted alphabetically)
	sortedProduct := make([]Product, 0)
	fmt.Println("Current Inventory:")

	for _, val := range inventory {
		sortedProduct = append(sortedProduct, Product{
			Name:     val.Name,
			Price:    val.Price,
			Quantity: val.Quantity,
		})
	}

	sort.Slice(sortedProduct, func(i, j int) bool {
		return sortedProduct[i].Name < sortedProduct[j].Name
	})

	for _, p := range sortedProduct {
		fmt.Printf("- %s: $%.2f (Stock: %d)\n", p.Name, p.Price, p.Quantity)
	}

	// 7. Calculate and display inventory statistics
	fmt.Println("Inventory Statistics:")

	fmt.Printf("Total Products: %d\n", len(sortedProduct))

	totalItems := 0
	totalValue := 0.0
	for _, item := range sortedProduct {
		totalItems += item.Quantity
		totalValue += item.Price * float64(item.Quantity)
	}
	// 8. Display system status
	fmt.Printf("Total Items in Stock: %d\n", totalItems)
	fmt.Printf("Total Inventory Value: $%g\n", totalValue)
	fmt.Println("System Status: Ready")
	fmt.Println("Inventory management system initialized successfully")

}
