package query

const UserSelectIdByName = `
	select id
	from t_user
	where (uname = ?)
`

const UserSelectPasswordByName = `
	select a.password
	from t_user_detail as a
	join t_user as b
		on (b.id = a.id_user)
	where (b.uname = ?);
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
