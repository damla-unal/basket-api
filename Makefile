createdb:
	createdb basket-api

dropdb:
	dropdb basket-api

sqlc:
	sqlc generate

migrateup:
	migrate -path ./db/migration -database "postgresql://localhost:5432/basket-api?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://localhost:5432/basket-api?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown sqlc