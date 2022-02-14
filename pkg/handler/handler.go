package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/oleg5896/go-web/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	list := router.Group("/list")
	{
		list.GET("/", h.getList)
	}

	item := router.Group("/items")
	{
		item.POST("/add/", h.addItem)
		item.GET("/add/", h.getItem)
		item.GET("/", h.getItem)
	}
	return router
}
