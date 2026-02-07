package sqlite

import "github.com/rlapz/mmweb/src/util"

type Repo struct {
	db *util.SqlitePool
}

func New(db *util.SqlitePool) *Repo {
	return &Repo{
		db: db,
	}
}
