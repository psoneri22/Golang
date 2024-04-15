// controllers/user_controller.go
package controllers

import (
	"mvc-gin/models"
	"mvc-gin/util"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related requests
type UserController struct{}

// NewUserController creates a new UserController
func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) Index(c *gin.Context) {
	// Get all users from the model
	users := models.GetAllUsers()

	// Define data to be passed to the template
	data := map[string]interface{}{
		"user": users,
	}

	// Define the path to the template file
	templates := []string{"views/layout.tmpl", "views/user_view.tmpl"}

	// Render the template with the provided data
	util.RenderTemplate(c, templates, data)
}
