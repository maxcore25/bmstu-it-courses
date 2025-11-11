package service

import (
	"errors"
	"time"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/repository"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

type AuthService interface {
	Login(req dto.LoginRequest) (*dto.AuthTokens, error)
	Refresh(refreshToken string) (*dto.AuthTokens, error)
	Logout(refreshToken string) error
}

type authService struct {
	userRepo    repository.UserRepository
	refreshRepo repository.RefreshTokenRepository
	jwt         *utils.JWTManager
}

func NewAuthService(u repository.UserRepository, r repository.RefreshTokenRepository, j *utils.JWTManager) AuthService {
	return &authService{userRepo: u, refreshRepo: r, jwt: j}
}

func (s *authService) Login(req dto.LoginRequest) (*dto.AuthTokens, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	access, err := s.jwt.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refresh, err := s.jwt.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, err
	}

	tokenModel := &model.RefreshToken{
		UserID:    user.ID,
		Token:     refresh,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.refreshRepo.Save(tokenModel); err != nil {
		return nil, err
	}

	return &dto.AuthTokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *authService) Refresh(refreshToken string) (*dto.AuthTokens, error) {
	claims, err := s.jwt.VerifyRefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	access, err := s.jwt.GenerateAccessToken(claims.UserID)
	if err != nil {
		return nil, err
	}

	newRefresh, err := s.jwt.GenerateRefreshToken(claims.UserID)
	if err != nil {
		return nil, err
	}

	_ = s.refreshRepo.Delete(refreshToken)

	tokenModel := &model.RefreshToken{
		UserID:    claims.UserID,
		Token:     newRefresh,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.refreshRepo.Save(tokenModel); err != nil {
		return nil, err
	}

	return &dto.AuthTokens{
		AccessToken:  access,
		RefreshToken: newRefresh,
	}, nil
}

func (s *authService) Logout(refreshToken string) error {
	return s.refreshRepo.Delete(refreshToken)
}
