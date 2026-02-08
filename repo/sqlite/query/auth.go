package query

const AuthTokenInvalidInsert = `
	INSERT INTO t_auth_token_invalid(value)
	VALUES (?)
`

const AuthTokenInvalidCheck = `
	SELECT 1
	FROM t_auth_token_invalid
	WHERE (value = ?)
`
