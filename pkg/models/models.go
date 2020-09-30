package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Entry struct {
	ID      int
	Title   string
	Content string
	User    string
	Created time.Time
}
