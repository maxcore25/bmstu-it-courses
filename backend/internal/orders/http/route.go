package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/middleware"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

func RegisterOrderRoutes(r *gin.RouterGroup, orderService service.OrderService, jwtManager *utils.JWTManager) {
	orderHandler := NewOrderHandler(orderService)

	orderGroup := r.Group("/orders")

	protected := orderGroup.Group("")
	protected.Use(middleware.AuthMiddleware(jwtManager))
	{
		protected.POST("", orderHandler.CreateOrder)
		protected.GET("/total", orderHandler.GetOrdersMetadata)
		protected.GET("", orderHandler.GetAllOrders)
		protected.GET("/my", orderHandler.GetMyOrders)
		protected.GET("/:id", orderHandler.GetOrder)
		protected.PATCH("/:id", orderHandler.UpdateOrderByID)
		protected.DELETE("/:id", orderHandler.DeleteOrderByID)
	}
}
