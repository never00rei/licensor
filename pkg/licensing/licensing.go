package licensing

import "time"

type License struct {
	Issuer           string
	ValidFrom        time.Time
	ValidUntil       time.Time
	OrganizationName string
	LicenseID        string
}
