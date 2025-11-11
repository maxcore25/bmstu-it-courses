package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
)

func RegisterAuthRoutes(r *gin.Engine, us service.UserService) {
	userHandler := NewUserHandler(us)

	g := r.Group("/users")
	{
		g.POST("", userHandler.CreateUser)
		g.GET("/:id", userHandler.GetUser)
	}

	// TODO: add login, refresh, logout routes later
}
