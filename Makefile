run:
	go run .

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

.PHONY: run dockerstart dockerstop migrationup migrationdown