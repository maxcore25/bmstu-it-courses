package dto

import "github.com/google/uuid"

// UserResponse represents the structure of a user in responses.
// @Description User response payload
// @Name UserResponse
type UserResponse struct {
	ID             uuid.UUID `json:"id" example:"a8098c1a-f86e-11da-bd1a-00112444be1e"`
	FirstName      string    `json:"first_name" example:"John"`
	LastName       string    `json:"last_name" example:"Doe"`
	MiddleName     *string   `json:"middle_name,omitempty" example:"Michael"`
	Email          string    `json:"email" example:"user@mail.ru"`
	Phone          *string   `json:"phone,omitempty" example:"+77010000000"`
	KnowledgeLevel string    `json:"knowledge_level" example:"beginner"`
}
