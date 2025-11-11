package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTManager struct {
	AccessSecret  string
	RefreshSecret string
}

type AccessTokenClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWTManager(accessSecret, refreshSecret string) *JWTManager {
	return &JWTManager{
		AccessSecret:  accessSecret,
		RefreshSecret: refreshSecret,
	}
}

func (j *JWTManager) GenerateAccessToken(userID uuid.UUID) (string, error) {
	claims := &AccessTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.AccessSecret))
}

func (j *JWTManager) GenerateRefreshToken(userID uuid.UUID) (string, error) {
	claims := &RefreshTokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.RefreshSecret))
}

func (j *JWTManager) VerifyAccessToken(tokenString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.AccessSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(*AccessTokenClaims), nil
}

func (j *JWTManager) VerifyRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.RefreshSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return token.Claims.(*RefreshTokenClaims), nil
}
