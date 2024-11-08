package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users/register", app.createUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users/login", app.createAuthenticationTokenHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/logout", app.requireAuthenticatedUser(app.invalidateAuthenticationTokenHandler))

	router.HandlerFunc(http.MethodGet, "/v1/mediaentries", app.requireAuthenticatedUser(app.listMediaEntriesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/mediaentries", app.requireAuthenticatedUser(app.createMediaEntryHandler))
	router.HandlerFunc(http.MethodGet, "/v1/mediaentries/:id", app.requireAuthenticatedUser(app.showMediaEntryHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/mediaentries/:id", app.requireAuthenticatedUser(app.updateMediaEntryHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/mediaentries/:id", app.requireAuthenticatedUser(app.deleteMediaEntryHandler))

	return app.recoverPanic(app.enableCORS(app.authenticate(router)))
}
