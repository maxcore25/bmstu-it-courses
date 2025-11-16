package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/model"
	"gorm.io/gorm"
)

type BranchRepository interface {
	Create(branch *model.Branch) error
	GetByID(id uuid.UUID) (*model.Branch, error)
	GetAll() ([]*model.Branch, error)
	UpdateByID(id uuid.UUID, updateData dto.UpdateBranchRequest) error
	DeleteByID(id uuid.UUID) error
}

type branchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) BranchRepository {
	return &branchRepository{db: db}
}

func (r *branchRepository) Create(branch *model.Branch) error {
	return r.db.Create(branch).Error
}

func (r *branchRepository) GetByID(id uuid.UUID) (*model.Branch, error) {
	var b model.Branch
	if err := r.db.First(&b, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *branchRepository) GetAll() ([]*model.Branch, error) {
	var branches []*model.Branch
	if err := r.db.Find(&branches).Error; err != nil {
		return nil, err
	}
	return branches, nil
}

func (r *branchRepository) UpdateByID(id uuid.UUID, updateData dto.UpdateBranchRequest) error {
	return r.db.Model(&model.Branch{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *branchRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.Branch{}).Error
}
