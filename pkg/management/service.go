package management

import (
	"context"

	"github.com/never00rei/licensor/domain"
	"github.com/never00rei/licensor/pkg/keygen"
	"github.com/pkg/errors"
)

type ManagementService struct {
	managementRepo domain.ManagementRepository
}

// NewManagementService will create an object that represent the management service
func NewManagementService(managementRepo domain.ManagementRepository) *ManagementService {
	return &ManagementService{managementRepo}
}

func (s *ManagementService) Delete(ctx context.Context, username string) error {
	return s.managementRepo.Delete(ctx, username)
}

// Create will create a new management user in the management user database.
func (s *ManagementService) Create(ctx context.Context, user *domain.ManagementUser) (string, error) {

	apiKey, err := keygen.GenerateApiKey(user.Username)
	if err != nil {
		return "", errors.Wrap(err, "could not generate api key")
	}

	apiHash, err := keygen.SaltAndHashAPIKey(apiKey)
	if err != nil {
		return "", errors.Wrap(err, "could not hash api key")
	}

	user.ApiKey = apiHash

	err = s.managementRepo.Create(ctx, user)
	if err != nil {
		return "", errors.Wrap(err, "could not create management user")
	}

	return apiKey, err
}

func (s *ManagementService) Get(ctx context.Context, username string) (*domain.ManagementUser, error) {
	return s.managementRepo.Get(ctx, username)
}

// GetAll will return all of the management users in the management user database.
func (s *ManagementService) GetAll(ctx context.Context) ([]*domain.ManagementUser, error) {
	return s.managementRepo.GetAll(ctx)
}
