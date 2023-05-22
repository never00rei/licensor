package domain

import "time"

// This represents a user in the tenant database.
// Users in this table will only be able to manage licenses, not tenants.
type User struct {
	UserID    string    `db:"user_id"`
	OrgUUID   string    `db:"org_uuid"`
	Username  string    `db:"username"`
	ApiKey    string    `db:"api_key"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Active    bool      `db:"active"`
	LastLogin time.Time `db:"last_login"`
	Deleted   bool      `db:"deleted"`
	DeletedAt time.Time `db:"deleted_at"`
}
