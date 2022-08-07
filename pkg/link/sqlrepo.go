package link

import (
	"database/sql"
	"ozonTask/shorter"
)

type LinkSQL struct {
	DB *sql.DB
}

func NewLinkSQL(db *sql.DB) *LinkSQL {
	return &LinkSQL{
		DB: db,
	}
}

func (ls *LinkSQL) Add(originalURL string) (string, error) {
	shortURL, err := shorter.GetShort(originalURL)
	if err != nil {
		return "", err
	}
	_, err = ls.Get(shortURL)
	if err != nil {
		_, err := ls.DB.Exec("INSERT INTO links (short, original) VALUES ($1, $2)", shortURL, originalURL)
		if err != nil {
			return "", err
		}
	}

	return shortURL, nil
}

func (ls *LinkSQL) Get(shortURL string) (string, error) {
	row := ls.DB.QueryRow("SELECT original FROM links WHERE short=$1", shortURL)
	var originalURL string
	err := row.Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
