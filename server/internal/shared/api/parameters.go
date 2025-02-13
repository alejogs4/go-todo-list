package api

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ErrPathParamNotFound = errors.New("path param not found")
)

func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func GetPathParam(r *http.Request, key string) (string, error) {
	muxVars := mux.Vars(r)

	if muxVars[key] == "" {
		return "", ErrPathParamNotFound
	}

	return muxVars[key], nil
}
