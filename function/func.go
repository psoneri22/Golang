package main

import "fmt"

// func myMessage() {
// 	fmt.Println("I just got executed!")
// }

// func main() {
// 	myMessage() // call the function
// }

// Function With Parameter
// func familyName(fname string) {
// 	fmt.Println("Hello", fname, "Refsnes")
// }

// func main() {
// 	familyName("Liam")
// 	familyName("Jenny")
// 	familyName("Anja")
// }

// Multiple Parameters
// func familyName(fname string, age int) {
// 	fmt.Println("Hello", age, "year old", fname, "Refsnes")
// }

// func main() {
// 	familyName("Liam", 3)
// 	familyName("Jenny", 14)
// 	familyName("Anja", 30)
// }

// Function Returns
// func myFunction(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	fmt.Println(myFunction(1, 2))
// }

// Named Return Values
// func myFunction(x int, y int) (result int) {
// 	result = x + y
// 	return
// }

// func main() {
// 	fmt.Println(myFunction(1, 2))
// }

// Store the Return Value in a Variable
// func myFunction(x int, y int) (result int) {
// 	result = x + y
// 	return
// }

// func main() {
// 	total := myFunction(1, 2)
// 	fmt.Println(total)
// }

// Multiple Return Values
// we can add an underscore (_), to omit this value
func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}
func main() {
	a, _ := myFunction(5, "Hello")
	fmt.Println(a)
}

//	func main() {
//		_, b := myFunction(5, "Hello")
//		fmt.Println(b)
//	}
//	func main() {
//		fmt.Println(myFunction(5, "Hello"))
//	}
// func main() {
// 	a, b := myFunction(5, "Hello")
// 	fmt.Println(a, b)
// }
