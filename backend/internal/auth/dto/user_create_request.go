package dto

// CreateUserRequest represents a user registration request.
// @Description User registration request payload
// @Name CreateUserRequest
type CreateUserRequest struct {
	FirstName      string `json:"first_name" binding:"required,min=2,max=50" example:"Иван"`
	LastName       string `json:"last_name" binding:"required,min=2,max=50" example:"Иванов"`
	MiddleName     string `json:"middle_name,omitempty" example:"Иванович"`
	Email          string `json:"email" binding:"required,email" example:"user@mail.ru"`
	Password       string `json:"password" binding:"required,min=6" example:"qwe123"`
	Phone          string `json:"phone,omitempty" example:"+77010000000"`
	KnowledgeLevel string `json:"knowledge_level" binding:"required,oneof=beginner intermediate advanced" example:"beginner"`
}
