package main

import "fmt"

// func main() {
// 	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
// 	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

// 	fmt.Printf("a\t%v\n", a)
// 	fmt.Printf("b\t%v\n", b)
// }

// create map using make function
// func main() {
// 	var a = make(map[string]string) // The map is empty now
// 	a["brand"] = "Ford"
// 	a["model"] = "Mustang"
// 	a["year"] = "1964"
// 	// a is no longer empty
// 	b := make(map[string]int)
// 	b["Oslo"] = 1
// 	b["Bergen"] = 2
// 	b["Trondheim"] = 3
// 	b["Stavanger"] = 4

// 	// Update and Add Map Elements
// 	a["year"] = "1970" // Updating an element
// 	a["color"] = "red" // Adding an element

// 	delete(a, "model")
// 	fmt.Printf("a\t%v\n", a)
// 	fmt.Printf("b\t%v\n", b)
// 	// accessing the value
// 	fmt.Println(a)
// }

// Create an Empty Map
// func main() {
// 	var a = make(map[string]string)
// 	var b map[string]string

// 	fmt.Println(a == nil)
// 	fmt.Println(b == nil)
// }

// Check For Specific Elements in a Map
func main() {
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

	val1, ok1 := a["brand"] // Checking for existing key and its value
	val2, ok2 := a["color"] // Checking for non-existing key and its value
	val3, ok3 := a["day"]   // Checking for existing key and its value
	_, ok4 := a["model"]    // Only checking for existing key and not its value

	if ok1 {
		fmt.Println(val1)
	}

	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)
}
