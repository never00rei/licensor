package http

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/never00rei/licensor/pkg/httputils"
	"github.com/never00rei/licensor/pkg/management"
	"golang.org/x/crypto/bcrypt"
)

func ApplyRoutes(r chi.Router, srv *management.ManagementService) {
	r.Use(authMiddleware(srv))
	r.Get("/", getAllHandler(srv))
}

func getAllHandler(srv *management.ManagementService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Call the service
		managementUsers, err := srv.GetAll(r.Context())
		if err != nil {
			http.Error(w, "failed to get management users", http.StatusInternalServerError)
			return
		}

		// Write the response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(managementUsers)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}

func authMiddleware(srv *management.ManagementService) httputils.Middleware {
	return func(next http.Handler) http.Handler {
		// Get the API token from the header
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Read the API key from the "Authorization" header
			apiKey := r.Header.Get("Authorization")

			apiKeySplit := strings.Split(apiKey, " ")

			if apiKeySplit[0] != "Bearer" {
				log.Println("Invalid authorization header")
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			// Check if the API key is valid
			token := strings.TrimSpace(apiKeySplit[1])

			// Base64 Decode the token
			tokenBytes, err := base64.RawURLEncoding.DecodeString(token)
			if err != nil {
				log.Println(err)
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			// extract username from token
			tokenSplit := strings.Split(string(tokenBytes), ":::")
			if len(tokenSplit) != 2 {
				log.Println("Token does not contain username")
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			username := strings.TrimSpace(tokenSplit[0])

			// Get user from db
			user, err := srv.Get(r.Context(), username)
			if err != nil {
				log.Println(err)
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			hash := sha256.Sum256([]byte(token))
			apiKeyHash := hash[:]

			err = bcrypt.CompareHashAndPassword([]byte(user.ApiKey), apiKeyHash)
			if err != nil {
				log.Println(err)
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
