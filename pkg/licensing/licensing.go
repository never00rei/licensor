package licensing

import "time"

type License struct {
	Issuer           string
	ValidFrom        time.Time
	ValidUntil       time.Time
	ExpiryDate       time.Time
	ValidityPeriod   int
	OrganizationName string
	LicenseID        string
}
