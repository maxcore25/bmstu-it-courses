package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/model"
	"gorm.io/gorm"
)

func applyExpansions(db *gorm.DB, expand map[string]bool) *gorm.DB {
	if expand["client"] {
		db = db.Preload("Client")
	}
	if expand["course"] {
		db = db.Preload("Course")
	}
	if expand["branch"] {
		db = db.Preload("Branch")
	}
	return db
}

type OrderRepository interface {
	Create(order *model.Order) error
	GetByID(id uuid.UUID) (*model.Order, error)
	GetByUserWithExpand(userID uuid.UUID, expand map[string]bool) ([]*model.Order, error)
	GetByUser(userID uuid.UUID) ([]*model.Order, error)
	GetAll() ([]*model.Order, error)
	GetByIDWithExpand(id uuid.UUID, expand map[string]bool) (*model.Order, error)
	GetAllWithExpand(expand map[string]bool) ([]*model.Order, error)
	UpdateByID(id uuid.UUID, updateData map[string]any) error
	DeleteByID(id uuid.UUID) error
	DB() *gorm.DB
}

type orderRepository struct {
	db *gorm.DB
}

func (r *orderRepository) DB() *gorm.DB {
	return r.db
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) GetByUserWithExpand(userID uuid.UUID, expand map[string]bool) ([]*model.Order, error) {
	var orders []*model.Order

	db := applyExpansions(r.db, expand)

	if err := db.Where("client_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) GetByUser(userID uuid.UUID) ([]*model.Order, error) {
	var orders []*model.Order

	if err := r.db.Where("client_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) GetByIDWithExpand(id uuid.UUID, expand map[string]bool) (*model.Order, error) {
	var s model.Order

	db := applyExpansions(r.db, expand)

	if err := db.First(&s, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *orderRepository) GetAllWithExpand(expand map[string]bool) ([]*model.Order, error) {
	var orders []*model.Order

	db := applyExpansions(r.db, expand)

	if err := db.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *orderRepository) GetByID(id uuid.UUID) (*model.Order, error) {
	var o model.Order
	if err := r.db.First(&o, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *orderRepository) GetAll() ([]*model.Order, error) {
	var orders []*model.Order
	if err := r.db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) UpdateByID(id uuid.UUID, updateData map[string]any) error {
	return r.db.Model(&model.Order{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *orderRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.Order{}).Error
}
