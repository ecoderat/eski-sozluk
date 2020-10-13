package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ecoderat/eski-sozluk/pkg/forms"

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
	t, err := app.entries.LatestTopics()
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content", "user")
	form.MaxLength("title", 50)
	form.MaxLength("user", 40)
	form.MinLength("user", 3)

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{
			Form:   form,
			Topics: t,
		})
		return
	}

	_, err = app.users.Authenticate(form.Get("user"), form.Get("password"))
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.Errors.Add("generic", "Name or Password is incorrect")
			app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		} else {
			app.serverError(w, err)
		}
		return
	}

	name, err := app.entries.Insert(form.Get("title"), form.Get("content"), form.Get("user"))
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

	form := forms.New(r.PostForm)

	app.render(w, r, "create.page.tmpl", &templateData{
		Form:   form,
		Topics: t,
	})

}
