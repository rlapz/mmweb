package query

const SelectUserPasswordByName = `
	SELECT b.password
	FROM t_user AS a
	JOIN t_user_detail AS b
		ON (b.id_user = a.id)
	WHERE (a.name = ?)
	ORDER BY b.id DESC
	LIMIT 1;

`
