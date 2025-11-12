package dto

// CreateUserRequest represents a user registration request.
// @Description User registration request payload
// @Name CreateUserRequest
type CreateUserRequest struct {
	FirstName      string `json:"first_name" binding:"required,min=2,max=50" example:"John"`
	LastName       string `json:"last_name" binding:"required,min=2,max=50" example:"Doe"`
	MiddleName     string `json:"middle_name,omitempty" example:"Michael"`
	Email          string `json:"email" binding:"required,email" example:"john.doe@example.com"`
	Password       string `json:"password" binding:"required,min=6" example:"secret123"`
	Phone          string `json:"phone,omitempty" example:"+77010000000"`
	KnowledgeLevel string `json:"knowledge_level" binding:"required,oneof=beginner intermediate advanced" example:"beginner"`
}
