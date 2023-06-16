package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

// Create will create a new tenant in the tenant database.
func (p *postgresqlTenantRepo) Create(ctx context.Context, tenant *domain.Tenant) error {
	args := pgx.NamedArgs{
		"org_name": tenant.OrgName,
	}

	query := `
	INSERT INTO management.tenant(
		org_name
	)
	VALUES (
		@org_name
	)
	`

	_, err := p.pool.Exec(ctx, query, args)
	if pgerr, ok := err.(*pgconn.PgError); ok {
		if pgerr.Code == "23505" {
			return domain.ErrDuplicateTenantExists
		}
	}
	return err
}

// GetAll will return all of the tenants in the tenant database.
func (p *postgresqlTenantRepo) GetAll(ctx context.Context) ([]*domain.Tenant, error) {
	var tenants []*domain.Tenant

	query := `
	SELECT 
		org_name,
		org_uuid,
		created_at,
		updated_at
	FROM management.tenant
	`

	rows, err := p.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t domain.Tenant
		err := rows.Scan(&t.OrgName, &t.OrgUUID, &t.CreatedAt, &t.UpdatedAt)
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
