package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
)

type Entry struct {
	ID      int
	Title   string
	Content string
	User    string
	Created time.Time
}

type User struct {
	ID             int
	Name           string
	HashedPassword []byte
	Created        time.Time
}
