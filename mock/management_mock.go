package mock

import (
	"context"

	"github.com/never00rei/licensor/domain"
)

type MockManagementDatabase struct {
	GetAllFunc func(ctx context.Context) ([]*domain.ManagementUser, error)
}

func (m *MockManagementDatabase) GetAll(ctx context.Context) ([]*domain.ManagementUser, error) {
	return m.GetAllFunc(ctx)
}
