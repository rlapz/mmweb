package sqlite

import (
	"context"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/repo/sqlite/query"
)

func (r *Repo) AuthTokenInvalidInsert(ctx context.Context, token string) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	res, err := conn.Db.ExecContext(ctx, query.AuthTokenInvalidInsert, token)
	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()
	if aff == 0 {
		return errorx.NoDataSaved
	}

	return err
}

func (r *Repo) AuthTokenInvalidCheck(ctx context.Context, token string) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	var ret bool
	row := conn.Db.QueryRowContext(ctx, query.AuthTokenInvalidCheck, token)
	return row.Scan(&ret)
}
