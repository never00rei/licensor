package user

// import (
// 	"context"
// 	"testing"

// 	mock "github.com/never00rei/licensor/mock"

// 	"github.com/never00rei/licensor/domain"
// 	"github.com/stretchr/testify/assert"
// )

// func exampleTenant(orgName string) *domain.Tenant {
// 	return &domain.Tenant{
// 		OrgName: "Example Org",
// 		OrgUUID: "00000000-0000-0000-0000-000000000000",
// 	}
// }

// func TestGetAll(t *testing.T) {
// 	testcases := map[string]struct {
// 		repoReturn []*domain.Tenant
// 		repoErr    error
// 	}{
// 		"success": {
// 			repoReturn: []*domain.Tenant{
// 				exampleTenant("Example Org"),
// 			},
// 			repoErr: nil,
// 		},
// 		"error": {
// 			repoReturn: nil,
// 			repoErr:    assert.AnError,
// 		},
// 		"emptyList": {
// 			repoReturn: []*domain.Tenant{},
// 			repoErr:    nil,
// 		},
// 	}

// 	for name, tc := range testcases {
// 		t.Run(name, func(t *testing.T) {

// 			// Generate the mock repository
// 			mockRepo := &mock.MockTenantDatabase{
// 				GetAllFunc: func(ctx context.Context) ([]*domain.Tenant, error) {
// 					return tc.repoReturn, tc.repoErr
// 				},
// 			}

// 			// Generate the service
// 			svc := NewTenantService(mockRepo)

// 			// Call the service
// 			tenants, err := svc.GetAll(context.Background())

// 			assert.Equal(t, tc.repoReturn, tenants)
// 			assert.Equal(t, tc.repoErr, err)
// 		})

// 	}

// }
