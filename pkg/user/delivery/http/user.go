package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/never00rei/licensor/pkg/user"
)

func ApplyRoutes(r chi.Router, srv *user.UserService) {
	r.Get("/", getAllHandler(srv))
	// r.Post("/", createHandler(srv))
}

// func createHandler(srv *tenant.TenantService) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		var tenantRequest TenantCreateRequest
// 		err := json.NewDecoder(r.Body).Decode(&tenantRequest)
// 		if err != nil {
// 			http.Error(w, "failed to decode request", http.StatusBadRequest)
// 			return
// 		}

// 		tenant := &domain.Tenant{
// 			OrgName: tenantRequest.OrgName,
// 		}

// 		// Call the service
// 		err = srv.Create(r.Context(), tenant)
// 		if err != nil {
// 			log.Println(err)
// 			http.Error(w, "failed to create tenant", http.StatusInternalServerError)
// 			return
// 		}

// 		response := TenantCreateResponse{
// 			OrgName: tenant.OrgName,
// 		}

// 		// Write the response
// 		w.Header().Add("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusCreated)
// 		err = json.NewEncoder(w).Encode(response)
// 		if err != nil {
// 			http.Error(w, "failed to encode response", http.StatusInternalServerError)
// 			return
// 		}

// 	}
// }

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
