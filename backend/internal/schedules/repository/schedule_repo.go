package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/model"
	"gorm.io/gorm"
)

func applyExpansions(db *gorm.DB, expand map[string]bool) *gorm.DB {
	if expand["course"] {
		db = db.Preload("Course")
	}
	if expand["branch"] {
		db = db.Preload("Branch")
	}
	return db
}

type ScheduleRepository interface {
	Create(schedule *model.Schedule) error
	GetByID(id uuid.UUID) (*model.Schedule, error)
	GetAll() ([]*model.Schedule, error)
	GetByIDWithExpand(id uuid.UUID, expand map[string]bool) (*model.Schedule, error)
	GetAllWithExpand(expand map[string]bool) ([]*model.Schedule, error)
	UpdateByID(id uuid.UUID, updateData map[string]any) error
	DeleteByID(id uuid.UUID) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

func (r *scheduleRepository) Create(schedule *model.Schedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepository) GetByIDWithExpand(id uuid.UUID, expand map[string]bool) (*model.Schedule, error) {
	var s model.Schedule

	db := applyExpansions(r.db, expand)

	if err := db.First(&s, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *scheduleRepository) GetAllWithExpand(expand map[string]bool) ([]*model.Schedule, error) {
	var schedules []*model.Schedule

	db := applyExpansions(r.db, expand)

	if err := db.Find(&schedules).Error; err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *scheduleRepository) GetByID(id uuid.UUID) (*model.Schedule, error) {
	var s model.Schedule
	if err := r.db.First(&s, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *scheduleRepository) GetAll() ([]*model.Schedule, error) {
	var schedules []*model.Schedule
	if err := r.db.Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *scheduleRepository) UpdateByID(id uuid.UUID, updateData map[string]any) error {
	return r.db.Model(&model.Schedule{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *scheduleRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.Schedule{}).Error
}
