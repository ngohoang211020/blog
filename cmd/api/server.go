package main

import (
	"errors"
	"fmt"
	"github.com/ngohoang211020/blog/internal/config"
	"github.com/ngohoang211020/blog/internal/jsonlog"
	"github.com/ngohoang211020/blog/internal/models"
	"net/http"
	"time"
)

type application struct {
	logger *jsonlog.Logger
	model  models.Models
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Config.Port),
		Handler: app.routes(),
		//ErrorLog:     log.New(app.logger, "", 0),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	app.logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  config.Config.Host,
	})
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	app.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})
	return nil
}
