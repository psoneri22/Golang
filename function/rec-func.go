package main

import "fmt"

// func testcount(x int) int32 {
// 	if x == 11 {
// 		return 0
// 	}
// 	fmt.Println(x)
// 	return testcount(x + 1)
// }

// func main() {
// 	testcount(1)
// }

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	// Recursive case: factorial(n) = n * factorial(n-1)
	return n * factorial(n-1)
}

func main() {
	// Calculate the factorial of 5 using recursion
	result := factorial(6)
	fmt.Println("Factorial of 5:", result) // Output: Factorial of 5: 120
}
