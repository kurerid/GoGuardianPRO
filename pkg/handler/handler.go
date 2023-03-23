package handler

import (
	"GuardianPRO/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up")
			auth.POST("/sign-in")
		}

		order := api.Group("/order")
		{
			private := order.Group("/private")
			{
				private.POST("")
			}

			group := order.Group("/group")
			{
				group.GET("/list")
			}
		}
	}
	return router
}
