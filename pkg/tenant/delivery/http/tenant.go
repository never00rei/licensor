package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/never00rei/licensor/pkg/tenant"
)

func ApplyRoutes(r chi.Router, srv *tenant.TenantService) {
	r.Get("/", getAllHandler(srv))
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
