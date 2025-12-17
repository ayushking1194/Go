package main

import (
	"fmt"
	// "reflect"
)

func main() {
	// Array of product names
	// Note: Arrays in Go have a fixed size, so we define it with a specific size.
	// If you need a dynamic size, use slices instead.
	var productNames [4]string
	productNames[0] = "Laptop"
	productNames[1] = "Smartphone"
	productNames[2] = "Tablet"
	productNames[3] = "Smartwatch"
	// fmt.Println("Product Names :", productNames)
	
	sliceProductNames := productNames[1:2] // Slicing the array to get a subset
	// fmt.Println("Slice of Product Names :", sliceProductNames)
	sliceProductNames = sliceProductNames[:3]
	// fmt.Println("Extended Slice of Product Names :", sliceProductNames)

	// Slice of product names
	// Slices are more flexible than arrays and can grow or shrink in size.
	prices := []float64{999.99, 499.99, 299.99, 199.99}
	// fmt.Println("Prices :", prices)
	prices = append(prices, 149.99) // Adding a new price
	// fmt.Println("Prices :", prices)

	products := []Product{
		{Name: "Laptop", Price: 999.99},
		{Name: "Smartphone", Price: 499.99},
	}
	// fmt.Println("Products :", products)

	details := []Product{
		{Name: "Tablet", Price: 299.99},
		{Name: "Smartwatch", Price: 199.99},
	}
	products = append(products, details...) // Appending another slice of products
	// fmt.Println("Updated Products :", products)
	

	// Using a map to associate website names with their urls
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Amazon Web Services": "https://aws.amazon.com",
	}
	fmt.Println("Websites :", websites)
	// Adding a new website
	websites["GitHub"] = "https://github.com"
	// fmt.Println("Updated Websites :", websites)
	// Iterating over the map
	// for name, url := range websites {	
	 	// fmt.Printf("Website: %s, URL: %s\n", name, url)
	// }
	// Deleting a website	
	delete(websites, "Google")
	// fmt.Println("Websites after deletion :", websites)

	// Using a slice to store user names with make
	userName := make([]string,2,5)	// Initializing a slice with a length of 2 and capacity of 5
	userName = append(userName, "John")
	userName = append(userName, "Doe")
	fmt.Println("User Names :", userName)
	// fmt.Println("userName[0] : ",userName[0],"Type of userName[0] : ",reflect.TypeOf(userName[0]))

	// Using a map to store user details
	userDetails := make(map[string]string,2) // Initializing a map with a capacity of 2
	userDetails["name"] = "John Doe"
	userDetails["email"] = "john.doe@example.com"
	fmt.Println("User Details :", userDetails)
}
type Product struct {
	Name  string
	Price float64
}