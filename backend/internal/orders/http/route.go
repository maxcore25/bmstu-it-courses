package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/middleware"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

// RegisterOrderRoutes registers the order routes with the provided gin router group.
func RegisterOrderRoutes(r *gin.RouterGroup, orderService service.OrderService, jwtManager *utils.JWTManager) {
	orderHandler := NewOrderHandler(orderService)

	orderGroup := r.Group("/orders")
	{
		orderGroup.POST("", orderHandler.CreateOrder)
		orderGroup.GET("", orderHandler.GetAllOrders)
		orderGroup.GET("/my", middleware.AuthMiddleware(jwtManager), orderHandler.GetMyOrders)
		orderGroup.GET("/:id", orderHandler.GetOrder)
		orderGroup.PATCH("/:id", orderHandler.UpdateOrderByID)
		orderGroup.DELETE("/:id", orderHandler.DeleteOrderByID)
	}
}
