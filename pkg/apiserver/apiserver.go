package apiserver

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/never00rei/licensor/pkg/config"
	"github.com/never00rei/licensor/pkg/management"
	managementDelivery "github.com/never00rei/licensor/pkg/management/delivery/http"
	managementRepo "github.com/never00rei/licensor/pkg/management/repository/postgresql"

	"github.com/never00rei/licensor/pkg/tenant"
	tenantDelivery "github.com/never00rei/licensor/pkg/tenant/delivery/http"
	tenantRepo "github.com/never00rei/licensor/pkg/tenant/repository/postgresql"
)

type Server struct {
	server *http.Server
	config *config.AppConfig
	Router *chi.Mux
}

// NewServer will create a new Server object
func NewServer(pool *pgxpool.Pool, config *config.AppConfig) *Server {

	s := &Server{config: config}

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

	s.Router = baseRouter

	// Set the http config
	s.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler:      s.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s
}

func (s *Server) Start() {
	log.Printf("Starting server on %s:%d", s.config.Host, s.config.Port)
	if err := s.server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
