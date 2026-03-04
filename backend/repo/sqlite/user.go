package sqlite

import (
	"context"
	"database/sql"

	"github.com/rlapz/mmweb/db"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/repo/sqlite/query"
	"github.com/rlapz/mmweb/util"
)

type User struct {
	db *db.SqlitePool
}

func UserNew(db *db.SqlitePool) *User {
	return &User{
		db: db,
	}
}

func (u *User) SelectByName(ctx context.Context, uname string) (*model.User, error) {
	conn := u.db.GetConn()
	defer u.db.PutConn(conn)

	ret := new(model.User)
	row := conn.Db.QueryRowContext(ctx, query.UserSelectByName, uname)
	err := row.Scan(&ret.Id, &ret.Name, &ret.FirstName, &ret.LastName, &ret.Email, &ret.Password,
		&ret.Flags, &ret.CreatedAt, &ret.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (u *User) Insert(ctx context.Context, user *model.User) error {
	conn := u.db.GetConn()
	defer u.db.PutConn(conn)

	now := util.Now()
	err := util.DbTxExecWithHandler(ctx, conn.Db, func(ctx context.Context, tx *sql.Tx, args ...any) error {
		res, err := util.DbTxTryExecPartial(ctx, tx, query.UserInsert, user.Name, now)
		if err != nil {
			return err
		}

		id, err := util.DbTryLastInsertId(ctx, res)
		if err != nil {
			return err
		}

		res, err = util.DbTxTryExecPartial(ctx, tx, query.UserDetailInsert, int(id),
			user.FirstName, user.LastName, user.Email, user.Password, user.Flags, now)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (u *User) IsExistsByName(ctx context.Context, uname string) (bool, error) {
	conn := u.db.GetConn()
	defer u.db.PutConn(conn)

	row := conn.Db.QueryRowContext(ctx, query.UserIsExists, uname)
	return util.DbDataIsExists(row)
}
