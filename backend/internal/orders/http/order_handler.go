package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/mapper"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/service"
	httphelper "github.com/maxcore25/bmstu-it-courses/backend/internal/shared/http"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandler(s service.OrderService) *OrderHandler {
	return &OrderHandler{service: s}
}

// CreateOrder godoc
// @Summary Create order
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body dto.CreateOrderRequest true "New order"
// @Param price query integer true "Order price (at time of order)"
// @Success 201 {object} dto.OrderResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest

	if !httphelper.BindJSON(c, &req) {
		return
	}

	order, err := h.service.CreateOrder(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.NewOrderResponse(order)

	c.JSON(http.StatusCreated, resp)
}

// GetMyOrders godoc
// @Summary Get orders of the authenticated user
// @Tags Orders
// @Produce json
// @Param expand query []string false "Relations to expand (client, course, branch). Example: expand=client,course"
// @Success 200 {array} dto.OrderResponse
// @Failure 401 {object} gin.H
// @Router /orders/me [get]
func (h *OrderHandler) GetMyOrders(c *gin.Context) {
	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID := uid.(uuid.UUID)

	expand := httphelper.ParseExpand(c.QueryArray("expand"))

	orders, err := h.service.GetOrdersByUser(userID, expand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]*dto.OrderResponse, len(orders))
	for i, order := range orders {
		resp[i] = mapper.NewOrderResponse(order)
	}

	c.JSON(http.StatusOK, resp)
}

// GetOrder godoc
// @Summary Get order by ID
// @Tags Orders
// @Produce json
// @Param id path string true "Order ID (uuid)"
// @Param expand query []string false "Relations to expand (client, course, branch). Example: expand=client,course,branch"
// @Success 200 {object} dto.OrderResponse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	expand := httphelper.ParseExpand(c.QueryArray("expand"))

	order, err := h.service.GetOrder(id, expand)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	resp := mapper.NewOrderResponse(order)

	c.JSON(http.StatusOK, resp)
}

// GetAllOrders godoc
// @Summary Get all orders
// @Tags Orders
// @Produce json
// @Param expand query []string false "Relations to expand (client, course, branch). Example: expand=client,course,branch"
// @Success 200 {array} dto.OrderResponse
// @Router /orders [get]
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	expand := httphelper.ParseExpand(c.QueryArray("expand"))

	orders, err := h.service.GetAllOrders(expand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]*dto.OrderResponse, len(orders))
	for i, order := range orders {
		resp[i] = mapper.NewOrderResponse(order)
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateOrderByID godoc
// @Summary Update order by ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID (uuid)"
// @Param order body map[string]interface{} true "Order update data"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/{id} [patch]
func (h *OrderHandler) UpdateOrderByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	var updateData map[string]any
	if !httphelper.BindJSON(c, &updateData) {
		return
	}
	if err := h.service.UpdateOrderByID(id, updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "order updated successfully"})
}

// DeleteOrderByID godoc
// @Summary Delete order by ID
// @Tags Orders
// @Produce json
// @Param id path string true "Order ID (uuid)"
// @Success 204 {object} nil
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /orders/{id} [delete]
func (h *OrderHandler) DeleteOrderByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	if err := h.service.DeleteOrderByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
