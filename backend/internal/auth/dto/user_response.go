package dto

import "github.com/google/uuid"

type UserResponse struct {
	ID             uuid.UUID `json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	MiddleName     *string   `json:"middle_name,omitempty"`
	Email          string    `json:"email"`
	Phone          *string   `json:"phone,omitempty"`
	KnowledgeLevel string    `json:"knowledge_level"`
}
