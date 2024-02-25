run:
	go run .

newcontainer:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:latest
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres postgres
dropdb:
	docker exec -it postgres dropdb postgres
dockerstart:
	docker start postgres
dockerstop:
	docker stop postgres_con

createmigrate:
	migrate create -ext sql -dir db/migrations -seq init_schema

migrationup:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migrationdown:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

pull:
	git pull origin main

push:
	git push origin main

sqlc:
	docker run --rm -v D:/golang/simplebank:/src -w /src sqlc/sqlc generate

dmigration:
	docker run --rm -v D:/golang/simplebank/db/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable up

dmigrationd:
	docker run --rm -v D:/golang/simplebank/db/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable down 1

.PHONY: run newcontainer createdb dropdb dockerstart dockerstop migrationup migrationdown pull push sqlc