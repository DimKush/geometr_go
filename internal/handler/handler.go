package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/DimKush/geometry_go/internal/service"
)

type Handler struct {
	services service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Content-type", "Authorization"}
	config.AllowMethods = []string{"POST", "GET", "DELETE", "PUT"}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	router.GET("/:id", h.getWarehouseById)
	router.GET("/unit/:id", h.getUnit)
	router.POST("/setUnit", h.setUnit)
	return router
}

func InitHandler(service *service.Service) *Handler {
	return &Handler{
		services: *service,
	}
}
