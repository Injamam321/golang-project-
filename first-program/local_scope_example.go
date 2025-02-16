package main

import "fmt"

func calculateSum(a int, b int) int {
	sum := a + b // 'sum' is a local variable
	return sum
}

func main() {
	result := calculateSum(5, 7) // Calling function
	fmt.Println("Sum:", result)

	// fmt.Println(sum) // ❌ This will cause an error! (sum is not accessible here)
}
