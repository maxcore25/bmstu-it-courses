package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/repository"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequest) (*model.User, error)
	GetUser(id uuid.UUID) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(req dto.CreateUserRequest) (*model.User, error) {
	// TODO: validate if email already exists before creation
	// TODO: hash password instead of storing plain string

	var middleNamePtr *string
	if req.MiddleName != "" {
		middleNamePtr = &req.MiddleName
	}

	var phonePtr *string
	if req.Phone != "" {
		phonePtr = &req.Phone
	}

	user := &model.User{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		MiddleName:     middleNamePtr,
		Email:          req.Email,
		Password:       req.Password, // TODO: hash before storing
		Phone:          phonePtr,
		KnowledgeLevel: model.KnowledgeLevel(req.KnowledgeLevel),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUser(id uuid.UUID) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
