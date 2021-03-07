package pgtypes

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib" // load driver
)

type fatalistic interface {
	Fatal(args ...interface{})
}

func openTestConnConninfo(conninfo string) (*sql.DB, error) {
	return sql.Open("pgx", conninfo)
}

func openTestConn(t fatalistic) *sql.DB {
	conn, err := openTestConnConninfo("")
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func getServerVersion(t *testing.T, db *sql.DB) int {
	var version int
	err := db.QueryRow("SHOW server_version_num").Scan(&version)
	if err != nil {
		t.Fatal(err)
	}
	return version
}
