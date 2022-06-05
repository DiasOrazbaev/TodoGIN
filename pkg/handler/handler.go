package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.login)
		auth.POST("/register", h.register)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAll)
			lists.POST("/", h.createList)
			lists.GET("/:id", h.getListByID)
			lists.DELETE("/:id", h.deleteList)
			lists.PUT("/:id", h.updateList)

			items := lists.Group(":id/items")
			{
				items.GET("/", h.getAllItem)
				items.POST("/", h.createItem)
				items.GET("/:item_id", h.getItemByID)
				items.DELETE("/:item_id", h.deleteItem)
				items.PUT("/:item_id", h.updateItem)
			}
		}
	}
	return router
}
