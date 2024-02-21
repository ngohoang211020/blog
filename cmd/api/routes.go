package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/ngohoang211020/blog/middleware"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodDelete, "/healthcheck", app.healthCheckHandler)

	router.HandlerFunc(http.MethodPost, "/users", app.registerUserHandler)
	return middleware.WithCorsMiddleware(router)
}
