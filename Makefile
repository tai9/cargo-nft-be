network:
	docker network create cargo-nft-network

postgres:
	docker run --name postgres12 --network cargo-nft-network -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root cargo-nft

dropdb:
	docker exec -it postgres12 dropdb cargo-nft

migrations:
	migrate create -ext sql -dir db/migration -format unix $(name)

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/cargo-nft?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/cargo-nft?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/tai9/cargo-nft-be/db/sqlc Store

.PHONY: network postgres createdb dropdb migrations migrateup migratedown sqlc test server mock