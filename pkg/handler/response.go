package handler

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}

var templates = template.Must(template.ParseFiles("public/upload.html"))

func Display(c gin.Context, page string, data interface{}) {
	templates.ExecuteTemplate(c.Writer, page+".html", data)
}
