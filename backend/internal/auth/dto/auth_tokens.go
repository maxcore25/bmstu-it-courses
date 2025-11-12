package dto

// AuthTokens represents authentication tokens.
// @Description Authentication tokens response payload
// @Name AuthTokens
type AuthTokens struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string `json:"refresh_token" example:"dGhpc2lzYXJlZnJlc2h0b2tlbg==..."`
}
