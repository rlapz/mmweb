package query

const InsertTodo = `
	INSERT INTO t_todo(id_user, title, description, flags, created_at, created_by)
	SELECT a.id, ?, ?, ?, ?, a.id
	FROM t_user AS a
	WHERE (a.name = ?)
	LIMIT 1
`

const SelectTodoById = `
	SELECT id, id_user, title, description, flags, created_at, created_by
	FROM t_todo
	WHERE (id = ?)
`

const SelectTodoByUsername = `
	SELECT a.id, a.id_user, a.title, a.description, a.flags, a.created_at, a.created_by
	FROM t_todo AS a
	JOIN t_user AS b ON (a.id_user = b.id)
	WHERE (b.name = ?)
`
