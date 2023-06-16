package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/never00rei/licensor/pkg/apiserver"
	"github.com/never00rei/licensor/pkg/config"
)

func main() {

	dbconf := config.NewDefaultDBConfig().ApplyOptions(
		config.WithDBHost(os.Getenv("DB_HOST")),
		config.WithDBPort(os.Getenv("DB_PORT")),
		config.WithDBUser(os.Getenv("DB_USER")),
		config.WithDBPassword(os.Getenv("DB_PASSWORD")),
		config.WithDBDatabase(os.Getenv("DB_DATABASE")),
	)

	ctx := context.Background()

	conn, err := pgxpool.New(ctx, dbconf.GetConnectionURL())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	cfg := config.NewDefaultAppConfig().ApplyOptions(
		config.WithHost(""),
		config.WithPort(8080),
	)

	// Create Server
	server := apiserver.NewServer(conn, cfg)

	server.Start()

}
