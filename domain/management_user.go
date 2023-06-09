package domain

import (
	"context"
	"errors"
	"time"
)

var ErrManagementUserNotFound error = errors.New("management user not found")
var ErrDuplicateUserExists error = errors.New("management user already exists")

// This represents a user in the administration database.
// Users in this table will only be able to manage tenants, not licenses.
type ManagementUser struct {
	UserID    int       `db:"user_id" json:"user_id"`
	Username  string    `db:"username" json:"username"`
	ApiKey    string    `db:"api_key" json:"-"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Active    bool      `db:"active" json:"active"`
	Deleted   bool      `db:"deleted" json:"-"`
	DeletedAt time.Time `db:"deleted_at" json:"-"`
	IsAdmin   bool      `db:"is_admin" json:"is_admin"`
}

// ManagementRepository is an interface that defines the methods that must be implemented
// in order to interact with the management user database.
type ManagementRepository interface {
	// Create will create a new management user in the management user database.
	Create(ctx context.Context, user *ManagementUser) error
	// Delete will delete a management user from the management user database.
	Delete(ctx context.Context, username string) error
	// Get will return a management user from the management user database.
	Get(ctx context.Context, username string) (*ManagementUser, error)
	// GetAll will return all of the management users in the management user database.
	GetAll(ctx context.Context) ([]*ManagementUser, error)
}
