package database

import (
	"database/sql"
)

var connectionString string = ""

func InitConnection() (db *sql.DB, err error) {
	//db, err = sql.Open("sqlserver", connectionString)
	db = nil
	err = nil
	return
}
