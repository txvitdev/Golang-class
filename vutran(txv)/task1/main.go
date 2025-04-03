package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"txv/task1/list"
	"txv/task1/product"
)

func main() {
	arrayList := list.NewArrayList[product.Product]()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Product Management Menu ---")
		fmt.Println("1. Add Product")
		fmt.Println("2. Delete Product")
		fmt.Println("3. View Products")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addProduct(scanner, arrayList)
		case "2":
			deleteProduct(scanner, arrayList)
		case "3":
			viewProducts(arrayList)
		case "4":
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func addProduct(scanner *bufio.Scanner, arrayList *list.ArrayList[product.Product]) {
	fmt.Print("Enter Product ID: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter Product Name: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Enter Product Price: ")
	scanner.Scan()
	price, _ := strconv.ParseFloat(scanner.Text(), 64)

	arrayList.Add(product.Product{Id: id, Name: name, Price: price})
	fmt.Println("Product added successfully!")
}

func deleteProduct(scanner *bufio.Scanner, arrayList *list.ArrayList[product.Product]) {
	fmt.Print("Enter Product ID to delete: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < arrayList.Size(); i++ {
		p, _ := arrayList.Get((i))
		if p.Id == id {
			arrayList.Remove(i)
			fmt.Println("Product deleted successfully!")
			return
		}
	}
	fmt.Println("Product not found!")
}

func viewProducts(arrayList *list.ArrayList[product.Product]) {
	if arrayList.Size() == 0 {
		fmt.Println("No products available.")
		return
	}

	fmt.Println("\nProduct List:")
	arrayList.Print()
}
