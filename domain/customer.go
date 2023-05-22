package domain

import "time"

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
