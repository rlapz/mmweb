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
