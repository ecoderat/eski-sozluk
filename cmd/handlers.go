package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/ecoderat/eski-sozluk/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

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
	name := r.URL.Query().Get(":name")
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
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	user := r.PostForm.Get("user")

	errors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 50 {
		errors["title"] = "This field is too long (maximum is 50 characters)"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}

	if strings.TrimSpace(user) == "" {
		errors["user"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(user) > 40 {
		errors["user"] = "This field is too long (maximum is 40 characters)"
	}

	if len(errors) > 0 {
		t, err := app.entries.LatestTopics()
		if err != nil {
			app.serverError(w, err)
			return
		}
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
			Topics:     t,
		})
		return
	}

	name, err := app.entries.Insert(title, content, user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/topic/%s", name), http.StatusSeeOther)
}

func (app *application) createEntryForm(w http.ResponseWriter, r *http.Request) {

	t, err := app.entries.LatestTopics()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "create.page.tmpl", &templateData{
		Topics: t,
	})

}
