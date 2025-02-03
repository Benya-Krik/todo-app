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

	auth := router.Group("/auth")
	{
		auth.POST("login", h.Login)
		auth.POST("signIn", h.signIn)

		auth.POST("logout", h.Logout)

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

			items := lists.Group("/items")
			{
				items.GET("/", h.getAllItems)
				items.POST("/", h.createItem)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.getAllItems)
				items.DELETE("/:item_id", h.getAllItems)
			}
		}
	}
	return router
}
