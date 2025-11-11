package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

func RegisterAuthRoutes(r *gin.Engine, userService service.UserService, authService service.AuthService, jwtManager *utils.JWTManager) {
	userHandler := NewUserHandler(userService)
	authHandler := NewAuthHandler(authService)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/refresh", authHandler.Refresh)
		authGroup.POST("/logout", authHandler.Logout)
	}

	userGroup := r.Group("/users")
	{
		userGroup.POST("", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUser)
	}
}
