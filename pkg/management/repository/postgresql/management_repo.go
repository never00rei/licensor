package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/never00rei/licensor/domain"
)

type postgresqlManagementRepo struct {
	pool *pgxpool.Pool
}

// NewPostgresqlManagementRepo will create an object that represent the Management.Repository interface
func NewPostgresqlManagementRepo(pool *pgxpool.Pool) domain.ManagementRepository {
	return &postgresqlManagementRepo{pool}
}

// GetAll will return all of the Managements in the Management database.
func (p *postgresqlManagementRepo) Get(ctx context.Context, username string) (*domain.ManagementUser, error) {
	var managementUser *domain.ManagementUser

	rows, err := p.pool.Query(ctx, "SELECT user_id, username, email, api_key, created_at, updated_at, active, is_admin FROM management_user WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}
	for rows.Next() {
		var m domain.ManagementUser
		err := rows.Scan(&m.UserID, &m.Username, &m.Email, &m.ApiKey, &m.CreatedAt, &m.UpdatedAt, &m.Active, &m.IsAdmin)
		if err != nil {
			return nil, err
		}
		managementUser = &m
	}

	if managementUser == nil { // TODO revist this as its a hack
		return nil, domain.ErrManagementUserNotFound
	}

	return managementUser, nil
}

// GetAll will return all of the Managements in the Management database.
func (p *postgresqlManagementRepo) GetAll(ctx context.Context) ([]*domain.ManagementUser, error) {
	var ManagementUsers []*domain.ManagementUser

	rows, err := p.pool.Query(ctx, "SELECT user_id, username, email, api_key, created_at, updated_at, active, is_admin FROM management_user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m domain.ManagementUser
		err := rows.Scan(&m.UserID, &m.Username, &m.Email, &m.ApiKey, &m.CreatedAt, &m.UpdatedAt, &m.Active, &m.IsAdmin)
		if err != nil {
			return nil, err
		}
		ManagementUsers = append(ManagementUsers, &m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ManagementUsers, nil
}
