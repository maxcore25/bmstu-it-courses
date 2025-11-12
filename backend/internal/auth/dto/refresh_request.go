package dto

// RefreshRequest represents a refresh token request.
// @Description Refresh token request payload
// @Name RefreshRequest
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"dGhpc2lzYXJlZnJlc2h0b2tlbg==..."`
}
