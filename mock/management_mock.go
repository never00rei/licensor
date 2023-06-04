package mock

import (
	"context"

	"github.com/never00rei/licensor/domain"
)

type MockManagementDatabase struct {
	GetAllFunc func(ctx context.Context) ([]*domain.ManagementUser, error)
}

func (m *MockManagementDatabase) Create(ctx context.Context, user *domain.ManagementUser) error {
	return nil
}

func (m *MockManagementDatabase) Delete(ctx context.Context, username string) error {
	return nil
}

func (m *MockManagementDatabase) GetAll(ctx context.Context) ([]*domain.ManagementUser, error) {
	return m.GetAllFunc(ctx)
}

func (m *MockManagementDatabase) Get(ctx context.Context, username string) (*domain.ManagementUser, error) {
	return nil, nil
}
