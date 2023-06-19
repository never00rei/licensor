package user

import (
	"context"

	"github.com/never00rei/licensor/domain"
	"github.com/never00rei/licensor/pkg/keygen"
	"github.com/pkg/errors"
)

type UserService struct {
	userRepo domain.UserRepository
}

// NewUserService will create an object that represent the tenant service
func NewUserService(userRepo domain.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Create(ctx context.Context, user *domain.User) (string, error) {
	apiKey, err := keygen.GenerateApiKey(user.Username)
	if err != nil {
		return "", errors.Wrap(err, "could not generate api key")
	}

	apiHash, err := keygen.SaltAndHashAPIKey(apiKey)
	if err != nil {
		return "", errors.Wrap(err, "could not hash api key")
	}

	user.ApiKey = apiHash

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return "", errors.Wrap(err, "could not create management user")
	}

	return apiKey, nil
}

func (s *UserService) GetAll(ctx context.Context, tenantID string) ([]*domain.User, error) {
	return s.userRepo.GetAll(ctx, tenantID)
}
