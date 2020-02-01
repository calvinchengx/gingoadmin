# gingoadmin

## PostgreSQL

```bash
# initialise database and user
psql -U postgres -f psql/init.sql
# set up goadmin's tables with database and user
psql -U gingoadmin_user -d gingoadmin_db -f psql/admin.sql
```
