postgres:
	docker run --name bank-postgres -p 5432:5432 -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=mysecret -d postgres

createdb:
	docker exec -it bank-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it bank-postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

unit-tests:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown unit-tests sqlc