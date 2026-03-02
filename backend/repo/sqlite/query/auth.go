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

const AuthTokenInsert = `
	insert into t_auth(token, flags)
	values(?, ?);
`

const AuthUdateFlagsByToken = `
	update t_auth
		set flags = ?
	where (token = ?);
`
