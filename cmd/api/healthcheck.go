package main

import (
	"github.com/ngohoang211020/blog/internal/config"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": config.Config.Host,
			"version":     "1.0.0",
		},
	}

	err := writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
