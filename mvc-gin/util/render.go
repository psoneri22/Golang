// util/render.go
package util

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderTemplate(c *gin.Context, templateFiles []string, data ...interface{}) {
	var templateData interface{}
	if len(data) > 0 {
		templateData = data[0]
	} else {
		templateData = nil // Set default value to nil
	}

	t, err := template.ParseFiles(templateFiles...)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = t.Execute(c.Writer, templateData)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
}
