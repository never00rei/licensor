package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/never00rei/licensor/domain"
)

type postgresqlTenantRepo struct {
	pool *pgxpool.Pool
}

// NewPostgresqlTenantRepo will create an object that represent the tenant.Repository interface
func NewPostgresqlTenantRepo(pool *pgxpool.Pool) domain.TenantRepository {
	return &postgresqlTenantRepo{pool}
}

// GetAll will return all of the tenants in the tenant database.
func (p *postgresqlTenantRepo) GetAll(ctx context.Context) ([]*domain.Tenant, error) {
	var tenants []*domain.Tenant

	rows, err := p.pool.Query(ctx, "SELECT org_id, org_name, org_uuid, schema_name, created_at, updated_at FROM tenant")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t domain.Tenant
		err := rows.Scan(&t.OrgID, &t.OrgName, &t.OrgUUID, &t.SchemaName, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tenants = append(tenants, &t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tenants, nil
}
