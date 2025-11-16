package dto

// UpdateBranchRequest represents a branch update request.
// @Description Branch update request payload
// @Name UpdateBranchRequest
type UpdateBranchRequest struct {
	Address *string `json:"address,omitempty" binding:"omitempty,min=2,max=255" example:"Москва, Тверская ул., д. 1, к. 1"`
	Rooms   *int    `json:"rooms,omitempty" binding:"omitempty,min=1" example:"8"`
}
