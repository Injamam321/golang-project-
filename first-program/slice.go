package main

import "fmt"

func main() {
	arr := [5]string{"i", "love", "Go", "programming", "language"}
	fmt.Println(arr)

	s := arr[1:3]
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
