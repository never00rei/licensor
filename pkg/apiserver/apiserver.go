package apiserver

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/never00rei/licensor/pkg/management"
	managementDelivery "github.com/never00rei/licensor/pkg/management/delivery/http"
	managementRepo "github.com/never00rei/licensor/pkg/management/repository/postgresql"

	"github.com/never00rei/licensor/pkg/tenant"
	tenantDelivery "github.com/never00rei/licensor/pkg/tenant/delivery/http"
	tenantRepo "github.com/never00rei/licensor/pkg/tenant/repository/postgresql"
)

type Server struct {
	Router *chi.Mux
}

// NewServer will create a new Server object
func NewServer(pool *pgxpool.Pool) *Server {

	// Generate base router
	baseRouter := chi.NewRouter()

	baseRouter.Use(middleware.Timeout(60 * time.Second))

	// A good base middleware stack
	baseRouter.Use(middleware.RequestID)
	baseRouter.Use(middleware.RealIP)
	baseRouter.Use(middleware.Logger)
	baseRouter.Use(middleware.Recoverer)

	managementRepo := managementRepo.NewPostgresqlManagementRepo(pool)

	// Generate the service
	managementService := management.NewManagementService(managementRepo)

	baseRouter.Route("/admin/user", func(r chi.Router) {
		managementDelivery.ApplyRoutes(r, managementService)
	})

	tenantRepo := tenantRepo.NewPostgresqlTenantRepo(pool)

	// Generate the service
	tenantService := tenant.NewTenantService(tenantRepo)

	baseRouter.Route("/admin/tenant", func(r chi.Router) {
		tenantDelivery.ApplyRoutes(r, tenantService)
	})

	return &Server{
		Router: baseRouter,
	}
}

func (s *Server) Start() {
	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", s.Router)
}
