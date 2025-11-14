package mapper

import (
	"time"

	"github.com/google/uuid"

	userMapper "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/mapper"
	branchMapper "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/mapper"
	courseMapper "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/mapper"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/model"
)

// NewOrderResponse creates an OrderResponse from model.Order, including expanded relations if preloaded.
func NewOrderResponse(o *model.Order) *dto.OrderResponse {
	resp := &dto.OrderResponse{
		ID:         o.ID,
		ClientID:   o.ClientID,
		CourseID:   o.CourseID,
		ScheduleID: o.ScheduleID,
		BranchID:   o.BranchID,
		Price:      o.Price,
		CreatedAt:  o.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  o.UpdatedAt.Format(time.RFC3339),
	}

	if o.Client.ID != uuid.Nil {
		resp.Client = userMapper.NewUserResponse(&o.Client)
	}

	if o.Course.ID != uuid.Nil {
		resp.Course = courseMapper.NewCourseResponse(&o.Course)
	}

	if o.Branch != nil && o.Branch.ID != uuid.Nil {
		resp.Branch = branchMapper.NewBranchResponse(o.Branch)
	}

	return resp
}
