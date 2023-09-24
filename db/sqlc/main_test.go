package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"swiftiesoft.com/simplebank/utils"
)

var testingQueries *Queries // ? to access globally
var testDB *sql.DB          // ? to access globally

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSrc)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testingQueries = New(testDB)

	os.Exit(m.Run())
}
