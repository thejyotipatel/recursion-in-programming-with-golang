package main

import "fmt"

func sum(num int, result int) int {
	if num <= 1 {
		return num
	}

	result = num + sum(num - 1, result)
	return result
}

func main() {
	fmt.Println("sum of natural number i.e. 1, 2, 3...")

	s := sum(2, 0)
	fmt.Println(s)
}