package repository

import (
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/model"
	"gorm.io/gorm"
)

func applyExpansions(db *gorm.DB, expand map[string]bool) *gorm.DB {
	if expand["author"] {
		db = db.Preload("Author")
	}
	return db
}

type CourseRepository interface {
	Create(course *model.Course) error
	GetByID(id uuid.UUID) (*model.Course, error)
	GetAll() ([]*model.Course, error)
	GetByIDWithExpand(id uuid.UUID, expand map[string]bool) (*model.Course, error)
	GetAllWithExpand(expand map[string]bool) ([]*model.Course, error)
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

func (r *courseRepository) GetByIDWithExpand(id uuid.UUID, expand map[string]bool) (*model.Course, error) {
	var s model.Course

	db := applyExpansions(r.db, expand)

	if err := db.First(&s, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *courseRepository) GetAllWithExpand(expand map[string]bool) ([]*model.Course, error) {
	var courses []*model.Course

	db := applyExpansions(r.db, expand)

	if err := db.Find(&courses).Error; err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *courseRepository) GetByID(id uuid.UUID) (*model.Course, error) {
	var c model.Course
	if err := r.db.First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *courseRepository) GetAll() ([]*model.Course, error) {
	var courses []*model.Course
	if err := r.db.Find(&courses).Error; err != nil {
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
