package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"swiftiesoft.com/simplebank/api"
	db "swiftiesoft.com/simplebank/db/sqlc"
	"swiftiesoft.com/simplebank/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSrc)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
