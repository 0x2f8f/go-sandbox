package main

import "fmt"

func main() {
	var products = [3]string{
		"bread",
		"butter",
	}

	products[2] = "milk"

	fmt.Printf("Count: %d\n", len(products))

	fmt.Printf("products: %v\n", products)

	fmt.Println(products)

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
	}

	for index, value := range products {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	for _, v := range products {
		fmt.Println(v)
	}
}
