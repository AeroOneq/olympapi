package handlers

import "../models"
import "../database"

func HandleError(rq *models.Request, errorStr string) {
	reqErr := &models.RequestError{
		Error:   errorStr,
		Request: rq,
	}

	errorDB := new(database.ErrorDB)
	if err := errorDB.OpenConnection(); err == nil {
		errorDB.WriteRequestErrorToDB(reqErr)
	}
}
