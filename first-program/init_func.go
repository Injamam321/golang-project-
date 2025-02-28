package main

import "fmt"

var a = 20

func main() {
	fmt.Println("Hello init function i love you")
}

func init() {
	fmt.Println("i am first function for executable") // first ata call hobe and ata first print hobe
}
