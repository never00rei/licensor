package main

import (
	"crypto/sha256"
	"log"

	"github.com/never00rei/licensor/pkg/keygen"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	// key, err := keygen.GenerateApiKey("admin")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println(key)

	key := "YWRtaW46Ojokxu6dGsK2kVf70VEmzGyIWGdVkCQI7YnuS6lQGHJD3Q"

	dbKey, err := keygen.SaltAndHashAPIKey(key)
	if err != nil {
		log.Fatal(err)
	}

	hash := sha256.Sum256([]byte(key))
	apiKeyHash := hash[:]

	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(dbKey)
	err = bcrypt.CompareHashAndPassword(byteHash, apiKeyHash)
	if err != nil {
		log.Println(err)
	}

	log.Println(dbKey)

}
