package sqlite

import (
	"context"
	"database/sql"

	"github.com/rlapz/mmweb/db"
	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
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

func (a *Auth) SelectByToken(ctx context.Context, token string) (*model.Auth, error) {
	conn := a.db.GetConn()
	defer a.db.PutConn(conn)

	ret := new(model.Auth)
	row := conn.Db.QueryRowContext(ctx, query.AuthSelectByToken, token)
	if err := row.Scan(&ret.Id, &ret.IdUser, &ret.Token, &ret.Flags); err != nil {
		return nil, err
	}

	return ret, nil
}

func (a *Auth) Insert(ctx context.Context, auth *model.Auth) error {
	conn := a.db.GetConn()
	defer a.db.PutConn(conn)

	_, err := util.DbTryExec(ctx, conn.Db, query.AuthInsert, auth.IdUser, auth.Token,
		auth.Flags)
	return err
}

func (a *Auth) UpdateFlagsByToken(ctx context.Context, token string, flags int32) error {
	conn := a.db.GetConn()
	defer a.db.PutConn(conn)

	now := util.Now()
	return util.DbTxExecWithHandler(ctx, conn.Db, func(ctx context.Context, tx *sql.Tx, _ ...any) error {
		_, err := util.DbTxTryExecPartial(ctx, tx, query.AuthInsertHistory, now, token)
		if err != nil {
			return err
		}

		res, err := util.DbTxTryExecPartial(ctx, tx, query.AuthUdateFlagsByToken, flags, now, token)
		if err != nil {
			return err
		}

		num, err := res.RowsAffected()
		if err != nil {
			return err
		}

		if num == 0 {
			return errorx.NoDataUpdated
		}

		return nil
	})
}
