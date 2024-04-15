package main

import (
	"fmt"
)

// range to iterate over an array and print both the indexes and the values at each (idx stores the index, val stores the value)
//
//	func main() {
//		fruits := [3]string{"apple", "orange", "banana"}
//		for idx, val := range fruits {
//			fmt.Printf("%v\t%v\n", idx, val)
//		}
//	}
//

// // omit the indexes (idx stores the index, val stores the value):
// func main() {
// 	fruits := [3]string{"apple", "orange", "banana"}
// 	for _, val := range fruits {
// 		fmt.Printf("%v\n", val)
// 	}
// }

// omit the indexes (idx stores the index, val stores the value):
func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}
}
