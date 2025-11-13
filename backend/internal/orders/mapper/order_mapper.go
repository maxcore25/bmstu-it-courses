package mapper

import (
	"time"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/model"
)

// NewOrderResponse creates an OrderResponse from model.Order.
func NewOrderResponse(o *model.Order) *dto.OrderResponse {
	return &dto.OrderResponse{
		ID:         o.ID,
		ClientID:   o.ClientID,
		CourseID:   o.CourseID,
		ScheduleID: o.ScheduleID,
		BranchID:   o.BranchID,
		Price:      o.Price,
		CreatedAt:  o.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  o.UpdatedAt.Format(time.RFC3339),
	}
}
