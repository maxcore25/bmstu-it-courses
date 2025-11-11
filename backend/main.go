package main

import (
	"fmt"

	_ "github.com/maxcore25/bmstu-it-courses/backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Example API
// @version 1.0
// @description This is a sample Gin server with Swagger
// @host localhost:8080
// @BasePath /api

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", pingHandler)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("🚀 Dev server started at http://localhost:8080")
	fmt.Println("🚀 Swagger started at http://localhost:8080/docs/index.html")
	r.Run(":8080")
}

// @Summary Ping example
// @Description Responds with pong
// @Tags example
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
