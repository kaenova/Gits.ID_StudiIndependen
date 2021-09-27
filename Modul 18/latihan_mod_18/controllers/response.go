package controllers

import "net/http"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse() Response {
	var r Response
	r.Status = http.StatusBadRequest
	r.Message = "Bad Reqeust"
	return r
}
