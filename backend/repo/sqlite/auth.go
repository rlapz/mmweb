package sqlite

import (
	"context"

	"github.com/rlapz/mmweb/db"
	"github.com/rlapz/mmweb/repo/sqlite/query"
	"github.com/rlapz/mmweb/util"
)

type Auth struct {
	db *db.SqlitePool
}

func AuthNew(db *db.SqlitePool) *Auth {
	return &Auth{
		db: db,
	}
}

func (a *Auth) TokenInsert(ctx context.Context, token string, flags int) error {
	conn := a.db.GetConn()
	defer a.db.PutConn(conn)

	_, err := util.DbTryExec(ctx, conn.Db, query.AuthTokenInsert, token, flags)
	return err
}

func (a *Auth) TokenUpdateFlags(ctx context.Context, token string, flags int) error {
	conn := a.db.GetConn()
	defer a.db.PutConn(conn)

	_, err := util.DbTryExec(ctx, conn.Db, query.AuthTokenUpdateFlags, flags, token)
	return err
}

func (a *Auth) TokenSelectFlags(ctx context.Context, token string) (int, error) {
	conn := a.db.GetConn()
	defer a.db.PutConn(conn)

	var ret int
	row := conn.Db.QueryRowContext(ctx, query.AuthTokenSelectFlags, token)
	if err := row.Scan(&ret); err != nil {
		return -1, err
	}

	return ret, nil
}
