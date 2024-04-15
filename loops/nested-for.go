package main

import (
	"fmt"
)

func main() {
	adj := [2]string{"big", "tasty"}
	fruits := [...]string{"apple", "orange", "banana", "guavava"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}
