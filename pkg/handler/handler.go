package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/oleg5896/go-web/pkg/service"
)

type Handler struct {
	services *service.Service
}

func newHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	list := router.Group("/list")
	{
		list.GET("/", h.getList)
	}

	item := router.Group("/item")
	{
		item.POST("/:id", h.addItem)
		item.GET("/:id", h.getItem)
	}
	return router
}
