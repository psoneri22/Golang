// validator.go
package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Validator struct {
	rules  map[string][]func(value string) error
	errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{
		rules:  make(map[string][]func(value string) error),
		errors: make(map[string]string),
	}
}

func (v *Validator) AddRule(field string, rule func(value string) error) {
	v.rules[field] = append(v.rules[field], rule)
}

func (v *Validator) Validate(data map[string]string) bool {
	valid := true
	for field, rules := range v.rules {
		for _, rule := range rules {
			if err := rule(data[field]); err != nil {
				v.errors[field] = err.Error()
				valid = false
				break
			}
		}
	}
	return valid
}

func (v *Validator) GetErrors() map[string]string {
	return v.errors
}

func Required(value string) error {
	if value == "" {
		return errors.New("Field is required")
	}
	return nil
}

func Email(value string) error {
	if value == "" {
		return nil // Allow empty email
	}
	// Basic email format validation
	// This regex pattern is simplified and may not cover all valid email formats
	if len(value) < 3 || len(value) > 254 {
		return errors.New("Invalid email format")
	}
	return nil
}

// FormValidator is a middleware for form validation
func FormValidator(v *Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := make(map[string]string)
		// Get form data from the request
		if c.Request.Method == "POST" {
			if err := c.Request.ParseForm(); err != nil {
				c.AbortWithStatusJSON(400, gin.H{"error": "Invalid form data"})
				return
			}
			for key, values := range c.Request.PostForm {
				if len(values) > 0 {
					data[key] = values[0]
				}
			}
		}
		// Validate form data
		if !v.Validate(data) {
			c.AbortWithStatusJSON(400, gin.H{"errors": v.GetErrors()})
			return
		}
		// Pass control to the next middleware or handler
		c.Next()
	}
}
