package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id uuid.UUID) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uuid.UUID) (*model.User, error) {
	var u model.User
	if err := r.db.First(&u, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var u model.User
	if err := r.db.First(&u, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
