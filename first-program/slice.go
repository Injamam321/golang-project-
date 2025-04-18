package main

import "fmt"

func main() {

	// s := []int{1, 2, 3, 4, 5, 6} //slice literal
	// fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))

	// s := make([]int, 3) //[0, 0, 0] len = 3, cap = 3
	// s[0] = 19
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// s := make([]int, 3, 6) //[0, 0, 0] len = 3, cap = 6
	// s[0] = 19              //[19, 0, 0] len = 3, cap = 6
	// s[2] = 20              //[19, 0, 20] len = 3, cap = 6
	// fmt.Println(s)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))

	// arr := [9]string{"i", "love", "Go", "programming", "language", "i", "want", "remote", "job"} // BODY
	// fmt.Println(arr)
	// s := arr[1:5]
	// fmt.Println(s) // [love, Go, programming, language] // HAAT
	// s1 := s[1:3]
	// fmt.Println(s1)      // [Go, programming] //RING
	// fmt.Println(len(s1)) //length ->2
	// fmt.Println(cap(s1)) //

	// var s []int //empty slice or nill slice
	// s = append(s, 1, 2, 9)
	// fmt.Println(s)

	var x []int //empty slice or nill slice
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x

	x = append(x, 4)
	y = append(y, 5)

	x[0] = 10

	fmt.Println(x) //[10, 2, 3, 5]

	fmt.Println(y) // [10, 2, 3, 5]
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
