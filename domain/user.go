package domain

import (
	"context"
	"time"
)

// This represents a user in the tenant database.
// Users in this table will only be able to manage licenses, not tenants.
type User struct {
	UserID    string    `db:"user_id" json:"user_id"`
	OrgUUID   string    `db:"org_uuid" json:"org_uuid"`
	Username  string    `db:"username" json:"username"`
	ApiKey    string    `db:"api_key" json:"-"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Active    bool      `db:"active" json:"active"`
	LastLogin time.Time `db:"last_login" json:"last_login"`
	Deleted   bool      `db:"deleted" json:"-"`
	DeletedAt time.Time `db:"deleted_at" json:"-"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetAll(ctx context.Context, tenantID string) ([]*User, error)
}
