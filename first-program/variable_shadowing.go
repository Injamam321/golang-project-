package main

import "fmt"

var a = 10

func main() {
	age := 30

	if age > 18 {
		a := 47        //atai variable shadowing agartar
		fmt.Println(a) // 47 print korbe
	}
	fmt.Println(a) //10 print korbe
}
