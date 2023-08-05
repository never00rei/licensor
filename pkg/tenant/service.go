package tenant

import (
	"context"

	"github.com/never00rei/licensor/domain"
	"github.com/never00rei/licensor/pkg/keygen"
	"github.com/pkg/errors"
)

type TenantService struct {
	tenantRepo domain.TenantRepository
	userRepo   domain.UserRepository
}

// NewTenantService will create an object that represent the tenant service
func NewTenantService(tenantRepo domain.TenantRepository, userRepo domain.UserRepository) *TenantService {
	return &TenantService{
		tenantRepo: tenantRepo,
		userRepo:   userRepo,
	}
}

// Create will create a new tenant in the tenant database.
func (s *TenantService) Create(ctx context.Context, tenant *domain.Tenant) (string, error) {
	err := s.tenantRepo.Create(ctx, tenant)
	if err != nil {
		return "", err
	}

	// Create the tenant's admin user
	tenantAdmin := &domain.User{
		Username: "admin",
		OrgUUID:  tenant.OrgUUID,
	}

	apiKey, err := keygen.GenerateApiKey(tenantAdmin.Username)
	if err != nil {
		return "", errors.Wrap(err, "could not generate api key")
	}

	apiHash, err := keygen.SaltAndHashAPIKey(apiKey)
	if err != nil {
		return "", errors.Wrap(err, "could not hash api key")
	}

	tenantAdmin.ApiKey = apiHash

	// Create the tenant's admin user
	err = s.userRepo.Create(ctx, tenantAdmin)
	if err != nil {
		return "", errors.Wrap(err, "could not create tenant admin user")
	}

	return apiKey, nil
}

// GetAll will return all of the tenants in the tenant database.
func (s *TenantService) GetAll(ctx context.Context) ([]*domain.Tenant, error) {
	return s.tenantRepo.GetAll(ctx)
}
