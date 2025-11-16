package service_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
)

type mockUserRepo struct {
	createdUser *model.User
	errOnCreate error

	userByEmail     *model.User
	errOnGetByEmail error

	userByID     *model.User
	errOnGetByID error

	allUsers    []*model.User
	errOnGetAll error

	usersByFilter []*model.User
	errOnFind     error

	updatedID   uuid.UUID
	updatedData map[string]any
	errOnUpdate error

	deletedID   uuid.UUID
	errOnDelete error
}

func (m *mockUserRepo) Create(user *model.User) error {
	if m.errOnCreate != nil {
		return m.errOnCreate
	}
	// Simulate ID creation for the user if not already set
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	m.createdUser = user
	return nil
}

func (m *mockUserRepo) GetByEmail(email string) (*model.User, error) {
	if m.errOnGetByEmail != nil {
		return nil, m.errOnGetByEmail
	}
	return m.userByEmail, nil
}

func (m *mockUserRepo) GetByID(id uuid.UUID) (*model.User, error) {
	if m.errOnGetByID != nil {
		return nil, m.errOnGetByID
	}
	return m.userByID, nil
}

func (m *mockUserRepo) GetAll() ([]*model.User, error) {
	if m.errOnGetAll != nil {
		return nil, m.errOnGetAll
	}
	return m.allUsers, nil
}

func (m *mockUserRepo) Find(filter dto.UserFilter) ([]*model.User, error) {
	if m.errOnFind != nil {
		return nil, m.errOnFind
	}
	return m.usersByFilter, nil
}

func (m *mockUserRepo) UpdateByID(id uuid.UUID, updateData map[string]any) error {
	m.updatedID = id
	m.updatedData = updateData
	if m.errOnUpdate != nil {
		return m.errOnUpdate
	}
	return nil
}

func (m *mockUserRepo) DeleteByID(id uuid.UUID) error {
	m.deletedID = id
	if m.errOnDelete != nil {
		return m.errOnDelete
	}
	return nil
}

func TestCreateUser_Success(t *testing.T) {
	mockRepo := &mockUserRepo{}
	svc := service.NewUserService(mockRepo)

	portfolio := "portfolio info"
	rating := 4.7
	testimonialsCount := 5
	req := dto.CreateUserRequest{
		FirstName:         "Jane",
		LastName:          "Doe",
		MiddleName:        "Middle",
		Email:             "jane@example.com",
		Password:          "password123",
		Phone:             "1234567890",
		KnowledgeLevel:    "beginner",
		Role:              "student",
		Portfolio:         &portfolio,
		Rating:            &rating,
		TestimonialsCount: &testimonialsCount,
	}

	user, err := svc.CreateUser(req)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, req.Email, user.Email)
	require.NotEqual(t, uuid.Nil, user.ID)
	require.Equal(t, model.KnowledgeLevel(req.KnowledgeLevel), user.KnowledgeLevel)
	require.Equal(t, model.UserRole(req.Role), user.Role)
	require.NotEmpty(t, user.Password)
}

func TestCreateUser_EmailAlreadyExists(t *testing.T) {
	mockRepo := &mockUserRepo{
		userByEmail: &model.User{ID: uuid.New(), Email: "duplicate@mail.com"},
	}
	svc := service.NewUserService(mockRepo)

	req := dto.CreateUserRequest{
		Email:    "duplicate@mail.com",
		Password: "pass",
	}
	user, err := svc.CreateUser(req)
	require.Error(t, err)
	require.Nil(t, user)
}

func TestCreateUser_RepoError(t *testing.T) {
	mockRepo := &mockUserRepo{
		errOnCreate: errors.New("create failed"),
	}
	svc := service.NewUserService(mockRepo)

	req := dto.CreateUserRequest{
		FirstName: "X",
		LastName:  "Y",
		Email:     "u@mail.org",
		Password:  "zzz",
	}
	user, err := svc.CreateUser(req)
	require.Error(t, err)
	require.Nil(t, user)
}

func TestGetByEmail_Success(t *testing.T) {
	expected := &model.User{ID: uuid.New(), Email: "abc@def.com"}
	mockRepo := &mockUserRepo{userByEmail: expected}
	svc := service.NewUserService(mockRepo)

	user, err := svc.GetByEmail("abc@def.com")
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, expected.ID, user.ID)
}

func TestGetByEmail_RepoError(t *testing.T) {
	mockRepo := &mockUserRepo{errOnGetByEmail: errors.New("no found")}
	svc := service.NewUserService(mockRepo)

	user, err := svc.GetByEmail("not@exists.com")
	require.Error(t, err)
	require.Nil(t, user)
}

func TestGetUser_Success(t *testing.T) {
	expected := &model.User{ID: uuid.New(), Email: "e@e.com"}
	mockRepo := &mockUserRepo{userByID: expected}
	svc := service.NewUserService(mockRepo)

	user, err := svc.GetUser(expected.ID)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, expected.ID, user.ID)
}

func TestGetUser_NotFound(t *testing.T) {
	mockRepo := &mockUserRepo{errOnGetByID: errors.New("not found")}
	svc := service.NewUserService(mockRepo)

	id := uuid.New()
	user, err := svc.GetUser(id)
	require.Error(t, err)
	require.Nil(t, user)
}

func TestGetAllUsers_Success(t *testing.T) {
	all := []*model.User{
		{ID: uuid.New(), Email: "one@a.com"},
		{ID: uuid.New(), Email: "two@a.com"},
	}
	mockRepo := &mockUserRepo{allUsers: all}
	svc := service.NewUserService(mockRepo)

	users, err := svc.GetAllUsers()
	require.NoError(t, err)
	require.Len(t, users, 2)
	require.Equal(t, all[0].ID, users[0].ID)
}

func TestGetAllUsers_RepoError(t *testing.T) {
	mockRepo := &mockUserRepo{errOnGetAll: errors.New("fail all")}
	svc := service.NewUserService(mockRepo)

	users, err := svc.GetAllUsers()
	require.Error(t, err)
	require.Nil(t, users)
}

func TestGetUsers_FilterSuccess(t *testing.T) {
	filtered := []*model.User{
		{ID: uuid.New()},
	}
	mockRepo := &mockUserRepo{usersByFilter: filtered}
	svc := service.NewUserService(mockRepo)

	role := "student"
	filter := dto.UserFilter{Role: &role}
	users, err := svc.GetUsers(filter)
	require.NoError(t, err)
	require.Equal(t, filtered, users)
}

func TestGetUsers_RepoError(t *testing.T) {
	mockRepo := &mockUserRepo{errOnFind: errors.New("repo find fail")}
	svc := service.NewUserService(mockRepo)

	filter := dto.UserFilter{}
	users, err := svc.GetUsers(filter)
	require.Error(t, err)
	require.Nil(t, users)
}

func TestUpdateUserByID_Success(t *testing.T) {
	mockRepo := &mockUserRepo{}
	svc := service.NewUserService(mockRepo)

	id := uuid.New()
	updateData := map[string]any{
		"first_name": "Updated",
	}
	err := svc.UpdateUserByID(id, updateData)
	require.NoError(t, err)
	require.Equal(t, id, mockRepo.updatedID)
	require.Equal(t, updateData, mockRepo.updatedData)
}

func TestUpdateUserByID_RepoError(t *testing.T) {
	mockRepo := &mockUserRepo{errOnUpdate: errors.New("fail update")}
	svc := service.NewUserService(mockRepo)

	id := uuid.New()
	updateData := map[string]any{
		"last_name": "XFail",
	}
	err := svc.UpdateUserByID(id, updateData)
	require.Error(t, err)
}

func TestDeleteUserByID_Success(t *testing.T) {
	mockRepo := &mockUserRepo{}
	svc := service.NewUserService(mockRepo)

	id := uuid.New()
	err := svc.DeleteUserByID(id)
	require.NoError(t, err)
	require.Equal(t, id, mockRepo.deletedID)
}

func TestDeleteUserByID_RepoError(t *testing.T) {
	mockRepo := &mockUserRepo{errOnDelete: errors.New("cannot delete")}
	svc := service.NewUserService(mockRepo)

	id := uuid.New()
	err := svc.DeleteUserByID(id)
	require.Error(t, err)
}
