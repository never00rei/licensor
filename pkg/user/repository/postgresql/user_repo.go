package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/never00rei/licensor/domain"
)

type postgresqlUserRepo struct {
	pool *pgxpool.Pool
}

// NewPostgresqlUserRepo will create an object that represent the user.Repository interface
func NewPostgresqlUserRepo(pool *pgxpool.Pool) domain.UserRepository {
	return &postgresqlUserRepo{pool}
}

// Create will create a new user in the user database.
func (p *postgresqlUserRepo) Create(ctx context.Context, user *domain.User) error {
	args := pgx.NamedArgs{
		"username": user.Username,
		"email":    user.Email,
		"org_uuid": user.OrgUUID,
		"api_key":  user.ApiKey,
	}

	query := `
	INSERT INTO application.user(
		username,
		email,
		org_uuid,
		api_key
	)
	VALUES (
		@username,
		@email,
		@org_uuid,
		@api_key
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

// GetAll will return all of the users in the user database.
func (p *postgresqlUserRepo) GetAll(ctx context.Context, tenantID string) ([]*domain.User, error) {
	var users []*domain.User

	args := pgx.NamedArgs{
		"org_uuid": tenantID,
	}

	query := `
	SELECT 
		user_id,
		username,
		email,
		org_uuid,
		created_at,
		updated_at
	FROM application.user
	
	WHERE deleted = false
	AND org_uuid = @org_uuid
	`

	rows, err := p.pool.Query(ctx, query, args)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.UserID, &u.Username, &u.Email, &u.OrgUUID, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
