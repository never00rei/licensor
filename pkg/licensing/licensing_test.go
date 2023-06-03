package licensing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateLicence(t *testing.T) {

	testcases := map[string]struct {
		issuer          string
		verifier        string
		orgUUID         string
		validFrom       time.Time
		validitiyPeriod int
	}{
		"success": {
			issuer:          "foo",
			verifier:        "bar",
			orgUUID:         "foobar",
			validFrom:       time.Now(),
			validitiyPeriod: 30,
		},
		"successValidFromSetDate": {
			issuer:          "foo",
			verifier:        "bar",
			orgUUID:         "foobar",
			validFrom:       time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
			validitiyPeriod: 30,
		},
		"successNoValidityPeriod": {
			issuer:          "foo",
			verifier:        "bar",
			orgUUID:         "foobar",
			validFrom:       time.Now(),
			validitiyPeriod: 0,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {

			// Generate the license using the factory
			license := NewLicense(tc.issuer, tc.verifier, tc.orgUUID, tc.validFrom, tc.validitiyPeriod)

			// Check the contents of the license
			assert.Equal(t, tc.issuer, license.Issuer)
			assert.Equal(t, tc.verifier, license.Verifier)
			assert.Equal(t, tc.orgUUID, license.OrgUUID)
			assert.Equal(t, tc.validFrom, license.ValidFrom)

			assert.Equal(t, tc.validitiyPeriod, license.ValidityPeriod)

			// Check the timestamps are within a second of now as we need to give time for the tests to actually
			// run
			assert.WithinDuration(t, tc.validFrom.AddDate(0, 0, tc.validitiyPeriod), license.ValidUntil, time.Second)
			assert.WithinDuration(t, time.Now(), license.CreatedAt, time.Second)
			assert.WithinDuration(t, time.Now(), license.UpdatedAt, time.Second)

		})
	}

}
