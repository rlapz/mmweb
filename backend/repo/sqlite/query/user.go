package query

const UserSelectByName = `
	select a.id, a.uname, b.first_name, b.last_name, b.email, b.password,
	       b.flags, a.created_at, b.updated_at
	from t_user as a
	join t_user_detail as b on (b.id_user = a.id)
	where (a.uname = ?)
`

const UserInsert = `
	insert into t_user(uname, created_at)
	values(?, ?);
`

const UserDetailInsert = `
	insert into t_user_detail(id_user, first_name, last_name, email,
				  password, flags, created_at)
	values(?, ?, ?, ?, ?, ?, ?);
`

const UserIsExists = `
	select 1
	from t_user_detail as a
	join t_user as b
		on (b.id = a.id_user)
	where (b.uname = ?)
	limit 1;
`
