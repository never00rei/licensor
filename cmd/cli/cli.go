package main

import (
	"context"
	"flag"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/never00rei/licensor/domain"
	"github.com/never00rei/licensor/pkg/dbconfig"
	"github.com/never00rei/licensor/pkg/management"
	managementRepo "github.com/never00rei/licensor/pkg/management/repository/postgresql"
)

var flagUsername = flag.String("username", "", "Please provide a username")
var flagEmail = flag.String("email", "", "Please provide an email")

func main() {
	flag.Parse()

	if *flagUsername == "" || *flagEmail == "" {
		flag.PrintDefaults()
		return
	}

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

	managementRepo := managementRepo.NewPostgresqlManagementRepo(conn)

	// Generate the service
	managementService := management.NewManagementService(managementRepo)

	user := domain.ManagementUser{
		Username: *flagUsername,
		Email:    *flagEmail,
		IsAdmin:  true,
	}

	apiKey, err := managementService.Create(ctx, &user)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Username: %s\nAPI Key: %s", *flagUsername, apiKey)

	// Create a new management user
	// Call the service
	// Write the response
}
