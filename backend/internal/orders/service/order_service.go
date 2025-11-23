package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	courseModel "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/repository"
	scheduleModel "github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/model"
)

type OrderService interface {
	CreateOrder(req *dto.CreateOrderRequest) (*model.Order, error)
	GetOrdersByUser(userID uuid.UUID, expand map[string]bool) ([]*model.Order, error)
	GetOrder(id uuid.UUID, expand map[string]bool) (*model.Order, error)
	GetAllOrders(expand map[string]bool) ([]*model.Order, error)
	GetOrdersMetadata(userID uuid.UUID) (*dto.OrdersMetadata, error)
	UpdateOrderByID(id uuid.UUID, updates map[string]any) error
	DeleteOrderByID(id uuid.UUID) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(r repository.OrderRepository) OrderService {
	return &orderService{repo: r}
}

func (s *orderService) CreateOrder(req *dto.CreateOrderRequest) (*model.Order, error) {
	tx := s.repo.DB().Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// --- 1. Validate schedule and lock it (if provided)
	var schedule scheduleModel.Schedule
	if req.ScheduleID != nil {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", *req.ScheduleID).
			First(&schedule).Error; err != nil {
			tx.Rollback()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("schedule not found")
			}
			return nil, err
		}

		// Check capacity
		if schedule.Reserved >= schedule.Capacity {
			tx.Rollback()
			return nil, fmt.Errorf("no seats available for this schedule")
		}
	}

	// --- 2. Get course for price snapshot
	var course courseModel.Course
	if err := tx.First(&course, "id = ?", req.CourseID).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("course not found")
		}
		return nil, err
	}

	price := course.Price // snapshot current price

	// --- 3. Create order entity
	order := &model.Order{
		ClientID:   req.ClientID,
		CourseID:   req.CourseID,
		ScheduleID: req.ScheduleID,
		BranchID:   req.BranchID,
		Price:      price,
	}

	// --- 4. Save order
	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	// --- 5. Increment reserved seats if schedule exists
	if req.ScheduleID != nil {
		if err := tx.Model(&schedule).
			UpdateColumn("reserved", gorm.Expr("reserved + ?", 1)).Error; err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to increment reserved seats: %w", err)
		}
	}

	// --- 6. Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (s *orderService) GetOrdersMetadata(userID uuid.UUID) (*dto.OrdersMetadata, error) {
	return s.repo.GetOrdersMetadata(userID)
}

func (s *orderService) GetOrdersByUser(userID uuid.UUID, expand map[string]bool) ([]*model.Order, error) {
	if len(expand) > 0 {
		return s.repo.GetByUserWithExpand(userID, expand)
	}
	return s.repo.GetByUser(userID)
}

func (s *orderService) GetOrder(id uuid.UUID, expand map[string]bool) (*model.Order, error) {
	if len(expand) > 0 {
		return s.repo.GetByIDWithExpand(id, expand)
	}
	return s.repo.GetByID(id)
}

func (s *orderService) GetAllOrders(expand map[string]bool) ([]*model.Order, error) {
	if len(expand) > 0 {
		return s.repo.GetAllWithExpand(expand)
	}
	return s.repo.GetAll()
}

func (s *orderService) UpdateOrderByID(id uuid.UUID, updates map[string]any) error {
	return s.repo.UpdateByID(id, updates)
}

func (s *orderService) DeleteOrderByID(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}
