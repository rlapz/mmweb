package query

const AuthTokenInsert = `
	insert into t_auth(token, flags)
	values(?, ?);
`

const AuthTokenUpdateFlags = `
	update t_auth
		set flags = ?
	where (token = ?);
`

const AuthTokenSelectFlags = `
	select flags
	from t_auth
	where (token = ?);
`
