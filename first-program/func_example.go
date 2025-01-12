package main

import "fmt"

func printWelcomeMessege() {
	//print welcome messege
	fmt.Println("wellcome to the application")
}

func getUserName() string {
	//get user name as a input
	var name string
	fmt.Println("Enter your name -")
	fmt.Scanln(&name)
	return name
}
func getTwoNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter first number -")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number -")
	fmt.Scanln(&num2)
	return num1, num2
}
func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}
func Display(name string, sum int) {
	//Display result or print
	fmt.Println("Hello, ", name)
	fmt.Println("Summation =", sum)
}

func Goodbye() {
	// Print a Goodbye me
	fmt.Println("Thank you for using my application")
	fmt.Println("Goodbye")
}
func main() {

	printWelcomeMessege()
	name := getUserName()
	num1, num2 := getTwoNumbers()
	sum := add(num1, num2)
	Display(name, sum)
	Goodbye()

}

//SOLID-
// S - Single Responsibility Principle (SRP)
// O - Open/Closed Principle (OCP)
// L - Liskov Substitution Principle (LSP)
// I - Interface Segregation Principle (ISP)
// D - Dependency Inversion Principle (DIP)
