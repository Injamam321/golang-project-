package main

import "fmt"

var a = 15

func main() {
	fmt.Println(a)
}

func init() {
	fmt.Println(a)
	a = 30
}
