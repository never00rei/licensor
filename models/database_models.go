package models

import "time"

// Administration Database Models.
//
// Tenant is a struct that represents a tenant in the database.
// Think of the "tenant" table as part of a top level administration instance.
// This will be in it's own database, and will maintain a list of tenants along
// with their associated database table names.
//
// This is to allow for multiple tenants to be stored in the same database server,
// but allow tenant data to be stored in seperate databases, providing data isolation.
// However, this adds complexity as we need to change the database connection at runtime...
type Tenant struct {
	OrgID     int       `db:"org_id"`
	OrgName   string    `db:"org_name"`
	OrgUUID   string    `db:"org_uuid"`
	TableName string    `db:"table_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// This represents a user in the administration database.
// Users in this table will only be able to manage tenants, not licenses.
type ManagementUser struct {
	UserID    int       `db:"user_id"`
	OrgUUID   string    `db:"org_uuid"`
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

// Tenant database models.

// This represents a key value which will be used to validate and sign JWTs.
type Key struct {
	KeyID   int    `db:"key_id"`
	OrgUUID string `db:"org_uuid"`
	Key     []byte `db:"key"`
}

// This represents a license in the tenant database.
type License struct {
	LicenseID      string    `db:"license_id"`
	Issuer         string    `db:"issuer"`
	Verifier       string    `db:"verifier"`
	OrgUUID        string    `db:"org_uuid"`
	ValidFrom      time.Time `db:"valid_from"`
	ValidUntil     time.Time `db:"valid_until"`
	ValidityPeriod int       `db:"validity_period"`
	Active         bool      `db:"active"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

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

// This represents a customer in the tenant database.
// Customers are the end users of the licenses provided by the tenant.
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
