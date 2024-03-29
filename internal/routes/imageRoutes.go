package routes

import (
	"github.com/Abedmuh/Paimon-bank/internal/image"
	"github.com/gin-gonic/gin"
)

func ImageRoutes(route *gin.RouterGroup) {
	controller := image.NewImageController()

	route.POST("/image", controller.PostImage)
}