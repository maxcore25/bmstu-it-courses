package service_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/model"
	branches_service "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/service"
)

type mockBranchRepo struct {
	created     *model.Branch
	errOnCreate error
}

func (m *mockBranchRepo) Create(branch *model.Branch) error {
	if m.errOnCreate != nil {
		return m.errOnCreate
	}
	branch.ID = uuid.New()
	m.created = branch
	return nil
}
func (m *mockBranchRepo) GetByID(id uuid.UUID) (*model.Branch, error)              { return nil, nil }
func (m *mockBranchRepo) GetAll() ([]*model.Branch, error)                         { return nil, nil }
func (m *mockBranchRepo) UpdateByID(id uuid.UUID, updateData map[string]any) error { return nil }
func (m *mockBranchRepo) DeleteByID(id uuid.UUID) error                            { return nil }

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
