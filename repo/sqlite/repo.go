package sqlite

import "github.com/rlapz/mmweb/db"

type Repo struct {
	db *db.SqlitePool
}

func New(dbb *db.SqlitePool) *Repo {
	return &Repo{
		db: dbb,
	}
}
