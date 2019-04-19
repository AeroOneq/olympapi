package models

type RequestError struct {
	Id      int      `json:"id"`
	Error   string   `json:"error"`
	Request *Request `json:"request"`
}
