package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSrc    = "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

var testingQueries *Queries // ? to access globally
var testDB *sql.DB          // ? to access globally

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSrc)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testingQueries = New(testDB)

	os.Exit(m.Run())
}
