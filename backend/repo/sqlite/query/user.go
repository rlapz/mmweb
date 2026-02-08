package query

const UserSelectPasswordByName = `
	SELECT password
	FROM t_user
	WHERE (name = ?)
`
