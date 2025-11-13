package dto

import (
	"time"

	"github.com/google/uuid"
)

// BranchResponse represents the structure of a branch returned in API responses.
// @Description Branch response payload.
// @Name BranchResponse
type BranchResponse struct {
	ID        uuid.UUID `json:"id" example:"1cf1b1f6-fbc2-4b5e-bd99-c71b4f9f67a2"`
	Address   string    `json:"address" example:"Москва, Тверская ул., д. 1, к. 1"`
	Rooms     int       `json:"rooms" example:"8"`
	CreatedAt time.Time `json:"createdAt" example:"2025-11-12T19:45:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2025-11-12T19:45:00Z"`
}
