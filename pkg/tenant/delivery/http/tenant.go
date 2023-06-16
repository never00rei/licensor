package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/never00rei/licensor/domain"
	"github.com/never00rei/licensor/pkg/tenant"
)

func ApplyRoutes(r chi.Router, srv *tenant.TenantService) {
	r.Get("/", getAllHandler(srv))
	r.Post("/", createHandler(srv))
}

func createHandler(srv *tenant.TenantService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var tenantRequest TenantCreateRequest
		err := json.NewDecoder(r.Body).Decode(&tenantRequest)
		if err != nil {
			http.Error(w, "failed to decode request", http.StatusBadRequest)
			return
		}

		tenant := &domain.Tenant{
			OrgName: tenantRequest.OrgName,
		}

		// Call the service
		err = srv.Create(r.Context(), tenant)
		if err != nil {
			log.Println(err)
			http.Error(w, "failed to create tenant", http.StatusInternalServerError)
			return
		}

		response := TenantCreateResponse{
			OrgName: tenant.OrgName,
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

func getAllHandler(srv *tenant.TenantService) http.HandlerFunc {

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
