package repository

import (
	"database/sql"
	"errors"
	"github.com/tacheshun/ejobsclone/pkg/models"
)

type AdModel struct {
	DB *sql.DB
}

func (m *AdModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO ejobs.ads (title, content, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *AdModel) Get(id int) (*models.Ad, error) {
	s := &models.Ad{}
	err := m.DB.QueryRow("SELECT id, title, content, created, expires FROM ejobs.ads WHERE expires > UTC_TIMESTAMP() AND id = ?", id).
		Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *AdModel) Latest() ([]*models.Ad, error) {
	stmt := `SELECT id, title, content, created, expires FROM ejobs.ads
    WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ads []*models.Ad
	for rows.Next() {
		s := &models.Ad{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		ads = append(ads, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ads, nil
}
