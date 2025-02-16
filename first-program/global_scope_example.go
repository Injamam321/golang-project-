package main

import "fmt"

// Global variable
var globalVar = "I'm a global variable"

func printGlobalVar() {
	fmt.Println(globalVar) // Accessing global variable
}

func main() {
	printGlobalVar()
	globalVar = "Global variable changed!"
	printGlobalVar()
}
