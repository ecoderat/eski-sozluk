package mysql

import (
	"database/sql"

	"github.com/ecoderat/eski-sozluk/pkg/models"
)

type EntryModel struct {
	DB *sql.DB
}

func (m *EntryModel) Insert(title, content, user string) (string, error) {
	stmt := `INSERT INTO sozluk (title, content, user, created)
	VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err := m.DB.Exec(stmt, title, content, user)
	if err != nil {
		return "", err
	}

	return title, nil
}

func (m *EntryModel) Get(id int) (*models.Entry, error) {

	return nil, nil
}

func (m *EntryModel) GetTopic(title string) ([]*models.Entry, error) {
	stmt := `SELECT id, title, content, user, created FROM sozluk 
	WHERE title = ?`

	rows, err := m.DB.Query(stmt, title)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	entries := []*models.Entry{}

	for rows.Next() {
		s := &models.Entry{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.User, &s.Created)
		if err != nil {
			return nil, err
		}

		entries = append(entries, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil

}

func (m *EntryModel) Latest() ([]*models.Entry, error) {
	stmt := `SELECT id, title, content, user, created FROM sozluk 
	ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	entries := []*models.Entry{}

	for rows.Next() {
		s := &models.Entry{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.User, &s.Created)
		if err != nil {
			return nil, err
		}

		entries = append(entries, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
