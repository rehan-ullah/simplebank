run:
	go run .

newcontainer:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:latest
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank
dockerstart:
	sudo docker start postgres_con

dockerstop:
	sudo docker stop postgres_con

createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrationup:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migrationdown:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

pull:
	git pull origin main

push:
	git push origin main

sqlc:
	docker run --rm -v D:/golang/simplebank:/src -w /src sqlc/sqlc generate

.PHONY: run dockerstart dockerstop migrationup migrationdown pull push sqlc