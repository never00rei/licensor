package main

import (
	"log"
	"os"
	"time"

	"github.com/never00rei/licensor/pkg/licensing"
)

func main() {

	// Generate License
	license := licensing.CreateLicense("issuer", "verifier", "orgUUID", time.Now(), 30)
	key := []byte("key")

	// Generate JWT
	jwt, err := licensing.GenerateJWT(license, key)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Println(jwt)

}
