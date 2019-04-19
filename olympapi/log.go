package olympapi

import (
	"OlympApi/models"
	"net/http"

	"Olympapi/apidatabase"

	"github.com/go-martini/martini"
)

func getLogger() martini.Handler {
	return func(r *http.Request) {
		var requestModel *models.Request = getRequestObject(r)
		apidatabase.WriteRequestToDB(requestModel)
	}
}

func getRequestObject(rq *http.Request) *models.Request {
	var request *models.Request = &models.Request{}
	(*request).GetRequest(*rq)

	return request
}
