package main

import "fmt"

// func main() {
// 	if 20 > 18 {
// 		fmt.Println("20 is greater than 18")
// 	}
// }

// func main() {
// 	var (
// 		x = 2
// 		y = 18
// 	)
// 	if x > y {
// 		fmt.Println("x is greater than y")
// 	} else {
// 		fmt.Println("y is greter than x")
// 	}
// }

// func main() {
// 	time := 18
// 	if time < 10 {
// 		fmt.Println("Good morning.")
// 	} else if time < 20 {
// 		fmt.Println("Good day.")
// 	} else {
// 		fmt.Println("Good evening.")
// 	}
// }

func main() {
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}
}
