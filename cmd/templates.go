package main

import "github.com/ecoderat/eski-sozluk/pkg/models"

type templateData struct {
	Entry   *models.Entry
	Entries []*models.Entry
	Topics  []string
}
