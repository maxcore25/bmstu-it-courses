package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/model"
	branches_service "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/service"
)

type mockBranchRepo struct {
	created     *model.Branch
	errOnCreate error

	branchByID   *model.Branch
	errOnGetByID error

	allBranches []*model.Branch
	errOnGetAll error

	errOnUpdate error
	updatedData map[string]any
	updatedID   uuid.UUID

	errOnDelete error
	deletedID   uuid.UUID
}

func (m *mockBranchRepo) Create(branch *model.Branch) error {
	if m.errOnCreate != nil {
		return m.errOnCreate
	}
	branch.ID = uuid.New()
	m.created = branch
	return nil
}

func (m *mockBranchRepo) GetByID(id uuid.UUID) (*model.Branch, error) {
	if m.errOnGetByID != nil {
		return nil, m.errOnGetByID
	}
	return m.branchByID, nil
}

func (m *mockBranchRepo) GetAll() ([]*model.Branch, error) {
	if m.errOnGetAll != nil {
		return nil, m.errOnGetAll
	}
	return m.allBranches, nil
}

func (m *mockBranchRepo) UpdateByID(id uuid.UUID, updateData map[string]any) error {
	m.updatedID = id
	m.updatedData = updateData
	if m.errOnUpdate != nil {
		return m.errOnUpdate
	}
	return nil
}

func (m *mockBranchRepo) DeleteByID(id uuid.UUID) error {
	m.deletedID = id
	if m.errOnDelete != nil {
		return m.errOnDelete
	}
	return nil
}

func TestCreateBranch_Success(t *testing.T) {
	mockRepo := &mockBranchRepo{}
	svc := branches_service.NewBranchService(mockRepo)

	req := dto.CreateBranchRequest{Address: "Test Address", Rooms: 5}
	b, err := svc.CreateBranch(req)
	require.NoError(t, err)
	require.NotNil(t, b)
	require.Equal(t, "Test Address", b.Address)
	require.Equal(t, 5, b.Rooms)
	require.NotEqual(t, uuid.Nil, b.ID)
}

func TestCreateBranch_RepoError(t *testing.T) {
	mockRepo := &mockBranchRepo{errOnCreate: errors.New("db fail")}
	svc := branches_service.NewBranchService(mockRepo)

	req := dto.CreateBranchRequest{Address: "X", Rooms: 1}
	b, err := svc.CreateBranch(req)
	require.Error(t, err)
	require.Nil(t, b)
}

func TestGetBranchByID_Success(t *testing.T) {
	branchID := uuid.New()
	expected := &model.Branch{
		ID:        branchID,
		Address:   "Addr 123",
		Rooms:     3,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockRepo := &mockBranchRepo{branchByID: expected}
	svc := branches_service.NewBranchService(mockRepo)

	res, err := svc.GetBranch(branchID)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expected.ID, res.ID)
	require.Equal(t, expected.Address, res.Address)
	require.Equal(t, expected.Rooms, res.Rooms)
}

func TestGetBranchByID_RepoError(t *testing.T) {
	branchID := uuid.New()
	mockRepo := &mockBranchRepo{errOnGetByID: errors.New("not found")}
	svc := branches_service.NewBranchService(mockRepo)

	res, err := svc.GetBranch(branchID)
	require.Error(t, err)
	require.Nil(t, res)
}

func TestGetAllBranches_Success(t *testing.T) {
	b1 := &model.Branch{ID: uuid.New(), Address: "A1", Rooms: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	b2 := &model.Branch{ID: uuid.New(), Address: "A2", Rooms: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	mockRepo := &mockBranchRepo{allBranches: []*model.Branch{b1, b2}}
	svc := branches_service.NewBranchService(mockRepo)

	list, err := svc.GetAllBranches()
	require.NoError(t, err)
	require.Len(t, list, 2)
	require.Equal(t, b1.ID, list[0].ID)
	require.Equal(t, b2.ID, list[1].ID)
}

func TestGetAllBranches_RepoError(t *testing.T) {
	mockRepo := &mockBranchRepo{errOnGetAll: errors.New("fail get all")}
	svc := branches_service.NewBranchService(mockRepo)

	list, err := svc.GetAllBranches()
	require.Error(t, err)
	require.Nil(t, list)
}

func TestUpdateBranchByID_Success(t *testing.T) {
	mockRepo := &mockBranchRepo{}
	svc := branches_service.NewBranchService(mockRepo)

	id := uuid.New()
	updateData := map[string]any{
		"address": "NewAddr",
		"rooms":   10,
	}
	err := svc.UpdateBranchByID(id, updateData)
	require.NoError(t, err)
	require.Equal(t, id, mockRepo.updatedID)
}

func TestUpdateBranchByID_RepoError(t *testing.T) {
	mockRepo := &mockBranchRepo{errOnUpdate: errors.New("update error")}
	svc := branches_service.NewBranchService(mockRepo)

	id := uuid.New()
	updateData := map[string]any{
		"address": "fail",
		"rooms":   0,
	}
	err := svc.UpdateBranchByID(id, updateData)
	require.Error(t, err)
}

func TestDeleteBranchByID_Success(t *testing.T) {
	mockRepo := &mockBranchRepo{}
	svc := branches_service.NewBranchService(mockRepo)

	id := uuid.New()
	err := svc.DeleteBranchByID(id)
	require.NoError(t, err)
	require.Equal(t, id, mockRepo.deletedID)
}

func TestDeleteBranchByID_RepoError(t *testing.T) {
	mockRepo := &mockBranchRepo{errOnDelete: errors.New("delete err")}
	svc := branches_service.NewBranchService(mockRepo)

	id := uuid.New()
	err := svc.DeleteBranchByID(id)
	require.Error(t, err)
}
