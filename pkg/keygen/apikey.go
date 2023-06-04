package keygen

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateApiKey(username string) (string, error) {
	// Create a byte slice of the specified length
	byteSlice := make([]byte, 32)
	usernameBytes := []byte(fmt.Sprintf("%s:::", username))

	// Read random data into the byte slice
	_, err := rand.Read(byteSlice)
	if err != nil {
		return "", err
	}

	randomBytes := append(usernameBytes, byteSlice...)

	// Base64 encode the random bytes
	token := base64.RawURLEncoding.EncodeToString(randomBytes)

	return token, nil
}

func SaltAndHashAPIKey(key string) (string, error) {
	hash := sha256.Sum256([]byte(key))
	apiKeyHash := hash[:]
	bcryptKey, err := bcrypt.GenerateFromPassword(apiKeyHash, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bcryptKey), nil
}
