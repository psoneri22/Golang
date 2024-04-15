package main

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// func main() {
// 	r := gin.Default()

// 	r.GET("/isprime/:num", func(c *gin.Context) {
// 		numStr := c.Param("num")
// 		num, err := strconv.Atoi(numStr)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number"})
// 			return
// 		}

// 		if isPrime(num) {
// 			c.JSON(http.StatusOK, gin.H{"number": num, "prime": true})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"number": num, "prime": false})
// 		}
// 	})

// 	port := 8080
// 	fmt.Printf("Server running on port %d\n", port)
// 	r.Run(fmt.Sprintf(":%d", port))
// }
