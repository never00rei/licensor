package keygen

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func GenerateApiKey(username string) (string, error) {
	// Create a byte slice of the specified length
	byteSlice := make([]byte, 32)
	usernameBytes := []byte(username)

	randomBytes := append(usernameBytes, byteSlice...)

	// Read random data into the byte slice
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64
	randomString := base64.RawURLEncoding.EncodeToString(randomBytes)

	// Create a new HMAC-SHA256 hasher
	h := sha512.New()

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
