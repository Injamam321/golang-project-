package main

import "fmt"

func main() {
	fmt.Println(PackageVar) // ✅ utils.go থেকে এক্সেস করা যাচ্ছে
	PrintMessage()          // ✅ utils.go থেকে ফাংশন এক্সেস করা যাচ্ছে
}
