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
	router.HandlerFunc(http.MethodGet, "/v1/users/me", app.requireAuthenticatedUser(app.getUserHandler))

	router.HandlerFunc(http.MethodGet, "/v1/mediaentries", app.requireAuthenticatedUser(app.listMediaEntriesHandler))
	router.HandlerFunc(http.MethodPost, "/v1/mediaentries", app.requireAuthenticatedUser(app.createMediaEntryHandler))
	router.HandlerFunc(http.MethodGet, "/v1/mediaentries/:id", app.requireAuthenticatedUser(app.showMediaEntryHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/mediaentries/:id", app.requireAuthenticatedUser(app.updateMediaEntryHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/mediaentries/:id", app.requireAuthenticatedUser(app.deleteMediaEntryHandler))

	router.HandlerFunc(http.MethodGet, "/v1/search/movies", app.requireAuthenticatedUser(app.searchMoviesByTitleHandler))
	router.HandlerFunc(http.MethodGet, "/v1/find/movie", app.requireAuthenticatedUser(app.getMoviesByIMDbHandler))
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.requireAuthenticatedUser(app.getMovieHandler))

	router.HandlerFunc(http.MethodGet, "/v1/search/tvshows", app.requireAuthenticatedUser(app.searchTVShowsByTitleHandler))
	router.HandlerFunc(http.MethodGet, "/v1/find/tvshow", app.requireAuthenticatedUser(app.getTVShowsByIMDbHandler))
	router.HandlerFunc(http.MethodGet, "/v1/tvshows/:id", app.requireAuthenticatedUser(app.getTVShowHandler))

	router.HandlerFunc(http.MethodGet, "/v1/search/books", app.requireAuthenticatedUser(app.searchBooksByTitleHandler))
	router.HandlerFunc(http.MethodGet, "/v1/find/book", app.requireAuthenticatedUser(app.searchBooksByISBNHandler))
	router.HandlerFunc(http.MethodGet, "/v1/books/:id", app.requireAuthenticatedUser(app.getBookHandler))

	router.HandlerFunc(http.MethodGet, "/v1/youtube", app.requireAuthenticatedUser(app.getYoutubeVideoHandler))

	router.HandlerFunc(http.MethodPost, "/v1/share", app.requireAuthenticatedUser(app.shareMediaEntryHandler))
	router.HandlerFunc(http.MethodGet, "/v1/shared", app.requireAuthenticatedUser(app.listSharedEntriesHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/shared/:id", app.requireAuthenticatedUser(app.deleteSharedEntryHandler))

	return app.recoverPanic(app.enableCORS(app.authenticate(router)))
}
