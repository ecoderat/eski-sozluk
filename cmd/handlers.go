package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ecoderat/eski-sozluk/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.entries.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	t, err := app.entries.LatestTopics()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Entries: s,
		Entry:   s[0],
		Topics:  t,
	})

}

func (app *application) showTopic(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		app.notFound(w)
		return
	}

	s, err := app.entries.GetTopic(name)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	t, err := app.entries.LatestTopics()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "topic.page.tmpl", &templateData{
		Entries: s,
		Entry:   s[0],
		Topics:  t,
	})

}

func (app *application) createEntry(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "pena"
	content := "gitar çalmak için kullanılır."
	user := "ssg"

	name, err := app.entries.Insert(title, content, user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/topic?name=%s", name), http.StatusSeeOther)
}
