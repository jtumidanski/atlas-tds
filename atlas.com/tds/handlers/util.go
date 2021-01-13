package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func readString(r *http.Request, param string) string {
	vars := mux.Vars(r)
	return vars[param]
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}
