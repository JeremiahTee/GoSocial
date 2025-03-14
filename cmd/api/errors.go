package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s", r.Method, r.URL.Path, err)

	app.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "err", err.Error())
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s", r.Method, r.URL.Path, err)

	app.logger.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "err", err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request error: %s path: %s", r.Method, r.URL.Path, err)

	app.logger.Errorf("status not found", "method", r.Method, "path", r.URL.Path, "err", err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("conflict error: %s path: %s", r.Method, r.URL.Path, err)

	app.logger.Warnf("status conflict", "method", r.Method, "path", r.URL.Path, "err", err.Error())
}
