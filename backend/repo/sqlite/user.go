package sqlite

import (
	"context"
	"database/sql"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/repo/sqlite/query"
	"github.com/rlapz/mmweb/util"
)

func (r *Repo) UserSelectIdByName(ctx context.Context, uname string) (int32, error) {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	var id int32
	row := conn.Db.QueryRowContext(ctx, query.UserSelectIdByName, uname)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repo) UserSelectPasswordByName(ctx context.Context, uname string) (string, error) {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	var passwd string
	row := conn.Db.QueryRowContext(ctx, query.UserSelectPasswordByName, uname)
	if err := row.Scan(&passwd); err != nil {
		return "", err
	}

	return passwd, nil
}

func (r *Repo) UserInsert(ctx context.Context, user *model.User) error {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	now := util.Now()
	err := util.DbTxExecWithHandler(ctx, conn.Db, func(ctx context.Context, tx *sql.Tx, args ...any) error {
		res, err := util.DbTxTryPartialExec(ctx, tx, query.UserInsert, user.Name, now)
		if err != nil {
			return err
		}

		aff, err := util.DbTryRowsAffected(ctx, res)
		if err != nil {
			return err
		}

		if aff == 0 {
			return errorx.NoDataSaved
		}

		id, err := util.DbTryLastInsertId(ctx, res)
		if err != nil {
			return err
		}

		res, err = util.DbTxTryPartialExec(ctx, tx, query.UserDetailInsert, int(id),
			user.FirstName, user.LastName, user.Email, user.Password, user.Flags,
			now)
		if err != nil {
			return err
		}

		aff, err = util.DbTryRowsAffected(ctx, res)
		if err != nil {
			return err
		}

		if aff == 0 {
			return errorx.NoDataSaved
		}

		return nil
	})

	return err
}

func (r *Repo) UserIsExists(ctx context.Context, uname string) (bool, error) {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	row := conn.Db.QueryRowContext(ctx, query.UserIsExists, uname)
	return util.DbDataIsExists(row)
}
