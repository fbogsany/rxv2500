package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes() http.Handler {
	router := httprouter.New()
	router.PUT("/volume/:value", volume)
	router.GET("/volume", volume)
	router.PUT("/power/:value", power)
	router.GET("/power", power)
	return router
}
