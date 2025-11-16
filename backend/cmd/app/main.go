package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	authHttp "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/http"
	authModel "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
	authRepo "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/repository"
	authService "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/seed"

	branchHttp "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/http"
	branchModel "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/model"
	branchRepo "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/repository"
	branchService "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/service"

	courseHttp "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/http"
	courseModel "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/model"
	courseRepo "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/repository"
	courseService "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/service"

	scheduleHttp "github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/http"
	scheduleModel "github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/model"
	scheduleRepo "github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/repository"
	scheduleService "github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/service"

	orderHttp "github.com/maxcore25/bmstu-it-courses/backend/internal/orders/http"
	orderModel "github.com/maxcore25/bmstu-it-courses/backend/internal/orders/model"
	orderRepo "github.com/maxcore25/bmstu-it-courses/backend/internal/orders/repository"
	orderService "github.com/maxcore25/bmstu-it-courses/backend/internal/orders/service"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/config"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/maxcore25/bmstu-it-courses/backend/docs"
)

// @title CodeCraft - IT Courses School API
// @version 1.0
// @description API documentation for CodeCraft IT Courses School.
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.type http
// @securityDefinitions.scheme bearer
// @securityDefinitions.bearerFormat JWT
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer {your_token}" to authorize

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
		&authModel.User{},
		&authModel.RefreshToken{},
		&branchModel.Branch{},
		&courseModel.Course{},
		&scheduleModel.Schedule{},
		&orderModel.Order{},
	); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	fmt.Println("✅ Database migrated successfully")

	// Initialize router and services
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost") || strings.HasPrefix(origin, "http://127.0.0.1")
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	jwtManager := utils.NewJWTManager(jwtAccessSecret, jwtRefreshSecret)
	jwtManager.AccessTokenTTL = accessExp
	jwtManager.RefreshTokenTTL = refreshExp

	userRepo := authRepo.NewUserRepository(db)
	userService := authService.NewUserService(userRepo)
	refreshTokenRepo := authRepo.NewRefreshTokenRepository(db)
	authService := authService.NewAuthService(userRepo, refreshTokenRepo, jwtManager)

	branchRepo := branchRepo.NewBranchRepository(db)
	branchService := branchService.NewBranchService(branchRepo)

	courseRepo := courseRepo.NewCourseRepository(db)
	courseService := courseService.NewCourseService(courseRepo)

	scheduleRepo := scheduleRepo.NewScheduleRepository(db)
	scheduleService := scheduleService.NewScheduleService(scheduleRepo)

	orderRepo := orderRepo.NewOrderRepository(db)
	orderService := orderService.NewOrderService(orderRepo)

	// Run seeds
	if err := seed.RunSeeds(db); err != nil {
		log.Fatalf("❌ Failed to seed: %v", err)
		panic(err)
	}

	// Register routes
	api := r.Group("/api")
	{
		authHttp.RegisterAuthRoutes(api, userService, authService, jwtManager)
		branchHttp.RegisterBranchRoutes(api, branchService, jwtManager)
		courseHttp.RegisterCourseRoutes(api, courseService, jwtManager)
		scheduleHttp.RegisterScheduleRoutes(api, scheduleService, jwtManager)
		orderHttp.RegisterOrderRoutes(api, orderService, jwtManager)
	}

	// Swagger docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.PersistAuthorization(true),
	))

	fmt.Printf("🚀 Dev server started at http://localhost:%s\n", port)
	fmt.Printf("📘 Swagger docs at http://localhost:%s/docs/index.html\n", port)

	r.Run(":" + port)
}
