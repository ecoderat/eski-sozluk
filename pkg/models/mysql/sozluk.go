package mysql

import (
	"database/sql"

	"github.com/ecoderat/eski-sozluk/pkg/models"
)

type SozlukModel struct {
	DB *sql.DB
}

func (m *SozlukModel) Insert(title, content, user string) (string, error) {
	stmt := `INSERT INTO sozluk (title, content, user, created)
	VALUES(?, ?, ?, UTC_TIMESTAMP())`

	_, err := m.DB.Exec(stmt, title, content, user)
	if err != nil {
		return "", err
	}

	return title, nil
}

func (m *SozlukModel) Get(id int) (*models.Sozluk, error) {
	return nil, nil
}

func (m *SozlukModel) Latest() ([]*models.Sozluk, error) {
	return nil, nil
}
