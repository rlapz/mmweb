package sqlite

import (
	"context"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/repo/sqlite/query"
	"github.com/rlapz/mmweb/util"
)

func (r *Repo) AuthTokenInsert(ctx context.Context, token string, flags int) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	res, err := util.DbTxTryExec(ctx, conn.Db, query.AuthTokenInsert, token, flags)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errorx.NoDataSaved
	}

	return nil
}

func (r *Repo) AuthTokenUpdateFlags(ctx context.Context, token string, flags int) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	res, err := util.DbTxTryExec(ctx, conn.Db, query.AuthTokenUpdateFlags, flags, token)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return errorx.NoDataUpdated
	}

	return nil
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
