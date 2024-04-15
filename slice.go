package main

import (
	"fmt"
	"slices"
)

// Create a Slice From an Array
// func main() {
// 	arr1 := [6]int{10, 11, 12, 13, 14, 15}
// 	myslice := arr1[2:4]

//		fmt.Printf("myslice = %v\n", myslice)
//		fmt.Printf("length = %d\n", len(myslice))
//		fmt.Printf("capacity = %d\n", cap(myslice))
//	}

// create slices using the make() function
// func main() {
// 	myslice1 := make([]int, 5, 10)
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))

// 	// with omitted capacity
// 	myslice2 := make([]int, 5)
// 	fmt.Printf("myslice2 = %v\n", myslice2)
// 	fmt.Printf("length = %d\n", len(myslice2))
// 	fmt.Printf("capacity = %d\n", cap(myslice2))
// }

// You can append elements to the end of a slice using the append()
// func main() {
// 	myslice1 := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))

// 	myslice1 = append(myslice1, 20, 21, 22, 23, 24)
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))
// }

// Append One Slice To Another Slice
// The '...' after slice2 is necessary when appending the elements of one slice to another
// func main() {
// 	myslice1 := []int{1, 2, 3}
// 	myslice2 := []int{4, 5, 6, 7}
// 	myslice3 := append(myslice1, myslice2...)

// 	fmt.Printf("myslice3=%v\n", myslice3)
// 	fmt.Printf("length=%d\n", len(myslice3))
// 	fmt.Printf("capacity=%d\n", cap(myslice3))
// }

// Change The Length of a Slice
// func main() {
// 	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
// 	myslice1 := arr1[1:5]                 // Slice array
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))

// 	myslice1 = arr1[1:3] // Change length by re-slicing the array
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))

// 	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
// 	fmt.Printf("myslice1 = %v\n", myslice1)
// 	fmt.Printf("length = %d\n", len(myslice1))
// 	fmt.Printf("capacity = %d\n", cap(myslice1))
// }

// Memory Efficiency
//
//	When using slices, Go loads all the underlying elements into the memory.
//
// If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
// The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program.
// func main() {
// 	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
// 	// Original slice
// 	fmt.Printf("numbers = %v\n", numbers)
// 	fmt.Printf("length = %d\n", len(numbers))
// 	fmt.Printf("capacity = %d\n", cap(numbers))

// 	// Create copy with only needed numbers
// 	neededNumbers := numbers[:len(numbers)-8]
// 	numbersCopy := make([]int, len(neededNumbers))
// 	copy(numbersCopy, neededNumbers)

// 	fmt.Printf("numbersCopy = %v\n", numbersCopy)
// 	fmt.Printf("length = %d\n", len(numbersCopy))
// 	fmt.Printf("capacity = %d\n", cap(numbersCopy))
// }

func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 5)
	for i := 0; i < 5; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
