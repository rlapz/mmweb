package query

const TodoInsert = `
	insert into t_todo(id_user, label, is_active, created_at)
	values(?, ?, ?, ?);
`

const TodoInsertItems = `
	insert into t_todo_item(id_todo, title, description, deadline, status, created_at)
	values
`

const TodoIsExists = `
	select 1
	from t_todo
	where (label = ?) and (id_user = ?)
	limit 1;
`

const TodoIsExistsById = `
	select 1
	from t_todo
	where (id = ?)
	limit 1;
`

const TodoItemIsExists = `
	select 1
	from t_todo_item
	where (id_todo = ?) and (title = ?)
	limit 1;
`

const TodoSelectById = `
	select id, id_user, label, is_active, created_at, updated_at
	from t_todo
	where (id = ?) and (is_active = 1);
`

const TodoSelectByUserId = `
	select a.id, a.id_user, a.label, a.is_active, a.created_at, a.updated_at
	from t_todo as a
	join t_user as b on (b.id = a.id_user)
	where (b.id = ?) and (is_active = 1);
`

const TodoSelectItemById = `
	select a.id, id_todo, title, description, deadline, status, a.created_at, a.updated_at
	from t_todo_item as a
	join t_todo as b on (b.id = a.id_todo)
	where (a.id = ?) and (b.is_active = 1);
`

const TodoSelectItemsByTodoId = `
	select a.id, id_todo, title, description, deadline, status, a.created_at, a.updated_at
	from t_todo_item as a
	join t_todo as b on (b.id = a.id_todo)
	where (a.id = ?) and (b.is_active = 1);
`

const TodoUpdate = `
	with src_t(id, label, updated_at) as (
		values(?, ?, ?)
	)
	update t_todo
		set label = a.label,
		    updated_at = a.updated_at
	from src_t as a
	where (t_todo.id = a.id)
		and (t_todo.label != a.label)
		and (t_todo.is_active = 1);
`

const TodoUpdateIsActive = `
	with src_t(id, is_active, updated_at) as (
		values(?, ?, ?)
	)
	update t_todo
		set is_active = a.is_active,
		    updated_at = a.updated_at
	from src_t as a
	where (t_todo.id = a.id);
`

const TodoUpdateItem = `
       with upd_src(id, title, description, deadline, status, updated_at) as (
               values (?, ?, ?, ?, ?, ?)
       )
       update t_todo_item
               set title = a.title,
                   description = a.description,
		   deadline = a.deadline,
                   status = a.status,
                   updated_at = a.updated_at 
       from upd_src as a
       where (t_todo_item.id = a.id) and (
               (t_todo_item.title != a.title)
               or (t_todo_item.deadline != a.deadline)
               or (t_todo_item.description != a.description)
               or (t_todo_item.status != a.status)
       );
`

const TodoInsertHistory = `
	insert into t_todo_history(id_todo, label, is_active, created_at)
	select id, label, is_active, ?
	from t_todo
	where (id = ?);
`

const TodoUpdateItemStatus = `
	update t_todo_item
		set status = ?,
		    updated_at = ?
	where (id = ?);
`

const TodoInsertItemHistory = `
	insert into t_todo_item_history(id_todo_item, title, description, deadline, status,
					created_at)
	select id, title, description, deadline, status, ?
	from t_todo_item
	where (id = ?);
`
