# gingoadmin

## PostgreSQL

```bash
# initialise database and user
psql -U postgres -f psql/init.sql
# set up goadmin's tables with database and user
psql -U gingoadmin_user -d gingoadmin_db -f psql/admin.sql
```

## Run gin go-admin

```bash
go get -v ./...
go run main.go
```

## Running tests

```bash
go test -v ./...
```

## Using the go-admin's cli program

```bash
# installs `adm` cli program into /usr/local/bin
GOBIN=/usr/local/bin go install github.com/GoAdminGroup/go-admin/adm
```
