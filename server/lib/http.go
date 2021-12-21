package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorObject struct {
	Message string `json:"message"`
}

func WriteJsonBadRequest(rw http.ResponseWriter, message string) {
	WriteJson(rw, http.StatusBadRequest, &errorObject{Message: message})
}

func WriteJsonInternalError(rw http.ResponseWriter) {
	WriteJson(rw, http.StatusBadRequest, &errorObject{Message: "internal error happened"})
}

func WriteJsonOk(rw http.ResponseWriter, res interface{}) {
	WriteJson(rw, http.StatusOK, res)
}

func WriteJson(rw http.ResponseWriter, status int, res interface{}) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(res)
	if err != nil {
		log.Printf("Error writing response: %s", err)
	}
}
