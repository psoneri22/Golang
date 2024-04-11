package main

import (
	"net/http"

	"github.com/pluralsight/webservice/controllers"
)

// func main() {
// 	u := models.User{
// 		ID:        1,
// 		FirstName: "Cloudester",
// 		LastName:  "cld",
// 	}
// 	fmt.Println(u)
// 	port := 3000
// 	_, isStart := startWebServer(port, 2)
// 	fmt.Println(isStart)
// }

//	func startWebServer(port, numberOfRetries int) (int, error) {
//		fmt.Println("starting server...")
//		fmt.Println("starting started on port", port)
//		fmt.Println("number Of Retries", numberOfRetries)
//		return port, nil
//	}
func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
