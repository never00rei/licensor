package domain

import "time"

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
