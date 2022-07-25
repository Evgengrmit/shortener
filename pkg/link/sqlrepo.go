package link

import "database/sql"

type LinkSQL struct {
	DB *sql.DB
}

func NewLinkSQL(db *sql.DB) *LinkSQL {
	return &LinkSQL{
		DB: db,
	}
}

func (ls *LinkSQL) Add(original string) (string, error) {
	var short string
	_, err := ls.DB.Exec("INSERT INTO links (`shortURL`, `longURL`) VALUES (&1, $2)", short, original)
	if err != nil {
		return "", err
	}
	return short, nil
}

func (ls *LinkSQL) Get(short string) (string, error) {
	row := ls.DB.QueryRow("SELECT longURL FROM links WHERE shortURL=$sl", short)
	var original string
	err := row.Scan(&original)
	if err != nil {
		return "", err
	}
	return original, nil
}
