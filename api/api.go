package api

import (
	"net/http"

	"../handlers"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/secure"
)

//typical martini handler, the start function in request handling accepts funcs
//like this as an argument (e.g. AuthorizeRequest(GetUserJsonByID), where GetUserJsonByID is of that type)
type MartiniHandler func(rq *http.Request, params martini.Params) (int, string)

func GetApi() (api *martini.Martini) {
	martini.Env = martini.Prod
	api = martini.New()

	api.Use(getLogger())
	api.Use(martini.Recovery)
	api.Use(secure.Secure(secure.Options{
		SSLRedirect: false,
		SSLHost:     "localhost:8443",
	}))

	router := martini.NewRouter()
	InitializeRouter(router)

	api.Action(router.Handle)

	return
}

func InitializeRouter(router martini.Router) {
	//users routes
	router.Get("/olympapi/v.1.0/users/(?P<id>[0-9]+)", AuthorizeRequest(handlers.GetUserJsonByID))
	router.Get("/olympapi/v.1.0/users", AuthorizeRequest(handlers.GetAllUsersJson))
	router.Post("/olympapi/v.1.0/users", AuthorizeRequest(handlers.CreateNewUserRecord))
}
