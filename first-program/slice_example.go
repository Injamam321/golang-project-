package main

import "fmt"

func changeslice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeslice(a)

	fmt.Println(x) //[1,2,3,4,10,6,7]
	fmt.Println(y) //[10,6,7,11]

	fmt.Println(x[0:8]) //[1,2,3,4,10,6,7,11]
}

/*
2.Phases:

	1.Compilation Phase (compile time)
	2.Execution Phase (run time)


***********Compile Phase ********

***Code Segment***
main=func(){...}

*/

/*
1.slice from an existing array
2.slice from an slice
3.slice literal
4.make function with len
5.make function with len & capacity
6.empty slice or nil slice
7.slice underlying array rule => 1024 -> 100% increase, 1024 < 25% increase
*/
