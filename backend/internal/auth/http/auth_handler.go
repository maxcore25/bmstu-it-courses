package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

// Login godoc
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "User credentials"
// @Success 200 {object} dto.AuthTokens
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// Refresh godoc
// @Summary Refresh access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param refresh body dto.RefreshRequest true "Refresh token request"
// @Success 200 {object} dto.AuthTokens
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var req dto.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.service.Refresh(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// Logout godoc
// @Summary Logout user (invalidate refresh token)
// @Tags Auth
// @Accept json
// @Produce json
// @Param logout body dto.RefreshRequest true "Refresh token request to logout"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	var req dto.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_ = h.service.Logout(req.RefreshToken)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
