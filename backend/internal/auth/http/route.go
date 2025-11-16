package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/middleware"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

func RegisterAuthRoutes(r *gin.RouterGroup, userService service.UserService, authService service.AuthService, jwtManager *utils.JWTManager) {
	userHandler := NewUserHandler(userService)
	authHandler := NewAuthHandler(authService)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/refresh", authHandler.Refresh)
		authGroup.POST("/logout", authHandler.Logout)
	}

	userGroup := r.Group("/users")
	{
		userGroup.POST("", middleware.AuthMiddleware(jwtManager), userHandler.CreateUser)
		userGroup.GET("", userHandler.GetAllUsers)
		userGroup.GET("/me", middleware.AuthMiddleware(jwtManager), userHandler.GetCurrentUser)
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.PATCH("/:id", middleware.AuthMiddleware(jwtManager), userHandler.UpdateUserByID)
		userGroup.DELETE("/:id", middleware.AuthMiddleware(jwtManager), userHandler.DeleteUserByID)
	}
}
