package database

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

var connectionString string = ""

func InitConnection() (db *sql.DB, err error) {
	//db, err = sql.Open("sqlserver", connectionString)
	db = nil
	err = nil
	return
}
