package sqlite

import (
	"context"

	"github.com/rlapz/mmweb/repo/sqlite/query"
	"github.com/rlapz/mmweb/util"
)

func (r *Repo) AuthTokenInsert(ctx context.Context, token string, flags int) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	_, err := util.DbTryExec(ctx, conn.Db, query.AuthTokenInsert, token, flags)
	return err
}

func (r *Repo) AuthTokenUpdateFlags(ctx context.Context, token string, flags int) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	_, err := util.DbTryExec(ctx, conn.Db, query.AuthTokenUpdateFlags, flags, token)
	return err
}

func (r *Repo) AuthTokenSelectFlags(ctx context.Context, token string) (int, error) {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	var ret int
	row := conn.Db.QueryRowContext(ctx, query.AuthTokenSelectFlags, token)
	if err := row.Scan(&ret); err != nil {
		return -1, err
	}

	return ret, nil
}
