package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
	httphelper "github.com/maxcore25/bmstu-it-courses/backend/internal/shared/http"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

// Register godoc
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body dto.RegisterRequest true "User registration data"
// @Success 200 {object} dto.AuthTokens
// @Failure 400 {object} gin.H
// @Failure 409 {object} gin.H
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		httphelper.JSONError(c, http.StatusBadRequest, err)
		return
	}

	tokens, err := h.service.Register(req)
	if err != nil {
		if err.Error() == "user with this email already exists" {
			httphelper.JSONError(c, http.StatusConflict, err)
		} else {
			httphelper.JSONError(c, http.StatusBadRequest, err)
		}
		return
	}

	httphelper.SetRefreshTokenCookie(c, tokens.RefreshToken)
	httphelper.JSONAccessToken(c, tokens.AccessToken)
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
		httphelper.JSONError(c, http.StatusBadRequest, err)
		return
	}

	tokens, err := h.service.Login(req)
	if err != nil {
		httphelper.JSONError(c, http.StatusUnauthorized, err)
		return
	}

	httphelper.SetRefreshTokenCookie(c, tokens.RefreshToken)
	httphelper.JSONAccessToken(c, tokens.AccessToken)
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
		httphelper.JSONError(c, http.StatusBadRequest, err)
		return
	}

	tokens, err := h.service.Refresh(req.RefreshToken)
	if err != nil {
		httphelper.JSONError(c, http.StatusUnauthorized, err)
		return
	}

	httphelper.SetRefreshTokenCookie(c, tokens.RefreshToken)
	httphelper.JSONAccessToken(c, tokens.AccessToken)
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
		httphelper.JSONError(c, http.StatusBadRequest, err)
		return
	}

	_ = h.service.Logout(req.RefreshToken)
	httphelper.ClearRefreshTokenCookie(c)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
