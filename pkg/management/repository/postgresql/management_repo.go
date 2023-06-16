package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func (p *postgresqlManagementRepo) Delete(ctx context.Context, username string) error {
	args := pgx.NamedArgs{
		"username": username,
	}

	query := `
	UPDATE management.management_user 
		SET deleted = true, 
		deleted_at = CURRENT_TIMESTAMP, 
		updated_at = CURRENT_TIMESTAMP 
	WHERE username = @username
	`

	_, err := p.pool.Exec(ctx, query, args)
	return err
}

func (p *postgresqlManagementRepo) Create(ctx context.Context, user *domain.ManagementUser) error {
	args := pgx.NamedArgs{
		"username": user.Username,
		"email":    user.Email,
		"api_key":  user.ApiKey,
		"is_admin": user.IsAdmin,
	}

	query := `
	INSERT INTO management.management_user(
		username, 
		email, 
		api_key, 
		is_admin
	) 
	
	VALUES (
		@username, 
		@email, 
		@api_key, 
		@is_admin
	)
	`

	_, err := p.pool.Exec(ctx, query, args)
	if pgerr, ok := err.(*pgconn.PgError); ok {
		if pgerr.Code == "23505" {
			return domain.ErrDuplicateUserExists
		}
	}
	return err
}

// Get will return the matching Management User in the Management Users database.
func (p *postgresqlManagementRepo) Get(ctx context.Context, username string) (*domain.ManagementUser, error) {
	var managementUser *domain.ManagementUser

	args := pgx.NamedArgs{
		"username": username,
	}

	query := `
	SELECT 
		user_id, 
		username, 
		email, 
		api_key, 
		created_at, 
		updated_at, 
		active,
		is_admin 
	FROM management.management_user 
	WHERE 
		deleted IS NOT true 
	AND 
		username = @username
	`

	// This query returns the user data from the database, as long as they are not deleted.
	rows, err := p.pool.Query(ctx, query, args)
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

	query := `
	SELECT 
		user_id, 
		username, 
		email, 
		api_key, 
		created_at, 
		updated_at, 
		active, 
		is_admin 
	FROM management.management_user 
	WHERE 
		deleted IS NOT true
	`

	// This query returns all of the user data from the database, as long as they are not deleted.
	rows, err := p.pool.Query(ctx, query)
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
