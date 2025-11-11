package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/http"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/repository"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/maxcore25/bmstu-it-courses/backend/docs"
)

// @title Gin Example API
// @version 1.0
// @description This is a sample Gin server with Swagger
// @host localhost:8080
// @BasePath /api

func main() {
	dsn := "postgres://postgres:powerfuldb22@localhost:5432/bmstu_it_courses?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	jwtManager := utils.NewJWTManager(
		"supersecretaccesskey123",
		"supersecretrefreshkey456",
	)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	authService := service.NewAuthService(userRepo, refreshTokenRepo, jwtManager)

	http.RegisterAuthRoutes(r, userService, authService, nil)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Println("🚀 Dev server started at http://localhost:8080")
	fmt.Println("🚀 Swagger started at http://localhost:8080/docs/index.html")

	r.Run(":8080")
}
