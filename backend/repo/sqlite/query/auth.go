package query

const AuthSelectByToken = `
	select id, id_user, token, flags
	from t_auth
	where (token = ?);
`

const AuthInsert = `
	insert into t_auth(id_user, token, flags)
	values(?, ?, ?);
`

const AuthUdateFlagsByToken = `
	update t_auth
		set flags = ?,
		    updated_at = ?
	where (token = ?);
`

const AuthInsertHistory = `
	insert into t_auth_history(id_auth, token, flags, created_at)
	select id, token, flags, ?
	from t_auth
	where (token = ?);
`
