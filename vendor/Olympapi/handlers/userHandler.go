package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"OlympApi/database"
	"OlympApi/models"

	"github.com/go-martini/martini"
)

func GetRequestObject(rq *http.Request) *models.Request {
	var request *models.Request = &models.Request{}
	(*request).GetRequest(*rq)

	return request
}

func GetUserDatabase() *database.UserDB {
	userDB := database.UserDB{}
	userDB.OpenConnection()
	return &userDB
}

func GetUserJsonByID(rq *http.Request, params martini.Params) (int, string) {
	request := GetRequestObject(rq)

	if len(params) != 1 {
		HandleError(request, "NOT APPROPRIATE PARAMS LENGTH, WANT 1")
		return http.StatusInternalServerError, INTERNAL_ERROR
	} else {
		id, ok := params["id"]

		if ok {
			intId, convErr := strconv.Atoi(id)

			if convErr == nil {
				userDB := GetUserDatabase()
				user, err := userDB.GetUserByID(intId)

				if err == nil {
					userBytes, jsonErr := json.Marshal(*user)
					if jsonErr == nil {
						return http.StatusOK, string(userBytes)
					}
				} else {
					HandleError(request, "USER NOT FOUND")
					return http.StatusNotFound, INTERNAL_ERROR
				}
			} else {
				HandleError(request, "ERROR WHILE PARSING ID TO INT")
				return http.StatusNotFound, INTERNAL_ERROR
			}
		}

		HandleError(request, "ERROR WHILE SERIALIZING OBJECT TO JSON")
		return http.StatusInternalServerError, INTERNAL_ERROR
	}
}

func GetAllUsersJson(rq *http.Request, params martini.Params) (int, string) {
	request := GetRequestObject(rq)
	defer database.WriteRequestToDB(request)

	userDB := GetUserDatabase()
	users, err := userDB.GetAllUsers()

	if err != nil {
		HandleError(request, "USER DATABASE ERROR")
		return http.StatusInternalServerError, INTERNAL_ERROR
	}

	usersBytes, jsonError := json.Marshal(users)
	if jsonError != nil {
		HandleError(request, "JSON SERIALIZING ERROR")
		return http.StatusInternalServerError, INTERNAL_ERROR
	}

	return http.StatusOK, string(usersBytes)
}

func CreateNewUserRecord(rq *http.Request, params martini.Params) (int, string) {
	request := GetRequestObject(rq)
	defer database.WriteRequestToDB(request)

	var newUser models.User
	err := json.Unmarshal(([]byte)((*request).Body), &newUser)

	if err != nil {
		HandleError(request, "UNPROCESSABLE ENTITY WAS GIVEN TO THE SERVER")
		return http.StatusUnprocessableEntity, UNPROCESSABLE_ENTITY
	}

	return http.StatusOK, ""
}
