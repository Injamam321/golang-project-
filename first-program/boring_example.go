package main

import "fmt"

var (
	a = 10
	b = 20
)

func printNum(num int) {
	fmt.Println(num)
}

func add(a int, b int) int {
	res := a + b
	printNum(res)
	return res
}

func main() {
	add(a, b)
}
