// controllers/user_controller.go
package controllers

import (
	"mvc-gin/models"
	"mvc-gin/util"

	"github.com/gin-gonic/gin"
)

// UserController handles user-related requests
type ContactController struct{}

// NewContactController creates a new ContactController
func NewContactController() *ContactController {
	return &ContactController{}
}

// Index handles GET /users
func (ca *ContactController) Index(c *gin.Context) {
	contacts := models.GetAllContacts()
	// c.HTML(http.StatusOK, "contact_view.tmpl", gin.H{"contact": contacts})
	data := map[string]interface{}{
		"contact": contacts,
	}
	templates := []string{"views/layout.tmpl", "views/contact_view.tmpl"}
	util.RenderTemplate(c, templates, data)
}
