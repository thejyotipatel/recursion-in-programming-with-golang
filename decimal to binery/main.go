package main

import (
	"fmt"
	"strconv"
)

func decimalToBinery( decimal int,result string) string  {
	
	 if decimal == 0 {
		 return	result
		}

		result =  strconv.Itoa(decimal % 2) + result

		return decimalToBinery(decimal / 2, result)
}

func main() {
	fmt.Println("Decimal to binery")

	input := decimalToBinery(233, "")

	fmt.Println(input)
}