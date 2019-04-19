package api

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/go-martini/martini"
)

var PASS string = "pass"
var LOGIN string = "login"

func AuthorizeRequest(pass MartiniHandler) MartiniHandler {
	return func(rq *http.Request, params martini.Params) (int, string) {
		authData := strings.SplitN(rq.Header.Get("Authorization"), " ", 2)

		if len(authData) != 2 || authData[0] != "Basic" {
			return http.StatusUnauthorized, "Unauthorized"
		}

		logNdPassString, err := base64.StdEncoding.DecodeString(authData[1])

		if err != nil {
			return http.StatusBadRequest, "Wrong encoding, use base64"
		}

		logAndPassPair := strings.SplitN(string(logNdPassString), ":", 2)
		if len(logAndPassPair) != 2 || !validateLoginAndPass(logAndPassPair[0], logAndPassPair[1]) {
			return http.StatusUnauthorized, "Unauthorized request"
		}

		return pass(rq, params)
	}
}

func validateLoginAndPass(login, password string) bool {
	if login == LOGIN || password == PASS {
		return true
	}
	
	return false
}
