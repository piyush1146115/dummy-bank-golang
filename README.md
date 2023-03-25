# dummy-bank-golang


## Steps:
- **Create DB_schema:** [db-schema](./db_schema/schema.md)
- **Install Postgres:** [docker hub](https://hub.docker.com/_/postgres)
- **Install TablePlus:** [TablePlus](https://tableplus.com/blog/2019/10/tableplus-linux-installation.html)
- Create a directory for db migration `db/migration`
- Create initial migration schema `migrate create -ext sql -dir db/migration -seq init_schema`
