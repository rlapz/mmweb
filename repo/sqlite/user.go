package sqlite

import (
	"context"

	"github.com/rlapz/mmweb/repo/sqlite/query"
)

func (r *Repo) SelectUserPasswordByName(ctx context.Context, uname string) (string, error) {
	conn := r.db.GetConn()
	defer r.db.PutConn(conn)

	var passwd string
	row := conn.Db.QueryRowContext(ctx, query.SelectUserPasswordByName, uname)
	if err := row.Scan(&passwd); err != nil {
		return "", err
	}

	return passwd, nil
}
