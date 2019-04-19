package handlers

import "OlympApi/apidatabase"
import "OlympApi/models"

func HandleError(rq *models.Request, errorStr string) {
	reqErr := &models.RequestError{
		Error:   errorStr,
		Request: rq,
	}

	errorDB := new(apidatabase.ErrorDB)
	if err := errorDB.OpenConnection(); err == nil {
		errorDB.WriteRequestErrorToDB(reqErr)
	}
}
