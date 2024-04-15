package main

import (
	"fmt"
)

// Comparison operators are used to compare two values.
// Note: The return value of a comparison is either true (1) or false (0)
func main() {
	var x = 5
	var y = 3
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)
}
