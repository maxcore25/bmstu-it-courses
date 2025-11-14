package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/model"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *model.Course) error
	GetByID(id uuid.UUID, preloadAuthor bool) (*model.Course, error)
	GetAll(preloadAuthor bool) ([]*model.Course, error)
	UpdateByID(id uuid.UUID, updateData map[string]any) error
	DeleteByID(id uuid.UUID) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(course *model.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) GetByID(id uuid.UUID, preloadAuthor bool) (*model.Course, error) {
	query := r.db
	if preloadAuthor {
		query = query.Preload("Author")
	}

	var c model.Course
	if err := query.First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *courseRepository) GetAll(preloadAuthor bool) ([]*model.Course, error) {
	query := r.db
	if preloadAuthor {
		query = query.Preload("Author")
	}

	var courses []*model.Course
	if err := query.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *courseRepository) UpdateByID(id uuid.UUID, updateData map[string]any) error {
	return r.db.Model(&model.Course{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *courseRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Where("id = ?", id).Delete(&model.Course{}).Error
}
