package main

import (
	"bufio"
	"fmt"
	"os"
	"product_management/handlers"
)

func showMenu() {
	fmt.Println("\n--------Menu--------")
	fmt.Println("Choose a number (0-5): ")
	fmt.Println("1. View Products")
	fmt.Println("2. Add Product")
	fmt.Println("3. Edit Product")
	fmt.Println("4. Delete Product")
	fmt.Println("5. Find product by name")
	fmt.Println("0. Exit")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			handlers.ListProducts()
		case "2":
			handlers.CreateProduct(scanner)
		case "3":
			handlers.UpdateProduct(scanner)
		case "4":
			handlers.DeleteProduct(scanner)
		case "5":
			handlers.SearchProduct(scanner)
		case "0":
			return
		default:
			fmt.Println("Invalid selection!")
		}
	}
}