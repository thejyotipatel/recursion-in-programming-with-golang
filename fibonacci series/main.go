package main

import "fmt"

// func fibonacci(memo map[int]int,num int) int {
func fibonacci(num int) int {
 // for _, item := range memo {
	// 	// fmt.Println(item)
	// 	if item == num{
	// 		return memo[num]
	// 	}
	// }

	if num <= 2 {
		return 1
	}
	// memo[num] = fibonacci(memo, num-1) + fibonacci(memo, num-2)
	// return memo[num]
	return fibonacci(num-1) + fibonacci(num-2)
}
// func add(num1 int, num2[] int) int {
// 	return num1 + num2[0]
// }

func main() {
	fmt.Println("fibonacci series")
	// n := make(map[int]int)
// fib := fibonacci(n,50)
fib := fibonacci(15)
// fib := add(2, n)
	fmt.Println("fab:", fib)
}