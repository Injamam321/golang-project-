package main

import "fmt"

var arr2 = [3]string{"imjamam", "siam", "mehedi"}

func main() {
	var arr [4]int

	arr[1] = 8
	arr[3] = 10

	fmt.Println(arr)
	fmt.Println(arr2)
}
