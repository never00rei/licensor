package management

import (
	"context"

	"github.com/never00rei/licensor/domain"
)

type ManagementService struct {
	managementRepo domain.ManagementRepository
}

// NewManagementService will create an object that represent the management service
func NewManagementService(managementRepo domain.ManagementRepository) *ManagementService {
	return &ManagementService{managementRepo}
}

func (s *ManagementService) Get(ctx context.Context, username string) (*domain.ManagementUser, error) {
	return s.managementRepo.Get(ctx, username)
}

// GetAll will return all of the management users in the management user database.
func (s *ManagementService) GetAll(ctx context.Context) ([]*domain.ManagementUser, error) {
	return s.managementRepo.GetAll(ctx)
}