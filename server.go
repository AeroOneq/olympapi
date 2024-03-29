package main

import (
	"OlympApi/olympapi"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	api := olympapi.GetApi()

	//start the HTTP and redirect all requests to HTTPS
	go func() {
		http.ListenAndServe(":8080", http.HandlerFunc(redirectToHTTPS))
	}()

	if err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", api); err != nil {
		fmt.Println(err)
	}
	appengine.Main()
}

func redirectToHTTPS(wr http.ResponseWriter, r *http.Request) {
	httpsURL := "https://" + (*r).Host + (*r).URL.Path
	if len((*r).URL.RawQuery) > 0 {
		httpsURL += "?" + (*r).URL.RawQuery
	}

	http.Redirect(wr, r, httpsURL, http.StatusTemporaryRedirect)
}
