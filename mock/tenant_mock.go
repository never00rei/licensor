package mock

import (
	"context"

	"github.com/never00rei/licensor/domain"
)

type MockTenantDatabase struct {
	GetAllFunc func(ctx context.Context) ([]*domain.Tenant, error)
}

func (m *MockTenantDatabase) GetAll(ctx context.Context) ([]*domain.Tenant, error) {
	return m.GetAllFunc(ctx)
}
