package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// CreateUser godoc
// @Summary Create user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "New user"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := dto.UserResponse{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		MiddleName:     user.MiddleName,
		Email:          user.Email,
		Phone:          user.Phone,
		KnowledgeLevel: string(user.KnowledgeLevel),
	}

	c.JSON(http.StatusCreated, resp)
}

// GetUser godoc
// @Summary Get user by ID
// @Tags Users
// @Produce json
// @Param id path string true "User ID (uuid)"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} gin.H
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	resp := dto.UserResponse{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		MiddleName:     user.MiddleName,
		Email:          user.Email,
		Phone:          user.Phone,
		KnowledgeLevel: string(user.KnowledgeLevel),
	}

	c.JSON(http.StatusOK, resp)
}
