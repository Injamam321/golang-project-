package main

import "fmt"

func main() {
	arr := [9]string{"i", "love", "Go", "programming", "language", "i", "want", "remote", "job"} // BODY
	fmt.Println(arr)

	s := arr[1:5]
	fmt.Println(s) // [love, Go, programming, language] // HAAT

	s1 := s[1:3]
	fmt.Println(s1)      // [Go, programming] //RING
	fmt.Println(len(s1)) //length ->2
	fmt.Println(cap(s1)) //
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
*/
