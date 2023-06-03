package management

import (
	"context"
	"testing"

	mock "github.com/never00rei/licensor/mock"

	"github.com/never00rei/licensor/domain"
	"github.com/stretchr/testify/assert"
)

func exampleManagementUser(username string) *domain.ManagementUser {
	return &domain.ManagementUser{
		UserID:   1,
		Username: username,
		ApiKey:   "bar",
		Email:    "foo@bar.com",
	}

}

func TestGetAll(t *testing.T) {
	testcases := map[string]struct {
		repoReturn []*domain.ManagementUser
		repoErr    error
	}{
		"success": {
			repoReturn: []*domain.ManagementUser{
				exampleManagementUser("foobar"),
			},
			repoErr: nil,
		},
		"error": {
			repoReturn: nil,
			repoErr:    assert.AnError,
		},
		"emptyList": {
			repoReturn: []*domain.ManagementUser{},
			repoErr:    nil,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {

			// Generate the mock repository
			mockRepo := &mock.MockManagementDatabase{
				GetAllFunc: func(ctx context.Context) ([]*domain.ManagementUser, error) {
					return tc.repoReturn, tc.repoErr
				},
			}

			// Generate the service
			svc := NewManagementService(mockRepo)

			// Call the service
			tenants, err := svc.GetAll(context.Background())

			assert.Equal(t, tc.repoReturn, tenants)
			assert.Equal(t, tc.repoErr, err)
		})

	}

}
