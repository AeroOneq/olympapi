package main

import (
	"fmt"
	"net/http"

	"./api"
)

func main() {
	api := api.GetApi()

	//start the HTTP and redirect all requests to HTTPS
	go func() {
		http.ListenAndServe(":8080", http.HandlerFunc(redirectToHTTPS))
	}()

	if err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", api); err != nil {
		fmt.Println(err)
	}
}

func redirectToHTTPS(wr http.ResponseWriter, r *http.Request) {
	httpsURL := "https://" + (*r).Host + (*r).URL.Path
	if len((*r).URL.RawQuery) > 0 {
		httpsURL += "?" + (*r).URL.RawQuery
	}

	http.Redirect(wr, r, httpsURL, http.StatusTemporaryRedirect)
}
