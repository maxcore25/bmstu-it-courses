package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/model"
	schedules_service "github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/service"
)

type mockScheduleRepo struct {
	created     *model.Schedule
	errOnCreate error

	scheduleByID    *model.Schedule
	errOnGetByID    error
	scheduleWithExp *model.Schedule
	errOnGetByIDExp error

	allSchedules    []*model.Schedule
	errOnGetAll     error
	allSchedulesExp []*model.Schedule
	errOnGetAllExp  error

	errOnUpdate error
	updatedData map[string]any
	updatedID   uuid.UUID

	errOnDelete error
	deletedID   uuid.UUID
}

func (m *mockScheduleRepo) Create(schedule *model.Schedule) error {
	if m.errOnCreate != nil {
		return m.errOnCreate
	}
	schedule.ID = uuid.New()
	schedule.CreatedAt = time.Now()
	schedule.UpdatedAt = time.Now()
	m.created = schedule
	return nil
}

func (m *mockScheduleRepo) GetByID(id uuid.UUID) (*model.Schedule, error) {
	if m.errOnGetByID != nil {
		return nil, m.errOnGetByID
	}
	return m.scheduleByID, nil
}

func (m *mockScheduleRepo) GetByIDWithExpand(id uuid.UUID, expand map[string]bool) (*model.Schedule, error) {
	if m.errOnGetByIDExp != nil {
		return nil, m.errOnGetByIDExp
	}
	return m.scheduleWithExp, nil
}

func (m *mockScheduleRepo) GetAll() ([]*model.Schedule, error) {
	if m.errOnGetAll != nil {
		return nil, m.errOnGetAll
	}
	return m.allSchedules, nil
}

func (m *mockScheduleRepo) GetAllWithExpand(expand map[string]bool) ([]*model.Schedule, error) {
	if m.errOnGetAllExp != nil {
		return nil, m.errOnGetAllExp
	}
	return m.allSchedulesExp, nil
}

func (m *mockScheduleRepo) UpdateByID(id uuid.UUID, updateData map[string]any) error {
	m.updatedID = id
	m.updatedData = updateData
	if m.errOnUpdate != nil {
		return m.errOnUpdate
	}
	return nil
}

func (m *mockScheduleRepo) DeleteByID(id uuid.UUID) error {
	m.deletedID = id
	if m.errOnDelete != nil {
		return m.errOnDelete
	}
	return nil
}

func TestCreateSchedule_Success(t *testing.T) {
	mockRepo := &mockScheduleRepo{}
	svc := schedules_service.NewScheduleService(mockRepo)

	req := &dto.CreateScheduleRequest{
		CourseID: uuid.New(),
		BranchID: &uuid.UUID{},
		StartAt:  time.Now().Add(time.Hour),
		EndAt:    time.Now().Add(2 * time.Hour),
		Capacity: 18,
	}
	s, err := svc.CreateSchedule(req)
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, req.CourseID, s.CourseID)
	require.Equal(t, req.BranchID, s.BranchID)
	require.Equal(t, req.Capacity, s.Capacity)
	require.Equal(t, req.StartAt, s.StartAt)
	require.Equal(t, req.EndAt, s.EndAt)
	require.NotEqual(t, uuid.Nil, s.ID)
}

func TestCreateSchedule_RepoError(t *testing.T) {
	mockRepo := &mockScheduleRepo{errOnCreate: errors.New("db fail")}
	svc := schedules_service.NewScheduleService(mockRepo)

	req := &dto.CreateScheduleRequest{
		CourseID: uuid.New(),
		BranchID: &uuid.UUID{},
		StartAt:  time.Now(),
		EndAt:    time.Now(),
		Capacity: 1,
	}
	s, err := svc.CreateSchedule(req)
	require.Error(t, err)
	require.Nil(t, s)
}

func TestGetScheduleByID_Success(t *testing.T) {
	scheduleID := uuid.New()
	expected := &model.Schedule{
		ID:        scheduleID,
		CourseID:  uuid.New(),
		BranchID:  &uuid.UUID{},
		StartAt:   time.Now(),
		EndAt:     time.Now().Add(time.Hour),
		Capacity:  15,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRepo := &mockScheduleRepo{scheduleByID: expected}
	svc := schedules_service.NewScheduleService(mockRepo)

	got, err := svc.GetSchedule(scheduleID, map[string]bool{})
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, expected.ID, got.ID)
	require.Equal(t, expected.CourseID, got.CourseID)
	require.Equal(t, expected.BranchID, got.BranchID)
	require.Equal(t, expected.Capacity, got.Capacity)
}

func TestGetScheduleByID_RepoError(t *testing.T) {
	mockRepo := &mockScheduleRepo{errOnGetByID: errors.New("not found")}
	svc := schedules_service.NewScheduleService(mockRepo)

	id := uuid.New()
	got, err := svc.GetSchedule(id, map[string]bool{})
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGetScheduleByID_WithExpand_Success(t *testing.T) {
	scheduleID := uuid.New()
	expected := &model.Schedule{ID: scheduleID, CourseID: uuid.New(), BranchID: &uuid.UUID{}}
	mockRepo := &mockScheduleRepo{scheduleWithExp: expected}
	svc := schedules_service.NewScheduleService(mockRepo)

	expand := map[string]bool{"course": true, "branch": true}
	got, err := svc.GetSchedule(scheduleID, expand)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, expected.ID, got.ID)
}

func TestGetScheduleByID_WithExpand_RepoError(t *testing.T) {
	mockRepo := &mockScheduleRepo{errOnGetByIDExp: errors.New("exp fail")}
	svc := schedules_service.NewScheduleService(mockRepo)

	id := uuid.New()
	got, err := svc.GetSchedule(id, map[string]bool{"course": true})
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGetAllSchedules_Success(t *testing.T) {
	s1 := &model.Schedule{ID: uuid.New(), Capacity: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	s2 := &model.Schedule{ID: uuid.New(), Capacity: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	mockRepo := &mockScheduleRepo{allSchedules: []*model.Schedule{s1, s2}}
	svc := schedules_service.NewScheduleService(mockRepo)

	got, err := svc.GetAllSchedules(map[string]bool{})
	require.NoError(t, err)
	require.Len(t, got, 2)
	require.Equal(t, s1.ID, got[0].ID)
	require.Equal(t, s2.ID, got[1].ID)
}

func TestGetAllSchedules_RepoError(t *testing.T) {
	mockRepo := &mockScheduleRepo{errOnGetAll: errors.New("fail get all")}
	svc := schedules_service.NewScheduleService(mockRepo)

	got, err := svc.GetAllSchedules(map[string]bool{})
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGetAllSchedules_WithExpand_Success(t *testing.T) {
	s1 := &model.Schedule{ID: uuid.New()}
	s2 := &model.Schedule{ID: uuid.New()}
	mockRepo := &mockScheduleRepo{allSchedulesExp: []*model.Schedule{s1, s2}}
	svc := schedules_service.NewScheduleService(mockRepo)

	expand := map[string]bool{"course": true}
	got, err := svc.GetAllSchedules(expand)
	require.NoError(t, err)
	require.Len(t, got, 2)
	require.Equal(t, s1.ID, got[0].ID)
	require.Equal(t, s2.ID, got[1].ID)
}

func TestGetAllSchedules_WithExpand_RepoError(t *testing.T) {
	mockRepo := &mockScheduleRepo{errOnGetAllExp: errors.New("exp fail")}
	svc := schedules_service.NewScheduleService(mockRepo)

	expand := map[string]bool{"branch": true}
	got, err := svc.GetAllSchedules(expand)
	require.Error(t, err)
	require.Nil(t, got)
}

func TestUpdateScheduleByID_Success(t *testing.T) {
	mockRepo := &mockScheduleRepo{}
	svc := schedules_service.NewScheduleService(mockRepo)

	id := uuid.New()
	updateData := map[string]any{
		"capacity": 22,
	}
	err := svc.UpdateScheduleByID(id, updateData)
	require.NoError(t, err)
	require.Equal(t, id, mockRepo.updatedID)
	require.Equal(t, updateData, mockRepo.updatedData)
}

func TestUpdateScheduleByID_RepoError(t *testing.T) {
	mockRepo := &mockScheduleRepo{errOnUpdate: errors.New("update error")}
	svc := schedules_service.NewScheduleService(mockRepo)

	id := uuid.New()
	updateData := map[string]any{
		"capacity": 0,
	}
	err := svc.UpdateScheduleByID(id, updateData)
	require.Error(t, err)
}

func TestDeleteScheduleByID_Success(t *testing.T) {
	mockRepo := &mockScheduleRepo{}
	svc := schedules_service.NewScheduleService(mockRepo)

	id := uuid.New()
	err := svc.DeleteScheduleByID(id)
	require.NoError(t, err)
	require.Equal(t, id, mockRepo.deletedID)
}

func TestDeleteScheduleByID_RepoError(t *testing.T) {
	mockRepo := &mockScheduleRepo{errOnDelete: errors.New("delete err")}
	svc := schedules_service.NewScheduleService(mockRepo)

	id := uuid.New()
	err := svc.DeleteScheduleByID(id)
	require.Error(t, err)
}
