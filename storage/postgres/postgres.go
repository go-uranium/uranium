package postgres

import (
	"database/sql"
)

type Postgres struct {
	db *sql.DB
}

func New(dataSourceName string) (*Postgres, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return &Postgres{}, err
	}
	if err := db.Ping(); err != nil {
		return &Postgres{}, err
	}
	return &Postgres{
		db: db,
	}, nil
}

func (pg *Postgres) Close() error {
	return pg.db.Close()
}
