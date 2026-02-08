package query

const SelectUserPasswordByName = `
	SELECT password
	FROM t_user
	WHERE (name = ?)
`
