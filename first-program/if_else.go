//example-1
// package main

// import "fmt"

// func main() {
// 	// Declare a variable
// 	age := 10

//		// Check conditions with if, else if, and else
//		if age < 13 {
//			fmt.Println("You are a child.")
//		} else if age >= 13 && age < 18 {
//			fmt.Println("You are a teenager.")
//		} else if age >= 18 && age < 60 {
//			fmt.Println("You are an adult.")
//		} else {
//			fmt.Println("You are a senior citizen.")
//		}
//	}

//example-2

// package main

// import "fmt"

// func main() {
// 	number := 10

// 	if number > 0 {
// 		fmt.Println("The number is positive.")
// 	} else if number < 0 {
// 		fmt.Println("The number is negative.")
// 	} else {
// 		fmt.Println("The number is zero.")
// 	}
// }

//example -3

// package main

// import "fmt"

// func main() {
// 	number := 8

// 	if number%2 == 0 {
// 		fmt.Println("The number is even.")
// 	} else {
// 		fmt.Println("The number is odd.")
// 	}
// }

//example-4

package main

import "fmt"

func main() {
	year := 2025

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Println("The year is a leap year.")
	} else {
		fmt.Println("The year is not a leap year.")
	}
}
