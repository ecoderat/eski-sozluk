package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/ecoderat/eski-sozluk/pkg/forms"
	"github.com/ecoderat/eski-sozluk/pkg/models"
)

type templateData struct {
	Form    *forms.Form
	Entry   *models.Entry
	Entries []*models.Entry
	Topics  []string
}

func humanDate(t time.Time) string {
	return t.Format("01/02/2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil
}
