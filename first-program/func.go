//1 : A general function

package main

import "fmt"

// Simple function to add two numbers
func add(a int, b int) int {
	result := a + b
	fmt.Println("Sum:", result) // Output: Sum: 30
	return result               // Add a return statement
}

func main() {
	a := 10
	b := 20
	add(a, b) // Call the function
}

//2 : many return value

package main

import "fmt"

// Function to return quotient and remainder

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q)  // Output: Quotient: 3
	fmt.Println("Remainder:", r) // Output: Remainder: 1
}


//3  : default argument function

package main

import "fmt"

// Function to find the sum of any number of integers
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5)) // Output: Sum: 15
}


//4 : recursive function

package main

import "fmt"

// Recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("Factorial of 5:", factorial(5)) // Output: Factorial of 5: 120
}


//5 :anonymous function or lamda function
package main

import "fmt"

func main() {
	// Anonymous function assigned to a variable
	multiply := func(a int, b int) int {
		return a * b
	}

	fmt.Println("Product:", multiply(4, 5)) // Output: Product: 20
}

//6 :Differ function

package main

import "fmt"

func main() {
	defer fmt.Println("This will run last.")
	fmt.Println("This will run first.")
}


//7 :using pointer a function

package main

import "fmt"

// Function to update a value using a pointer

func updateValue(val *int) {
	*val = 100
}

func main() {
	num := 50
	fmt.Println("Before:", num)
	updateValue(&num)
	fmt.Println("After:", num) // Output: After: 100
}
