package router

import (
	"web-service-gin/controllers"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	err := r.Run("localhost:8080")
	if err != nil {
		return
	}
}
