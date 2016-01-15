package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func volume(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if r.Method == "PUT" {
	} else {
	}
}

func power(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}
