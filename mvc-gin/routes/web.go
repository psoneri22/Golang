package routes

import (
	"mvc-gin/controllers"
	"mvc-gin/util"

	"github.com/gin-gonic/gin"
)

func LoadRoutes() *gin.Engine {
	// Create a new Gin engine
	r := gin.Default()
	// Set the template directory
	r.LoadHTMLGlob("views/*.tmpl")

	// Define controllers
	userController := controllers.NewUserController()
	contactController := controllers.NewContactController()

	// Define routes
	r.GET("/", HomeHandler)
	r.GET("/users", userController.Index)
	r.GET("/contacts", contactController.Index)

	return r
}

// Exported function
func HomeHandler(c *gin.Context) {
	// Define the path to the template file
	templates := []string{"views/layout.tmpl", "views/home.tmpl"}

	// Render the template with the provided data
	util.RenderTemplate(c, templates, nil)
}
