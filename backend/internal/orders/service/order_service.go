package service

import (
	"github.com/google/uuid"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/repository"
)

type OrderService interface {
	GetOrder(id uuid.UUID) (*model.Order, error)
	CreateOrder(req *dto.CreateOrderRequest, price int64) (*model.Order, error)
	GetAllOrders() ([]*model.Order, error)
	UpdateOrderByID(id uuid.UUID, updates map[string]any) error
	DeleteOrderByID(id uuid.UUID) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(r repository.OrderRepository) OrderService {
	return &orderService{repo: r}
}

func (s *orderService) GetOrder(id uuid.UUID) (*model.Order, error) {
	return s.repo.GetByID(id)
}

// price is provided externally and should be calculated beforehand
func (s *orderService) CreateOrder(req *dto.CreateOrderRequest, price int64) (*model.Order, error) {
	order := &model.Order{
		ClientID:   req.ClientID,
		CourseID:   req.CourseID,
		ScheduleID: req.ScheduleID,
		BranchID:   req.BranchID,
		Price:      price,
	}
	err := s.repo.Create(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *orderService) GetAllOrders() ([]*model.Order, error) {
	return s.repo.GetAll()
}

func (s *orderService) UpdateOrderByID(id uuid.UUID, updates map[string]any) error {
	return s.repo.UpdateByID(id, updates)
}

func (s *orderService) DeleteOrderByID(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}
