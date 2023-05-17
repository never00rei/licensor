package licensing

import "time"

type License struct {
	Issuer           string    `json:"iss"`
	IssuedAt         time.Time `json:"iat"`
	ValidFrom        time.Time `json:"nbf"`
	ValidUntil       time.Time `json:"exp"`
	ValidityPeriod   int       `json:"validity_period"`
	OrganizationName string    `json:"org_name"`
	LicenseID        string    `json:"license_id"`
}
