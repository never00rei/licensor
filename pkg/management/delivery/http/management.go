package http

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/never00rei/licensor/domain"
	"github.com/never00rei/licensor/pkg/httputils"
	"github.com/never00rei/licensor/pkg/management"
	"golang.org/x/crypto/bcrypt"
)

func ApplyRoutes(r chi.Router, srv *management.ManagementService) {
	// r.Use(authMiddleware(srv))
	r.Post("/", createHandler(srv))
	r.Delete("/{username}", deleteHandler(srv))
	r.Get("/", getAllHandler(srv))
	r.Get("/{username}", getHandler(srv))
}

func createHandler(srv *management.ManagementService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Decode the request
		var user ManagementUserCreateRequest
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "failed to decode request", http.StatusBadRequest)
			return
		}

		managementUser := domain.ManagementUser{
			Username: user.Username,
			Email:    user.Email,
			IsAdmin:  user.IsAdmin,
		}

		// Call the service
		apiKey, err := srv.Create(r.Context(), &managementUser)
		if errors.Is(err, domain.ErrDuplicateUserExists) {
			http.Error(w, "user already exists", http.StatusBadRequest)
			return
		} else if err != nil {
			log.Println(err)
			http.Error(w, "failed to create management user", http.StatusInternalServerError)
			return
		}

		// Create the response
		response := ManagementUserCreateResponse{
			Username: user.Username,
			ApiKey:   apiKey,
		}

		// Write the response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}

func deleteHandler(srv *management.ManagementService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		username := chi.URLParam(r, "username")

		// Call the service
		err := srv.Delete(r.Context(), username)
		if err != nil {
			http.Error(w, "failed to delete management user", http.StatusInternalServerError)
			return
		}

		// Write the response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode("success")
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

	}
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

func getHandler(srv *management.ManagementService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		username := chi.URLParam(r, "username")

		// Call the service
		managementUser, err := srv.Get(r.Context(), username)
		if err != nil {
			http.Error(w, "failed to get management user", http.StatusInternalServerError)
			return
		}

		// Write the response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(managementUser)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}

func AuthMiddleware(srv *management.ManagementService) httputils.Middleware {
	return func(next http.Handler) http.Handler {
		// Get the API token from the header
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Read the API key from the "Authorization" header
			apiKey := r.Header.Get("Authorization")

			token, err := extractToken(apiKey)
			if err != nil {
				log.Println(err)
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			username, err := extractUsername(token)
			if err != nil {
				log.Println(err)
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			log.Println(username)

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

func extractToken(bearerToken string) (string, error) {
	tokenSplit := strings.Split(bearerToken, " ")

	if tokenSplit[0] != "Bearer" {
		log.Println("Invalid authorization header")
		return "", errors.New("invalid authorization header")
	}

	token := strings.TrimSpace(tokenSplit[1])

	return token, nil
}

func extractUsername(token string) (string, error) {
	tokenBytes, err := base64.RawURLEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}

	tokenSplit := strings.Split(string(tokenBytes), ":::")
	if len(tokenSplit) != 2 {
		return "", errors.New("token does not contain username")
	}

	username := strings.TrimSpace(tokenSplit[0])

	return username, nil
}
