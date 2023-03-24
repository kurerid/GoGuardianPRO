package handler

import (
	"GuardianPRO/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

		order := api.Group("/order")
		{
			private := order.Group("/private")
			{
				private.POST("", h.createPrivateOrder)
			}

			group := order.Group("/group")
			{
				group.GET("/list", h.getGroupOrderList)
			}
		}

		account := api.Group("/account")
		{
			account.GET("/list", h.getAccountList)
		}
	}
	return router
}
