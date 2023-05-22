# Licensor

This is an opensource SAAS licensing system.


## Important information


### Database Migrations
This project uses Tern to migrate database schemas.

To install tern:

1. git clone https://github.com/jackc/tern
2. Then build tern and move it into `/usr/local/bin` or the mac equivilent.
3. Once installed you can run tern from the migrations directory, preferably one of the sub folders.


### Setting up Postgres

1. Run `docker-compose up -d`
2. To login `psql -U postgres -h localhost -p 5432 postgres` use the username as the password.
3. In the PSQL command line run `CREATE DATABASE licensor_management;`
4. Once complete run `\c licensor_management`
5. Then run `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";` to enable UUID v4 auto generation.

### Migrating the management database schema

1. set the following environment variables (use the username as the password, replacing "CHANGEME" below):
```
export DB_HOST="localhost"
export DB_PASSWORD="CHANGEME"
export DB_USER="postgres"
export DB_PORT="5432"
export DB_DATABASE="licensor_management"

export MIGRATOR_USER="postgres"
export MIGRATOR_PASSWORD="CHANGEME"
```

2. `cd migrations/management_db && tern migrate`
