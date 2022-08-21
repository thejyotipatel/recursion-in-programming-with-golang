package main

import (
	"fmt"
)

// 1. base case, find when recursion will stop
// 2. 
func reverse(input string) string {
	if input == "" {
		return ""
	}

	//  input[1:] + input[:1]
	return reverse(input[1:]) + input[:1]
}

func main() {
	fmt.Println("Reverse string")

	rev := reverse("hello")
	fmt.Println(rev)

}