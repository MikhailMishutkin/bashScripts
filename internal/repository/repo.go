package repository

import "github.com/jackc/pgx/v5"

type Repo struct {
	db *pgx.Conn
}

func NewRepo(db *pgx.Conn) *Repo {
	return &Repo{
		db: db,
	}
}
