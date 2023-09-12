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

var testingQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSrc)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testingQueries = New(conn)

	os.Exit(m.Run())
}
