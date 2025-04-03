package handlers

import (
	"bufio"
	"fmt"
	"product_management/models"
	"strconv"
	"strings"
)
var (
	products      = make(map[int]*models.Product)
	nextProductID = 1
)


func ListProducts() {
	if len(products) == 0 {
		fmt.Println("No products found.")
		return
	}

	fmt.Println("Product List:")
	for _, p := range products {
		fmt.Printf("- ID: %d, Name: %s, Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}


func CreateProduct(scanner *bufio.Scanner) {
	fmt.Println("Enter product name:")
	scanner.Scan()
	nameStr := scanner.Text()

	fmt.Println("Enter product price:")
	scanner.Scan()
	priceStr := scanner.Text()
	price64, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		fmt.Println("Invalid price.")
		return
	}

	product := &models.Product{
		ID:    nextProductID,
		Name:  nameStr,
		Price: float32(price64),
	}

	products[nextProductID] = product
	nextProductID++

	fmt.Println("Product added successfully.")
}

func UpdateProduct(scanner *bufio.Scanner) {
	fmt.Println("Enter the product ID to update:")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	product, ok := products[id]
	if !ok {
		fmt.Println("Product not found.")
		return
	}

	fmt.Printf("Current Name: %s\n", product.Name)
	fmt.Printf("Current Price: %.2f\n", product.Price)

	fmt.Print("Do you want to update the name? (y/n): ")
	scanner.Scan()
	AcceptNewName := strings.ToLower(scanner.Text())

	newName := product.Name
	
	if AcceptNewName == "y" || AcceptNewName == "yes" {
		fmt.Print("Enter the new name: ")
		scanner.Scan()
		newName = scanner.Text()
	}

	fmt.Print("Do you want to update the price? (y/n): ")
	scanner.Scan()
	acceptNewPrice := strings.ToLower(scanner.Text())

	newPrice := product.Price
	if acceptNewPrice == "y" || acceptNewPrice == "yes" {
		fmt.Print("Enter the new price: ")
		scanner.Scan()
		priceStr := scanner.Text()
		priceF64, err := strconv.ParseFloat(priceStr, 32)
		if err != nil {
			fmt.Println("Invalid price.")
			return
		}
		newPrice = float32(priceF64)
	}

	product.Name = newName
	product.Price = newPrice
	products[id] = product

	fmt.Println("Product updated successfully.")
}

func DeleteProduct(scanner *bufio.Scanner) {
	fmt.Println("Enter the product ID to delete:")
	scanner.Scan()
	idStr := scanner.Text()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}
	product, ok := products[id]
	if !ok {
		fmt.Println("Product not found.")
		return
	}

	fmt.Printf("Product Info: ID: %d- Name: %s- Price: %.2f\n", product.ID, product.Name, product.Price)
	fmt.Print("Are you sure? (y/n): ")
	scanner.Scan()

	acceptDeleteProduct := strings.ToLower(scanner.Text())
	if acceptDeleteProduct != "y" && acceptDeleteProduct != "yes" {
		return
	}
	delete(products, id)
	fmt.Println("Product deleted successfully.")
}

func SearchProduct(scanner *bufio.Scanner) {
	fmt.Println("Enter the product name to search:")
	scanner.Scan()
	nameStr := strings.ToLower(scanner.Text())

	for _, product := range products {
		if strings.Contains(strings.ToLower(product.Name), nameStr) {
			fmt.Printf("- ID: %d, Name: %s, Price: %.2f\n", product.ID, product.Name, product.Price)
		}
	}
}