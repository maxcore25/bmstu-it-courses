package mapper

import (
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/model"
)

// NewBranchResponse maps a Branch model to a BranchResponse DTO.
func NewBranchResponse(b *model.Branch) *dto.BranchResponse {
	return &dto.BranchResponse{
		ID:        b.ID,
		Address:   b.Address,
		Rooms:     b.Rooms,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}
