package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Salary float64
}

// func print(numbers *[3]int) {
// 	fmt.Println(numbers)
// }

func main() {
	Imjamam := User{ //instance or obj
		Name:   "Siam",
		Age:    21,
		Salary: 29,
	}

	x := &Imjamam

	fmt.Println(*x)

	//pointer or address memory(ram)
	// x := 50
	// fmt.Println("x =", x) // x=50

	// p := &x //ampersand & => address of |  824,633,755,368
	// *p = 70

	// fmt.Println("x =", x)

	// fmt.Println("Address: ", p)               //p is the address of x
	// fmt.Println("Value at the address: ", *p) // * => value at address

	// arr := [3]int{1, 2, 3}
	// print(&arr)

}

/*
2.Phases:

	1.Compilation Phase (compile time)
	2.Execution Phase (run time)


***********Compile Phase ********

***Code Segment***
*/
