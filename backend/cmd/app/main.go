package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/http"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/repository"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/config"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/maxcore25/bmstu-it-courses/backend/docs"
)

// @title BMSTU IT Courses API
// @version 1.0
// @description This is a sample Gin server with Swagger
// @host localhost:8080
// @BasePath /api

func main() {
	config.LoadEnv()

	// Read environment variables
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")

	jwtAccessSecret := os.Getenv("JWT_ACCESS_SECRET")
	jwtRefreshSecret := os.Getenv("JWT_REFRESH_SECRET")
	accessExpStr := os.Getenv("JWT_ACCESS_EXPIRATION")
	refreshExpStr := os.Getenv("JWT_REFRESH_EXPIRATION")

	accessExp, err := utils.ParseDuration(accessExpStr)
	if err != nil {
		log.Fatalf("Invalid JWT_ACCESS_EXPIRATION: %v", err)
	}
	refreshExp, err := utils.ParseDuration(refreshExpStr)
	if err != nil {
		log.Fatalf("Invalid JWT_REFRESH_EXPIRATION: %v", err)
	}

	// Build DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect database: %v", err)
	}

	// Run AutoMigrate
	if err := db.AutoMigrate(
		&model.User{},
		&model.RefreshToken{},
	); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	fmt.Println("✅ Database migrated successfully")

	// Initialize router and services
	r := gin.Default()

	jwtManager := utils.NewJWTManager(jwtAccessSecret, jwtRefreshSecret)
	jwtManager.AccessTokenTTL = accessExp
	jwtManager.RefreshTokenTTL = refreshExp

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	authService := service.NewAuthService(userRepo, refreshTokenRepo, jwtManager)

	// main.go
	api := r.Group("/api")
	{
		http.RegisterAuthRoutes(api, userService, authService)
	}

	// Swagger docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Printf("🚀 Dev server started at http://localhost:%s\n", port)
	fmt.Printf("📘 Swagger docs at http://localhost:%s/docs/index.html\n", port)

	r.Run(":" + port)
}
