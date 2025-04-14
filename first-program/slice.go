package main

import "fmt"

func main() {
	arr := [9]string{"i", "love", "Go", "programming", "language", "i", "want", "remote", "job"}
	fmt.Println(arr)

	s := arr[1:5]
	fmt.Println(s) // [love go programming language]

	s1 := s[1:3]
	fmt.Println(s1)
}

/*
2.Phases:

	1.Compilation Phase (compile time)
	2.Execution Phase (run time)


***********Compile Phase ********

***Code Segment***
main=func(){...}



*/
