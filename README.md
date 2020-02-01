# gingoadmin

## PostgreSQL

```bash
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

## Running tests

```bash
go test -v ./...
```