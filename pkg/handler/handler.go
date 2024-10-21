package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp())
		auth.POST("/sign-in", h.signIn())
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists())
			lists.GET("/:id", h.getListById())
			lists.POST("/", h.creatList())
			lists.PUT("/:id", h.updateList())
			lists.DELETE("/:id", h.deleteList())
			items := lists.Group(":id/is")
			{
				items.POST("/", h.creatItem())
				items.PUT("/:item_id", h.updateItem())
				items.GET("/", h.getAllItems())
				items.GET("/:item_id", h.getItemById())
				items.DELETE("/:item_id", h.deleteItem())
			}

		}

	}
	return router
}
