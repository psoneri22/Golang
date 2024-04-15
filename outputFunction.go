package main

// import "fmt"

// The Print() function prints its arguments with their default format.
// func main() {
// 	var i, j string = "Hello", "World"

// 	fmt.Print(i)
// 	fmt.Print(j)
// }

// If we want to print the arguments in new lines, we need to use \n.
// func main() {
// 	var i, j string = "Hello", "World"

// 	fmt.Print(i, "\n", j)
// }

// Print() inserts a space between the arguments if neither are strings:
// func main() {
// 	var i, j = 10, 20

// 	fmt.Print(i, "\n", j)
// }

// The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end
// func main() {
// 	var i, j string = "Hello", "World"

// 	fmt.Println(i, j)
// }

// The Printf() function first formats its argument based on the given formatting verb and then prints them.
// %v is used to print the value of the arguments
// %T is used to print the type of the arguments
// func main() {
// 	var i string = "Hello"
// 	var j int = 15

// 	fmt.Printf("i has value: %v and type: %T\n", i, i)
// 	fmt.Printf("j has value: %v and type: %T\n", j, j)
// }

// Formatting Verbs for Printf()
// %v	Prints the value in the default format
// %#v	Prints the value in Go-syntax format
// %T	Prints the type of the value
// %%	Prints the % sign
// func main() {
// 	var i = 15.5
// 	var txt = "Hello World!"

// 	fmt.Printf("%v\n", i)
// 	fmt.Printf("%#v\n", i)
// 	fmt.Printf("%v%%\n", i)
// 	fmt.Printf("%T\n", i)

// 	fmt.Printf("%v\n", txt)
// 	fmt.Printf("%#v\n", txt)
// 	fmt.Printf("%T\n", txt)
// }
