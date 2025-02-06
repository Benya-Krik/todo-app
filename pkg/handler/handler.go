package handler

import (
	"eduProject/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	ping := router.Group("/ping")
	{
		ping.GET("/", func(c *gin.Context) {})
	}

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/signIn", h.signIn)

		auth.POST("/logout", h.Logout)

	}
	api := router.Group("/api", h.GetUserID)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.POST("/", h.createList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group("/:id/items")
			{
				items.GET("/", h.getAllItems)
				items.POST("/", h.createItem)
			}
		}
		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}
	return router
}
