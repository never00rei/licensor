package models

import "time"

// Tenant is a struct that represents a tenant in the database.
// This is also split into separate schemas in the database.
// This is to allow for multiple tenants to be stored in the same database server,
// but allow tenant data to be stored in seperate databases providing isolation.
type Tenant struct {
	OrgID     int       `db:"org_id"`
	OrgName   string    `db:"org_name"`
	OrgUUID   string    `db:"org_uuid"`
	TableName string    `db:"table_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// Tenant database models.
type License struct {
	LicenseID      string    `db:"license_id"`
	OrgUUID        string    `db:"org_uuid"`
	ValidFrom      time.Time `db:"valid_from"`
	ValidUntil     time.Time `db:"valid_until"`
	ValidityPeriod int       `db:"validity_period"`
	Active         bool      `db:"active"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type User struct {
	UserID    string    `db:"user_id"`
	OrgUUID   string    `db:"org_uuid"`
	Username  string    `db:"username"`
	ApiKey    string    `db:"api_key"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Active    bool      `db:"active"`
	lastLogin time.Time `db:"last_login"`
	Deleted   bool      `db:"deleted"`
	DeletedAt time.Time `db:"deleted_at"`
}

type Customer struct {
	CustomerID string    `db:"customer_id"`
	OrgUUID    string    `db:"org_uuid"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	Active     bool      `db:"active"`
	Deleted    bool      `db:"deleted"`
	DeletedAt  time.Time `db:"deleted_at"`
}
