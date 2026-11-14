package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Product struct {
	ID      int
	Name    string
	Price   float64
	InStock bool
}

func main() {
	user := User{
		Name:  "John Park",
		Email: "john@gmail.com",
	}
	user.Age = 24
	user.ID = 1

	fmt.Printf("User: %+v\n", user)

	// Named initialization
	product := Product{
		ID:      2,
		Name:    "Phone",
		Price:   40000,
		InStock: true,
	}

	fmt.Printf("Product 1 (Named): %+v\n", product)

	// Positional initialization
	product2 := Product{
		3,
		"Mouse",
		1200,
		true,
	}

	fmt.Printf("Product 2 (Positional): %+v\n", product2)

	// Zero-value initialization followed by field assignment
	productNew := Product{}

	productNew.ID = 4
	productNew.Name = "Laptop"
	productNew.Price = 75000
	productNew.InStock = false

	fmt.Printf("Product 3 (Zero value + assignment): %+v\n", productNew)
}
