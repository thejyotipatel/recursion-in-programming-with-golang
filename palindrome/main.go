package main

import "fmt"

func palindrome( input string) bool  {
	if len(input) == 1 || len(input) == 0 {
		return true
	}

	if input[:1] == input[(len(input) - 1):] {
		return palindrome(input[1:(len(input) - 1)])
	}
	
	// if not palindram
	return false
}

func main() {
	fmt.Println("Decimal to binery")

	input := palindrome("racecar")
	// input := "racecar" 
	fmt.Println("is palindrome: ",input)
}