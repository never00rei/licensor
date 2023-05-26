package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/never00rei/licensor/pkg/dbconfig"
	"github.com/never00rei/licensor/pkg/tenant"
	"github.com/never00rei/licensor/pkg/tenant/repository/postgresql"
)

func main() {

	config, err := dbconfig.GetDBConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	conn, err := pgxpool.New(ctx, config.GetConnectionURL())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// Create a tenant repo
	tenantRepo := postgresql.NewPostgresqlTenantRepo(conn)

	tenantService := tenant.NewTenantService(tenantRepo)

	tenants, err := tenantService.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(tenants)
	for _, t := range tenants {
		log.Println(t.OrgUUID, t.OrgName)
	}
}
