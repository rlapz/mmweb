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
