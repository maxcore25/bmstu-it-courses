package service

import (
	"github.com/google/uuid"

	authModel "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/repository"
)

type CourseService interface {
	CreateCourse(req *dto.CreateCourseRequest) (*model.Course, error)
	GetCourse(id uuid.UUID, includeAuthor bool) (*model.Course, error)
	GetAllCourses(includeAuthor bool) ([]*model.Course, error)
	UpdateCourseByID(id uuid.UUID, updates map[string]any) error
	DeleteCourseByID(id uuid.UUID) error
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(r repository.CourseRepository) CourseService {
	return &courseService{repo: r}
}

func (s *courseService) CreateCourse(req *dto.CreateCourseRequest) (*model.Course, error) {
	course := &model.Course{
		Name:       req.Name,
		Difficulty: authModel.KnowledgeLevel(req.Difficulty),
		Duration:   req.Duration,
		Price:      req.Price,
		Format:     model.CourseFormat(req.Format),
		AuthorID:   req.AuthorID,
	}
	err := s.repo.Create(course)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (s *courseService) GetCourse(id uuid.UUID, includeAuthor bool) (*model.Course, error) {
	return s.repo.GetByID(id, includeAuthor)
}

func (s *courseService) GetAllCourses(includeAuthor bool) ([]*model.Course, error) {
	return s.repo.GetAll(includeAuthor)
}

func (s *courseService) UpdateCourseByID(id uuid.UUID, updates map[string]any) error {
	return s.repo.UpdateByID(id, updates)
}

func (s *courseService) DeleteCourseByID(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}
