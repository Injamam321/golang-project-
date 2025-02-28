package main

import "fmt"

//standard or named function
// func add(a, b int) {
// 	fmt.Println(a + b)
// }

func main() {
	// Invoke\calling add function
	// add(40, 7) //Expression

	// Anonymous function
	//Immediately Invoked  Function Exprssion
	//IIFE Function
	func(a int, b int) {
		c := a + b //Expression
		fmt.Println(c)

	}(5, 6)

}

func init() {
	fmt.Println("i am call first")
}
