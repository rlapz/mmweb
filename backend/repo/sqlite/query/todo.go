package query

const TodoInsert = `
	insert into t_todo(id_user, label, created_at)
	values(?, ?, ?);
`

const TodoInsertItems = `
	insert into t_todo_item(id_todo, title, description, flags, created_at)
	values
`

const TodoIsExists = `
	select 1
	from t_todo
	where (label = ?) and (id_user = ?)
	limit 1;
`

const TodoItemIsExists = `
	select 1
	from t_todo_item
	where (id_todo = ?) and (title = ?)
	limit 1;
`

const TodoSelectById = `
	select id, id_user, label, created_at, updated_at
	from t_todo
	where (id = ?);
`

const TodoSelectByUserId = `
	select a.id, a.id_user, a.label, a.created_at, a.updated_at
	from t_todo as a
	join t_user as b on (b.id = a.id_user)
	where (b.id = ?);
`

const TodoSelectItemById = `
	select id, id_todo, title, description, flags, created_at, updated_at
	from t_todo_item
	where (id = ?);
`

const TodoSelectItemsByTodoId = `
	select id, id_todo, title, description, flags, created_at, updated_at
	from t_todo_item
	where (id_todo = ?);
`

const TodoUpdateItemFlags = `
	update t_todo_item
		set flags = ?,
		    updated_at = ?
	where (id = ?);
`

const TodoInsertItemHistory = `
	insert into t_todo_item_history(id_todo_item, title, description, flags,
					created_at)
	select id, title, description, flags, ?
	from t_todo_item
	where (id = ?);
`
