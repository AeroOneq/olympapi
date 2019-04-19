package models

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Request struct {
	Id int `json:"id"`

	RequestDate time.Time `josn:"requestTime"`

	Method        string              `json:"method"`
	Proto         string              `json:"proto"`
	Header        http.Header         `json:"header"`
	ContentLength int64               `json:"contentLength"`
	RemoteAddr    string              `json:"remoteAddr"`
	Host          string              `json:"host"`
	RequestURI    string              `json:"requestURI"`
	Body          string              `json:"body"`
	FormValue     map[string][]string `json:"formValue"`
	FormFile      []string            `json:"formFile"`
}

func (request *Request) GetRequest(rq http.Request) {
	var body string

	reqURI, err := ioutil.ReadAll(rq.Body)
	if err == nil {
		body = string(reqURI)
	}

	*request = Request{
		Id:            0,
		RequestDate:   time.Now(),
		Method:        rq.Method,
		Proto:         rq.Proto,
		Header:        rq.Header,
		ContentLength: rq.ContentLength,
		RemoteAddr:    rq.RemoteAddr,
		Host:          rq.Host,
		RequestURI:    rq.URL.String(),
		Body:          body,
	}
}
