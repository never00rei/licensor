package licensing_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/never00rei/licensor/domain"
	"github.com/never00rei/licensor/pkg/licensing"
	"github.com/stretchr/testify/assert"
)

var EXAMPLEKEY = []byte("xdsMO3GMXGTcfnX8h88oC2t_fg3sxTHA22AyajQQ_P3jRqr71RnwzKBLtBriAGuWS5Rw87lDoWy8K8xTD18OHg")

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
			license := licensing.NewLicense(tc.issuer, tc.verifier, tc.orgUUID, tc.validFrom, tc.validitiyPeriod)

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

func TestGenerateJWT(t *testing.T) {

	testcases := map[string]struct {
		licenseID       string
		issuer          string
		verifier        string
		orgUUID         string
		createdAt       time.Time
		validUntil      time.Time
		validFrom       time.Time
		validitiyPeriod int
		active          bool
		key             []byte
	}{
		"success": {
			licenseID:       "foo",
			issuer:          "foo",
			verifier:        "bar",
			orgUUID:         "foobar",
			createdAt:       time.Now(),
			validFrom:       time.Now(),
			validUntil:      time.Now().AddDate(0, 0, 30),
			validitiyPeriod: 30,
			active:          true,
			key:             EXAMPLEKEY,
		},
	}

	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {

			license := &domain.License{
				LicenseID:      tc.licenseID,
				Issuer:         tc.issuer,
				Verifier:       tc.verifier,
				OrgUUID:        tc.orgUUID,
				CreatedAt:      tc.createdAt,
				ValidFrom:      tc.validFrom,
				ValidUntil:     tc.validUntil,
				ValidityPeriod: tc.validitiyPeriod,
				Active:         tc.active,
			}

			// Generate the JWT
			jwtVal, err := licensing.GenerateJWT(license, tc.key)
			assert.NoError(t, err)

			jwtToken, err := jwt.Parse(jwtVal, func(token *jwt.Token) (interface{}, error) {
				return tc.key, nil
			})
			assert.NoError(t, err)

			// Generate the Signing String to later verify the signature
			signingString, err := jwtToken.SigningString()
			assert.NoError(t, err)

			// Verify the signature
			err = jwt.SigningMethodHS512.Verify(signingString, jwtToken.Signature, tc.key)
			assert.NoError(t, err)

			assert.Equal(t, tc.issuer, jwtToken.Claims.(jwt.MapClaims)["iss"])
			assert.Equal(t, license.LicenseID, jwtToken.Claims.(jwt.MapClaims)["sub"])
			assert.Equal(t, tc.orgUUID, jwtToken.Claims.(jwt.MapClaims)["org_id"])
			assert.Equal(t, []interface{}{tc.issuer}, jwtToken.Claims.(jwt.MapClaims)["aud"])
			assert.Equal(t, float64(license.CreatedAt.Unix()), jwtToken.Claims.(jwt.MapClaims)["iat"])
			assert.Equal(t, float64(license.ValidFrom.Unix()), jwtToken.Claims.(jwt.MapClaims)["nbf"])
			assert.Equal(t, float64(license.ValidUntil.Unix()), jwtToken.Claims.(jwt.MapClaims)["exp"])
			assert.Equal(t, float64(license.ValidityPeriod), jwtToken.Claims.(jwt.MapClaims)["validityperiod"])
			assert.Equal(t, license.Active, jwtToken.Claims.(jwt.MapClaims)["active"])

		})
	}

}
