package main

import (
	"fmt"
	"html/template"
	"net/http"
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

	data := &templateData{Entries: s, Entry: s[0], Topics: t}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) showTopic(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		app.notFound(w)
		return
	}

	s, err := app.entries.GetTopic(name)
	if err != nil {
		app.serverError(w, err)
		return
	}

	t, err := app.entries.LatestTopics()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// for _, entry := range s {
	// 	fmt.Fprintf(w, "%v\n", entry)
	// }

	data := &templateData{Entries: s, Entry: s[0], Topics: t}

	files := []string{
		"./ui/html/topic.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}

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
