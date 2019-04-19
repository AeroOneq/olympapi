package database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"../models"
)

type ErrorDB struct {
	Database *sql.DB
}

func (errorDB *ErrorDB) OpenConnection() (err error) {
	db, err := InitConnection()
	if err == nil {
		errorDB.Database = db
	}

	return
}

func (errorDB *ErrorDB) WriteRequestErrorToDB(reqErr *models.RequestError) {
	reqErrBytes, err := json.Marshal(reqErr)

	if err == nil {
		fmt.Println(string(reqErrBytes))
	}
}
