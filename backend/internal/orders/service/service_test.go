package service_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/orders/model"
	orders_service "github.com/maxcore25/bmstu-it-courses/backend/internal/orders/service"
)

type mockOrderRepo struct {
	created     *model.Order
	errOnCreate error

	ordersByUser   []*model.Order
	errOnGetByUser error

	ordersByUserExp   []*model.Order
	errOnGetByUserExp error

	orderByID    *model.Order
	errOnGetByID error

	orderByIDExp    *model.Order
	errOnGetByIDExp error

	allOrders   []*model.Order
	errOnGetAll error

	allOrdersExp   []*model.Order
	errOnGetAllExp error

	updatedID   uuid.UUID
	updatedData map[string]any
	errOnUpdate error

	deletedID   uuid.UUID
	errOnDelete error

	db *gorm.DB // can be nil for tests, but must exist
}

func (m *mockOrderRepo) Create(order *model.Order) error {
	if m.errOnCreate != nil {
		return m.errOnCreate
	}
	order.ID = uuid.New()
	m.created = order
	return nil
}

func (m *mockOrderRepo) GetByUser(userID uuid.UUID) ([]*model.Order, error) {
	if m.errOnGetByUser != nil {
		return nil, m.errOnGetByUser
	}
	return m.ordersByUser, nil
}

func (m *mockOrderRepo) GetByUserWithExpand(userID uuid.UUID, expand map[string]bool) ([]*model.Order, error) {
	if m.errOnGetByUserExp != nil {
		return nil, m.errOnGetByUserExp
	}
	return m.ordersByUserExp, nil
}

func (m *mockOrderRepo) GetByID(id uuid.UUID) (*model.Order, error) {
	if m.errOnGetByID != nil {
		return nil, m.errOnGetByID
	}
	return m.orderByID, nil
}

func (m *mockOrderRepo) GetByIDWithExpand(id uuid.UUID, expand map[string]bool) (*model.Order, error) {
	if m.errOnGetByIDExp != nil {
		return nil, m.errOnGetByIDExp
	}
	return m.orderByIDExp, nil
}

func (m *mockOrderRepo) GetAll() ([]*model.Order, error) {
	if m.errOnGetAll != nil {
		return nil, m.errOnGetAll
	}
	return m.allOrders, nil
}

func (m *mockOrderRepo) GetAllWithExpand(expand map[string]bool) ([]*model.Order, error) {
	if m.errOnGetAllExp != nil {
		return nil, m.errOnGetAllExp
	}
	return m.allOrdersExp, nil
}

func (m *mockOrderRepo) UpdateByID(id uuid.UUID, updateData map[string]any) error {
	m.updatedID = id
	m.updatedData = updateData
	if m.errOnUpdate != nil {
		return m.errOnUpdate
	}
	return nil
}

func (m *mockOrderRepo) DeleteByID(id uuid.UUID) error {
	m.deletedID = id
	if m.errOnDelete != nil {
		return m.errOnDelete
	}
	return nil
}

// DB mock to satisfy repository.OrderRepository interface
func (m *mockOrderRepo) DB() *gorm.DB {
	return m.db
}

func TestGetOrdersByUser_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{
		ordersByUser: []*model.Order{
			{ID: uuid.New(), Price: 50},
			{ID: uuid.New(), Price: 70},
		},
	}
	svc := orders_service.NewOrderService(mockRepo)

	userID := uuid.New()
	res, err := svc.GetOrdersByUser(userID, map[string]bool{})
	require.NoError(t, err)
	require.Len(t, res, 2)
	require.Equal(t, mockRepo.ordersByUser[0].ID, res[0].ID)
}

func TestGetOrdersByUser_RepoError(t *testing.T) {
	mockRepo := &mockOrderRepo{errOnGetByUser: errors.New("fail by user")}
	svc := orders_service.NewOrderService(mockRepo)

	userID := uuid.New()
	res, err := svc.GetOrdersByUser(userID, map[string]bool{})
	require.Error(t, err)
	require.Nil(t, res)
}

func TestGetOrdersByUserWithExpand_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{
		ordersByUserExp: []*model.Order{
			{ID: uuid.New()}, {ID: uuid.New()},
		},
	}
	svc := orders_service.NewOrderService(mockRepo)

	userID := uuid.New()
	expand := map[string]bool{"course": true}
	res, err := svc.GetOrdersByUser(userID, expand)
	require.NoError(t, err)
	require.Len(t, res, 2)
	require.Equal(t, mockRepo.ordersByUserExp[0].ID, res[0].ID)
}

func TestGetOrdersByUserWithExpand_RepoError(t *testing.T) {
	mockRepo := &mockOrderRepo{errOnGetByUserExp: errors.New("expand fail")}
	svc := orders_service.NewOrderService(mockRepo)

	userID := uuid.New()
	res, err := svc.GetOrdersByUser(userID, map[string]bool{"whatever": true})
	require.Error(t, err)
	require.Nil(t, res)
}

func TestGetOrderByID_Success(t *testing.T) {
	expected := &model.Order{ID: uuid.New()}
	mockRepo := &mockOrderRepo{orderByID: expected}
	svc := orders_service.NewOrderService(mockRepo)

	res, err := svc.GetOrder(expected.ID, map[string]bool{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expected.ID, res.ID)
}

func TestGetOrderByID_RepoError(t *testing.T) {
	mockRepo := &mockOrderRepo{errOnGetByID: errors.New("no such")}
	svc := orders_service.NewOrderService(mockRepo)

	id := uuid.New()
	res, err := svc.GetOrder(id, map[string]bool{})
	require.Error(t, err)
	require.Nil(t, res)
}

func TestGetOrderByIDWithExpand_Success(t *testing.T) {
	expected := &model.Order{ID: uuid.New()}
	mockRepo := &mockOrderRepo{orderByIDExp: expected}
	svc := orders_service.NewOrderService(mockRepo)

	expand := map[string]bool{"course": true}
	res, err := svc.GetOrder(expected.ID, expand)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expected.ID, res.ID)
}

func TestGetOrderByIDWithExpand_RepoError(t *testing.T) {
	mockRepo := &mockOrderRepo{errOnGetByIDExp: errors.New("expand fail")}
	svc := orders_service.NewOrderService(mockRepo)

	id := uuid.New()
	res, err := svc.GetOrder(id, map[string]bool{"x": true})
	require.Error(t, err)
	require.Nil(t, res)
}

func TestGetAllOrders_Success(t *testing.T) {
	orders := []*model.Order{
		{ID: uuid.New(), Price: 25},
		{ID: uuid.New(), Price: 40},
	}
	mockRepo := &mockOrderRepo{allOrders: orders}
	svc := orders_service.NewOrderService(mockRepo)

	got, err := svc.GetAllOrders(map[string]bool{})
	require.NoError(t, err)
	require.Len(t, got, 2)
	require.Equal(t, orders[0].ID, got[0].ID)
	require.Equal(t, orders[1].ID, got[1].ID)
}

func TestGetAllOrders_RepoError(t *testing.T) {
	mockRepo := &mockOrderRepo{errOnGetAll: errors.New("fail get all")}
	svc := orders_service.NewOrderService(mockRepo)

	got, err := svc.GetAllOrders(map[string]bool{})
	require.Error(t, err)
	require.Nil(t, got)
}

func TestGetAllOrdersWithExpand_Success(t *testing.T) {
	orders := []*model.Order{{ID: uuid.New()}, {ID: uuid.New()}}
	mockRepo := &mockOrderRepo{allOrdersExp: orders}
	svc := orders_service.NewOrderService(mockRepo)

	got, err := svc.GetAllOrders(map[string]bool{"client": true})
	require.NoError(t, err)
	require.Len(t, got, 2)
	require.Equal(t, orders[0].ID, got[0].ID)
	require.Equal(t, orders[1].ID, got[1].ID)
}

func TestGetAllOrdersWithExpand_RepoError(t *testing.T) {
	mockRepo := &mockOrderRepo{errOnGetAllExp: errors.New("expand fail")}
	svc := orders_service.NewOrderService(mockRepo)

	got, err := svc.GetAllOrders(map[string]bool{"course": true})
	require.Error(t, err)
	require.Nil(t, got)
}

func TestUpdateOrderByID_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{}
	svc := orders_service.NewOrderService(mockRepo)

	id := uuid.New()
	updateData := map[string]any{
		"price": 99.5,
	}
	err := svc.UpdateOrderByID(id, updateData)
	require.NoError(t, err)
	require.Equal(t, id, mockRepo.updatedID)
	require.Equal(t, updateData, mockRepo.updatedData)
}

func TestUpdateOrderByID_RepoError(t *testing.T) {
	mockRepo := &mockOrderRepo{errOnUpdate: errors.New("update error")}
	svc := orders_service.NewOrderService(mockRepo)

	id := uuid.New()
	updateData := map[string]any{
		"price": 10,
	}
	err := svc.UpdateOrderByID(id, updateData)
	require.Error(t, err)
}

func TestDeleteOrderByID_Success(t *testing.T) {
	mockRepo := &mockOrderRepo{}
	svc := orders_service.NewOrderService(mockRepo)

	id := uuid.New()
	err := svc.DeleteOrderByID(id)
	require.NoError(t, err)
	require.Equal(t, id, mockRepo.deletedID)
}

func TestDeleteOrderByID_RepoError(t *testing.T) {
	mockRepo := &mockOrderRepo{errOnDelete: errors.New("delete err")}
	svc := orders_service.NewOrderService(mockRepo)

	id := uuid.New()
	err := svc.DeleteOrderByID(id)
	require.Error(t, err)
}
