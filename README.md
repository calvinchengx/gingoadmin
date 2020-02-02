# gingoadmin

## PostgreSQL

```bash
# load postgresql environment variables
cp .env.sample .env
cd .

# initialise database and user
psql -U postgres -f psql/init.sql
# set up goadmin's tables with database and user
psql -U gingoadmin_user -d gingoadmin_db -f psql/admin.sql
```

## Using the go-admin's cli program

```bash
# installs `adm` cli program into /usr/local/bin
GOBIN=/usr/local/bin go install github.com/GoAdminGroup/go-admin/adm
```

## Using air for hot reload

```bash
GOBIN=/usr/local/bin go install github.com/cosmtrek/air
```

## Run gin go-admin

```bash
go get -v ./...
# go run .
air
```

## Managing postgresql tables and models

```bash
# model generator from postgresql tables
GOBIN=/usr/local/bin go install github.com/dizzyfool/genna

# postgresql client and orm
go get github.com/go-pg/pg/v9

# useful for translating between url/API params and model queries, via filter
go get github.com/go-pg/urlstruct
```

## Generating structs (models) from existing postgresql tables

```bash
bash -c "genna model-named -c postgres://gingoadmin_user:gingoadmin_password@localhos
t:5432/gingoadmin_db?sslmode=disable -o ./tmp/model.go"
```

## Schema migrations

```bash
# create [migration_name]
# migrate
# rollback
# help
go run . [command]
```

## Running tests

```bash
go test -v ./...
```