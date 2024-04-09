package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/higatu/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user_banner := router.Group("/user_banner")
	{
		user_banner.GET("/tag_id", h.getTagId)	
		user_banner.GET("/feature_id", h.getFeatureId)
		user_banner.GET("/use_last_version", h.getLastVersion)
		user_banner.GET("user_token", h.getUserToken)

	}

	banner := router.Group("/banner")
	{
		banner.GET("/token", h.getToken)
		banner.GET("/feature_id", h.getFeatureId)
		banner.GET("/tag_id", h.getTagId)
		banner.GET("/limit", h.getLimit)
		banner.GET("/offset", h.getOffset)
		banner.POST("/token", h.postToken)

			id := banner.Group("/:id")
			{
				id.PATCH("/id", h.patchId)
				id.PATCH("/token", h.patchToken)
				id.DELETE("/id", h.deleteId)
				id.DELETE("/token", h.deleteToken)
			}
		
	}
	
	return router
}
