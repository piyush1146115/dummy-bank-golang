postgres:
	docker run --name bank-postgres -p 5432:5432 -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=mysecret -d postgres

remove-postgres:
	docker stop bank-postgres && docker rm bank-postgres

createdb:
	docker exec -it bank-postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it bank-postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" --verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" --verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate

unit-tests:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/piyush1146115/dummy-bank-golang/db/sqlc Store

migrate-create:
	migrate create -ext sql -dir db/migration -seq add_users

.PHONY: postgres createdb dropdb migrateup migratedown unit-tests sqlc remove-postgres server mock migrate-create migrateup1 migratedown1