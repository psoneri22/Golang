package main

import "fmt"

// func main() {
// 	var a = 15 + 25
// 	fmt.Println(a)
// }

// func main() {
// 	var (
// 		sum1 = 100 + 50    // 150 (100 + 50)
// 		sum2 = sum1 + 250  // 400 (150 + 250)
// 		sum3 = sum2 + sum2 // 800 (400 + 400)
// 	)
// 	fmt.Println(sum3)
// }

func main() {
	// x := 22
	// y := 5
	// x++
	// fmt.Println("Addition:", x+y)
	// fmt.Println("Substraction:", x-y)
	// fmt.Println("Multiplication:", x*y)
	// fmt.Println("Divison:", x/y)
	// fmt.Println("Modulus:", x%y)
	// fmt.Println("increment:", x)
	c := complex(3, 4)
	fmt.Println(c)

	r, i := real(c), imag(c)
	fmt.Println(r, i)
}