package sqlite

import (
	"context"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/repo/sqlite/query"
	"github.com/rlapz/mmweb/util"
)

func (r *Repo) AuthTokenInvalidInsert(ctx context.Context, token string) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	count, err := util.DbTxTryExec(ctx, conn.Db, query.AuthTokenInvalidInsert, token)
	if err != nil {
		return err
	}

	if count == 0 {
		return errorx.NoDataSaved
	}

	return nil
}

func (r *Repo) AuthTokenInvalidCheck(ctx context.Context, token string) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	var ret bool
	row := conn.Db.QueryRowContext(ctx, query.AuthTokenInvalidCheck, token)
	return row.Scan(&ret)
}
