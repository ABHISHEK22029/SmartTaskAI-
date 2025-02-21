package main

import (
	"github.com/gin-gonic/gin"
	"backend/middleware"
	"backend/repository"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/login", repository.Login)

	protected := router.Group("/api")
	protected.Use(middleware.JWTMiddleware())
	{
		protected.POST("/task-suggestions", repository.GetTaskSuggestions)
	}

	router.GET("/ws", func(c *gin.Context) {
		middleware.HandleWebSocket(c.Writer, c.Request)
	})

	router.Run(":8080")
}
