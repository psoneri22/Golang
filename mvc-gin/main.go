// main.go
package main

import (
	"mvc-gin/routes"
)

func main() {
	r := routes.LoadRoutes()

	// Run the server
	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
