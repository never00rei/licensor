package mock

import (
	"context"

	"github.com/never00rei/licensor/domain"
)

type MockTenantDatabase struct {
	CreateFunc func(ctx context.Context, tenant *domain.Tenant) error
	GetAllFunc func(ctx context.Context) ([]*domain.Tenant, error)
}

func (m *MockTenantDatabase) Create(ctx context.Context, tenant *domain.Tenant) error {
	return m.CreateFunc(ctx, tenant)
}

func (m *MockTenantDatabase) GetAll(ctx context.Context) ([]*domain.Tenant, error) {
	return m.GetAllFunc(ctx)
}
