package main

import "fmt"

// global variable declaration
var (
	a = 20
	b = 30
)

func add(x int, y int) {
	z := x + y
	fmt.Println(z) // Corrected `fmt.println` to `fmt.Println`
}

func main() {
	p := 30
	q := 40
	add(p, q) // This will correctly print the sum
	add(a, q)
	add(p, b)
}
