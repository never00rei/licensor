package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/never00rei/licensor/domain"
)

type postgresqlTenantRepo struct {
	Conn *pgxpool.Conn
}

// NewPostgresqlTenantRepo will create an object that represent the tenant.Repository interface
func NewPostgresqlTenantRepo(Conn *pgxpool.Conn) domain.TenantRepository {
	return &postgresqlTenantRepo{Conn}
}

// GetAll will return all of the tenants in the tenant database.
func (p *postgresqlTenantRepo) GetAll(ctx context.Context) ([]*domain.Tenant, error) {
	var tenants []*domain.Tenant

	rows, err := p.Conn.Query(ctx, "SELECT org_id, org_name, org_uuid, table_name, created_at, updated_at FROM tenant")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t domain.Tenant
		err := rows.Scan(&t.OrgID, &t.OrgName, &t.OrgUUID, &t.TableName, &t.CreatedAt, &t.UpdatedAt)
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
