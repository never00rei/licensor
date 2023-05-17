package licensing

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/never00rei/licensor/models"
)

// Createlicense creates a new license with the given parameters. It will calculate the valid until date based on the valid from date and the validity period which
// should be given in days.
func CreateLicense(issuer, verifier, orgUUID string, validFrom time.Time, validitiyPeriod int) *models.License {
	licenseID := uuid.New().String()
	validUnitl := validFrom.AddDate(0, 0, validitiyPeriod)

	license := &models.License{
		LicenseID:      licenseID,
		Issuer:         issuer,
		Verifier:       verifier,
		OrgUUID:        orgUUID,
		ValidFrom:      validFrom,
		ValidUntil:     validUnitl,
		ValidityPeriod: validitiyPeriod,
		Active:         true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return license
}

func GenerateJWT(license *models.License, key []byte) (string, error) {

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS512)

	// Set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = license.LicenseID
	claims["org_id"] = license.OrgUUID
	claims["iss"] = license.Issuer
	claims["aud"] = []string{license.Issuer}
	claims["iat"] = license.CreatedAt.Unix()
	claims["nbf"] = license.ValidFrom.Unix()
	claims["exp"] = license.ValidUntil.Unix()
	claims["validityperiod"] = license.ValidityPeriod
	claims["active"] = license.Active

	// Sign the token with the secret key
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
