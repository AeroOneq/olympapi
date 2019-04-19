package api

import (
	"net/http"

	"../database"
	"../models"
	"github.com/go-martini/martini"
)

func getLogger() martini.Handler {
	return func(r *http.Request) {
		var requestModel *models.Request = getRequestObject(r)
		database.WriteRequestToDB(requestModel)
	}
}

func getRequestObject(rq *http.Request) *models.Request {
	var request *models.Request = &models.Request{}
	(*request).GetRequest(*rq)

	return request
}
