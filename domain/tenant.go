package domain

import (
	"context"
	"errors"
	"time"
)

var ErrDuplicateTenantExists error = errors.New("duplicate tenant exists")

// Administration Database Models.
//
// Tenant is a struct that represents a tenant in the database.
// Think of the "tenant" table as part of a top level administration instance.
// This will be in it's own database, and will maintain a list of tenants along
// with their associated database table names.
//
// This is to allow for multiple tenants to be stored in the same database server,
// but allow tenant data to be stored in seperate databases, providing data isolation.
// However, this adds complexity as we need to change the database connection at runtime...
type Tenant struct {
	OrgName   string    `db:"org_name"`
	OrgUUID   string    `db:"org_uuid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TenantRepository interface {
	Create(ctx context.Context, tenant *Tenant) error
	GetAll(ctx context.Context) ([]*Tenant, error)
}
