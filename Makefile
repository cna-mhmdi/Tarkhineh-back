DB_URL=postgresql://root:nyco@localhost:5432/Tarkhineh-db?sslmode=disable

postgres:
	docker run --name Tarkhineh-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=nyco -d postgres:12-alpine

createdb:
	docker exec -it Tarkhineh-db createdb --username=root --owner=root Tarkhineh-db

dropdb:
	docker exec -it Tarkhineh-db dropdb Tarkhineh-db

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/cna-mhmdi/Tarkhineh-back/db/sqlc Store

test:
	go test -v -cover -short ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown mock test server