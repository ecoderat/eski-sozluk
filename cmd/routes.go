package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/topic/:name", http.HandlerFunc(app.showTopic))
	mux.Get("/topic/create", http.HandlerFunc(app.createEntry))
	mux.Post("/topic/create", http.HandlerFunc(app.createEntry))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
