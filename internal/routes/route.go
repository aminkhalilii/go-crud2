package routes

import (
	"crud/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.SaveBook)
}
