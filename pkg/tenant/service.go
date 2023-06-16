package tenant

import (
	"context"

	"github.com/never00rei/licensor/domain"
)

type TenantService struct {
	tenantRepo domain.TenantRepository
}

// NewTenantService will create an object that represent the tenant service
func NewTenantService(tenantRepo domain.TenantRepository) *TenantService {
	return &TenantService{tenantRepo}
}

// Create will create a new tenant in the tenant database.
func (s *TenantService) Create(ctx context.Context, tenant *domain.Tenant) error {
	return s.tenantRepo.Create(ctx, tenant)
}

// GetAll will return all of the tenants in the tenant database.
func (s *TenantService) GetAll(ctx context.Context) ([]*domain.Tenant, error) {
	return s.tenantRepo.GetAll(ctx)
}
