package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/never00rei/licensor/pkg/apiserver"
	"github.com/never00rei/licensor/pkg/dbconfig"
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

	// Create Server
	server := apiserver.NewServer(conn)

	server.Start()

}
