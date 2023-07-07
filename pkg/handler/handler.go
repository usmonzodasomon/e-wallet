package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usmonzodasomon/e-wallet/pkg/service"
)

type handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *handler {
	return &handler{service}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		})

		// auth := api.Group("/auth")
		// {
		// 	// auth.POST("/sign-up", h.signUp)
		// 	// auth.POST("/sign-in", h.signIn)
		// }

	}
	return router
}
