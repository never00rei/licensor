package keygen

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

// GenerateKey generates a HS512 key for signing JWTs.
// This is a simple wrapper around the GenerateKey function in the jwt package.
func GenerateKey(secretKey string) (string, error) {

	// Create a byte slice of the specified length
	randomBytes := make([]byte, 32)

	// Read random data into the byte slice
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64
	randomString := base64.RawURLEncoding.EncodeToString(randomBytes)

	// Create a new HMAC-SHA256 hasher
	h := hmac.New(sha512.New, []byte(secretKey))

	// Write the data to the hasher
	_, err = h.Write([]byte(randomString))
	if err != nil {
		return "", err
	}

	// Get the raw HMAC-SHA512 sum
	sum := h.Sum(nil)

	// Encode the raw sum to base64
	token := base64.RawURLEncoding.EncodeToString(sum)

	return token, nil
}
