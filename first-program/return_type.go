package main

import "fmt"

// Function to add two numbers
func add(a int, b int) int {
	sum := a + b
	return sum // Returns the sum of a and b
}

// Function to return sum and product
func getnumbers(a int, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

func main() {
	num1 := 10
	num2 := 20

	// Using the add function
	sum := add(num1, num2)
	fmt.Println("Sum from add function:", sum)

	// Using the getnumbers function
	p, q := getnumbers(num1, num2)
	fmt.Println("Sum from getnumbers function:", p)
	fmt.Println("Product from getnumbers function:", q)
}
