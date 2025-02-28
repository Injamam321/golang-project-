package main

import "fmt"

func main() {

	// Function expression or assign function in variable
	// a := 10 //expression

	add := func(a int, b int) {
		c := a + b //Expression
		fmt.Println(c)

	}
	add(8, 9)

}

func init() {
	fmt.Println("i am call first")
}
