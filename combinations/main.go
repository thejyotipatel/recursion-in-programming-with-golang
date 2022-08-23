package main

import "fmt"

func combinations(element []string) []string {
if len(element) == 0 {
	return []string{}
}
const firstEl = element[:1]
const rest = element[1:]

combsWithoutFirst := combinations(rest)
combsWithFirst := []string{}

// combsWithoutFirst
		for _ , comb := range combsWithoutFirst{
			combWithFirst := []string{comb,firstEl[0]}
			combsWithFirst = append(combsWithFirst, combWithFirst...)
			// return combWithFirst
		}

return combsWithFirst
// ,combsWithFirst
}

func main() {
	fmt.Println("Combinations")

	comb := combinations([]string{"a","b","c"})
fmt.Println(comb)
}