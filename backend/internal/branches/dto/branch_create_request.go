package dto

// CreateBranchRequest represents a branch creation request.
// @Description Branch creation request payload
// @Name CreateBranchRequest
type CreateBranchRequest struct {
	Address string `json:"address" binding:"required,min=2,max=255" example:"Москва, Тверская ул., д. 1, к. 1"`
	Rooms   int    `json:"rooms" binding:"required,min=1" example:"8"`
}
