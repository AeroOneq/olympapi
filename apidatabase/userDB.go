package apidatabase

import (
	"database/sql"

	"OlympApi/models"
)

type UserDB struct {
	Database *sql.DB
}

func (userDB *UserDB) OpenConnection() (err error) {
	db, err := InitConnection()
	if err == nil {
		userDB.Database = db
		err = nil
	}

	return
}

func (userDB *UserDB) GetUserByID(id int) (user *models.User, err error) {
	user = &(models.User{})
	err = nil

	return
}

func (userDB *UserDB) GetAllUsers() (users *[]models.User, err error) {
	users = &[]models.User{models.User{}, models.User{}}
	err = nil

	return
}
