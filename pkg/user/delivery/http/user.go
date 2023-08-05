package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/never00rei/licensor/domain"
	"github.com/never00rei/licensor/pkg/user"
)

func ApplyRoutes(r chi.Router, srv *user.UserService) {
	r.Get("/", getAllHandler(srv))
	r.Post("/", createHandler(srv))
}

func createHandler(srv *user.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		tenantID := chi.URLParam(r, "tenant_id")

		var userRequest UserCreateRequest
		err := json.NewDecoder(r.Body).Decode(&userRequest)
		if err != nil {
			http.Error(w, "failed to decode request", http.StatusBadRequest)
			return
		}

		user := &domain.User{
			OrgUUID:  tenantID,
			Email:    userRequest.Email,
			Username: userRequest.Username,
		}

		// Call the service
		apiKey, err := srv.Create(r.Context(), user)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}

		response := UserCreateResponse{
			Username: user.Username,
			Email:    user.Email,
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

func getAllHandler(srv *user.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		tenantID := chi.URLParam(r, "tenant_id")

		// Call the service
		users, err := srv.GetAll(r.Context(), tenantID)
		if err != nil {
			http.Error(w, "failed to get users", http.StatusInternalServerError)
			return
		}

		// Write the response
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}
