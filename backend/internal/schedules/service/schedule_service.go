package service

import (
	"github.com/google/uuid"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/repository"
)

type ScheduleService interface {
	GetSchedule(id uuid.UUID) (*model.Schedule, error)
	CreateSchedule(req *dto.CreateScheduleRequest) (*model.Schedule, error)
	GetAllSchedules() ([]*model.Schedule, error)
	UpdateScheduleByID(id uuid.UUID, updates map[string]any) error
	DeleteScheduleByID(id uuid.UUID) error
}

type scheduleService struct {
	repo repository.ScheduleRepository
}

func NewScheduleService(r repository.ScheduleRepository) ScheduleService {
	return &scheduleService{repo: r}
}

func (s *scheduleService) GetSchedule(id uuid.UUID) (*model.Schedule, error) {
	return s.repo.GetByID(id)
}

func (s *scheduleService) CreateSchedule(req *dto.CreateScheduleRequest) (*model.Schedule, error) {
	schedule := &model.Schedule{
		CourseID: req.CourseID,
		BranchID: req.BranchID,
		StartAt:  req.StartAt,
		EndAt:    req.EndAt,
		Capacity: req.Capacity,
	}
	err := s.repo.Create(schedule)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

func (s *scheduleService) GetAllSchedules() ([]*model.Schedule, error) {
	return s.repo.GetAll()
}

func (s *scheduleService) UpdateScheduleByID(id uuid.UUID, updates map[string]any) error {
	return s.repo.UpdateByID(id, updates)
}

func (s *scheduleService) DeleteScheduleByID(id uuid.UUID) error {
	return s.repo.DeleteByID(id)
}
