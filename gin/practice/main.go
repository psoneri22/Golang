// main.go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Create a new form validator
	validator := NewValidator()
	// validator.AddRule("name", Required)
	validator.AddRule("email", Required)
	validator.AddRule("email", Email)

	// Register form validator middleware globally
	r.Use(FormValidator(validator))

	// Example route that requires form validation
	r.POST("/submit", func(c *gin.Context) {
		// If form validation passes, this handler will be executed
		c.JSON(200, gin.H{"message": "Form submitted successfully"})
	})

	// Run the server
	r.Run(":8080")
}
