package handler

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Error string `json: "error"`
}

func SendError(err error, w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ResponseError{Error: err.Error()})
}

func SendResponse(w http.ResponseWriter, res interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
