package mysql

// --- userの新規作成 ---
var CreateUserState = `
	insert into
		users(
			name,
			email,
			password,
			created_at
		)
	values (?, ?, ?, ?)
`

// --- userの取得 ---
var FindUserState = `
	select
		id,
		name,
		email,
		password,
		created_at
	from
		users
	where
		id = ?
`
