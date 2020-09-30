package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Sozluk struct {
	ID      int
	Title   string
	Content string
	User    string
	Created time.Time
}
