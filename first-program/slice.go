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

	var s []int //empty slice or nill slice
	fmt.Println(s)
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
*/
