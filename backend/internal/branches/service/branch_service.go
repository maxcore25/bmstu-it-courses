package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/repository"
)

type BranchService interface {
	CreateBranch(req dto.CreateBranchRequest) (*model.Branch, error)
	GetBranch(id uuid.UUID) (*model.Branch, error)
	GetAllBranches() ([]*model.Branch, error)
	UpdateBranchByID(id uuid.UUID, updateData dto.UpdateBranchRequest) error
	DeleteBranchByID(id uuid.UUID) error
}

type branchService struct {
	repo repository.BranchRepository
}

func NewBranchService(r repository.BranchRepository) BranchService {
	return &branchService{repo: r}
}

// Note: The DTO 'BranchResponse' is designed for responses, not for creation
// In a real-world scenario, you might want a separate 'CreateBranchRequest' DTO
func (s *branchService) CreateBranch(req dto.CreateBranchRequest) (*model.Branch, error) {
	branch := &model.Branch{
		Address: req.Address,
		Rooms:   req.Rooms,
	}

	if err := s.repo.Create(branch); err != nil {
		return nil, err
	}
	return branch, nil
}

func (s *branchService) GetBranch(id uuid.UUID) (*model.Branch, error) {
	branch, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("branch not found")
	}
	return branch, nil
}

func (s *branchService) GetAllBranches() ([]*model.Branch, error) {
	branches, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return branches, nil
}

func (s *branchService) UpdateBranchByID(id uuid.UUID, updateData dto.UpdateBranchRequest) error {
	// Could add validation here (e.g., allowed fields)
	return s.repo.UpdateByID(id, updateData)
}

func (s *branchService) DeleteBranchByID(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}
