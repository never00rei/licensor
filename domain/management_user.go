package domain

import "time"

// This represents a user in the administration database.
// Users in this table will only be able to manage tenants, not licenses.
type ManagementUser struct {
	UserID    int       `db:"user_id"`
	Username  string    `db:"username"`
	ApiKey    string    `db:"api_key"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Active    bool      `db:"active"`
	Deleted   bool      `db:"deleted"`
	DeletedAt time.Time `db:"deleted_at"`
	IsAdmin   bool      `db:"is_admin"`
}
