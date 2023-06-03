package main

import "github.com/never00rei/licensor/pkg/keygen"

func main() {

	key, err := keygen.GenerateKey("foo")
	if err != nil {
		panic(err)
	}

	println(key)

	// config, err := dbconfig.GetDBConfigFromEnv()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// ctx := context.Background()

	// conn, err := pgxpool.New(ctx, config.GetConnectionURL())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer conn.Close()

	// // Create a tenant repo
	// tenantRepo := tenantRepo.NewPostgresqlTenantRepo(conn)

	// tenantService := tenant.NewTenantService(tenantRepo)

	// tenants, err := tenantService.GetAll(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Tenants: ", tenants)
	// for _, t := range tenants {
	// 	log.Println(t.OrgID, t.OrgUUID, t.OrgName)
	// }

	// managementUserRepo := managementUserRepo.NewPostgresqlManagementRepo(conn)

	// managementService := management.NewManagementService(managementUserRepo)

	// managementUsers, err := managementService.GetAll(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Management Users: ", managementUsers)
	// for _, m := range managementUsers {
	// 	log.Println(m.UserID, m.Username, m.Email)
	// }

}
